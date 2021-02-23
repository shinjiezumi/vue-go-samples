package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin/binding"

	"github.com/shinjiezumi/vue-go-samples/src/api/usecase/auth"

	"github.com/shinjiezumi/vue-go-samples/src/api/domain/user"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var r auth.RegisterRequest
	if err := c.ShouldBindBodyWith(&r, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	auth.NewRegisterUseCase().Execute(c, r)
}

func Login(c *gin.Context) {
	var r auth.LoginRequest
	if err := c.ShouldBindBodyWith(&r, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	auth.NewLoginUseCase().Execute(c, r)
}

func CurrentUser(c *gin.Context) {
	u := GetLoginUser(c)
	c.JSON(http.StatusOK, gin.H{
		"name": u.Name,
	})
}

func RefreshToken(c *gin.Context) {
	authMiddleware, err := auth.CreateAuthMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	authMiddleware.RefreshHandler(c)
}

func GetLoginUser(c *gin.Context) *user.User {
	authMiddleware, err := auth.CreateAuthMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return authMiddleware.IdentityHandler(c).(*user.User)
}

func MiddlewareFunc() gin.HandlerFunc {
	authMiddleware, err := auth.CreateAuthMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return authMiddleware.MiddlewareFunc()
}
