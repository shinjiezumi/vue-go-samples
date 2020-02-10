package auth

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/shinjiezumi/vue-go-samples/src/api/models"
	"log"
	"net/http"
	"time"
)

// TODO 秘匿化
var secretKey = "75c92a074c341e9964329c0550c2673730ed8479c885c43122c90a2843177d5ef21cb50cfadcccb20aeb730487c11e09ee4dbbb02387242ef264e74cbee97213"

type login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type registerParams struct {
	Name string `form:"name" json:"name" binding:"required"`
	login
}

var IdentityKey = "id"

func Register(c *gin.Context) {
	var registerParams registerParams
	if err := c.ShouldBind(&registerParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email and password is required",
		})
		return
	}

	// TODO バリデーション

	if err := models.StoreUser(registerParams.Name, registerParams.Email, registerParams.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "registration failure",
		})
		return
	}

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

func createAuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(secretKey),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					IdentityKey: v.Name,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.User{
				Name: claims[IdentityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginParams login
			if err := c.ShouldBind(&loginParams); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			user := models.FindUser(loginParams.Email, loginParams.Password)
			if user.Name != "" {
				return &models.User{
					Id:   user.Id,
					Name: user.Name,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*models.User); ok && v.Name == "admin" {
				return true
			}

			return false
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
