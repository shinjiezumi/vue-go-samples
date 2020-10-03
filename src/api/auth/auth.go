package auth

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/shinjiezumi/vue-go-samples/src/api/common"
	"github.com/shinjiezumi/vue-go-samples/src/api/database"
	"github.com/shinjiezumi/vue-go-samples/src/api/messages"
	"github.com/shinjiezumi/vue-go-samples/src/api/models/user"
	"log"
	"net/http"
	"os"
	"time"
)

var secretKey = os.Getenv("JWT_SECRET_KEY")

type login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type registerParams struct {
	Name string `form:"name" json:"name" binding:"required"`
	login
}

var IdentityKey = "name"

func Register(c *gin.Context) {
	var params registerParams
	if err := c.ShouldBindBodyWith(&params, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": messages.RequiredError,
		})
		return
	}

	// TODO バリデーション＋CSRF

	repo := user.NewRepository(database.Conn)
	u := repo.GetUserByEmail(params.Email)
	if u != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": messages.EmailAlreadyExists,
		})
		return
	}

	u = &user.User{
		Name:     params.Name,
		Email:    params.Email,
		Password: common.HashPassword(params.Password),
	}
	user.NewRepository(database.Conn).Create(u)

	// JWTトークン発行
	Login(c)
}

func Login(c *gin.Context) {
	authMiddleware, err := createAuthMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	authMiddleware.LoginHandler(c)
}

func CurrentUser(c *gin.Context) {
	u := GetLoginUser(c)
	c.JSON(http.StatusOK, gin.H{
		"name": u.Name,
	})
}

func RefreshToken(c *gin.Context) {
	authMiddleware, err := createAuthMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	authMiddleware.RefreshHandler(c)
}

func MiddlewareFunc() gin.HandlerFunc {
	authMiddleware, err := createAuthMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return authMiddleware.MiddlewareFunc()
}

func GetLoginUser(c *gin.Context) *user.User {
	authMiddleware, err := createAuthMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return authMiddleware.IdentityHandler(c).(*user.User)
}

func createAuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(secretKey),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*user.User); ok {
				return jwt.MapClaims{
					"id":        v.Id,
					IdentityKey: v.Name,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			id := claims["id"].(float64)
			return &user.User{
				Id:   uint64(id),
				Name: claims[IdentityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginParams login
			if err := c.ShouldBindBodyWith(&loginParams, binding.JSON); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			u := user.NewRepository(database.Conn).FindUser(loginParams.Email, loginParams.Password)
			if u.Name != "" {
				return &user.User{
					Id:   u.Id,
					Name: u.Name,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	return authMiddleware, err
}
