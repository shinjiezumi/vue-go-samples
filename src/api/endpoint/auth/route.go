package auth

import "github.com/gin-gonic/gin"

func SetupRoute(api *gin.RouterGroup) {
	api.POST("/login", Login)
	api.GET("/refresh_token", RefreshToken)
	api.POST("/register", Register)
}

func SetupAuthenticatedRoute(api *gin.RouterGroup) {
	api.GET("/user", CurrentUser)
}
