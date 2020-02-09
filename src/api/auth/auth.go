package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// TODO 秘匿化
var secretKey = "75c92a074c341e9964329c0550c2673730ed8479c885c43122c90a2843177d5ef21cb50cfadcccb20aeb730487c11e09ee4dbbb02387242ef264e74cbee97213"

type User struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

func HandleRoot(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Vue Go Samples",
	})
}

func HandleLogin(context *gin.Context) {
	// TODO DB認証
	user := User{Id: 1, Name: "test"}

	tokenString, err := createToken(user)
	if err == nil {
		context.JSON(http.StatusOK, gin.H{
			"token": tokenString,
		})
	} else {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not generate token",
		})
	}
}

func HandleRegister(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Vue Go Samples",
	})
}

func createToken(user User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// アルゴリズムの指定
	token.Claims = jwt.MapClaims{
		"id":   user.Id,
		"name": user.Name,
		"exp":  time.Now().Add(1 * time.Hour).Unix(),
	}

	// 署名の付与
	return token.SignedString([]byte(secretKey))
}
