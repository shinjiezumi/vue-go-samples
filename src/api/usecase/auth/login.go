package auth

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"github.com/shinjiezumi/vue-go-samples/src/api/common"
	"github.com/shinjiezumi/vue-go-samples/src/api/database"
	"github.com/shinjiezumi/vue-go-samples/src/api/domain/user"
)

var identityKey = "name"
var secretKey = os.Getenv("JWT_SECRET_KEY")

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,gte=8,lte=16"`
}

type loginUseCase struct{}

func NewLoginUseCase() *loginUseCase {
	return &loginUseCase{}
}

func (s *loginUseCase) Execute(c *gin.Context, r LoginRequest) {
	v := validator.New()
	if err := v.Struct(r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": common.ExtractValidationErrorMsg(err)})
		return
	}

	authMiddleware, err := CreateAuthMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	authMiddleware.LoginHandler(c)
}

func CreateAuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(secretKey),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*user.User); ok {
				return jwt.MapClaims{
					"id":        v.Id,
					identityKey: v.Name,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			id := claims["id"].(float64)
			return &user.User{
				Id:   uint64(id),
				Name: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var r LoginRequest
			if err := c.ShouldBindBodyWith(&r, binding.JSON); err != nil {
				return "", fmt.Errorf(err.Error())
			}

			if u := user.NewRepository(database.Conn).FindUser(r.Email, r.Password); u != nil {
				return &user.User{
					Id:   u.Id,
					Name: u.Name,
				}, nil
			}

			// TODO リファクタ対象
			return "", fmt.Errorf(common.InvalidEmailOrPassword.String())
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
