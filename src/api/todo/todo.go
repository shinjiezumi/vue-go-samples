package todo

import (
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/shinjiezumi/vue-go-samples/src/api/auth"
	"github.com/shinjiezumi/vue-go-samples/src/api/models"
)

func Todos(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(auth.IdentityKey)
	c.JSON(200, gin.H{
		"userID":   claims[auth.IdentityKey],
		"userName": user.(*models.User).Name,
		"text":     "Hello World.",
	})
}
