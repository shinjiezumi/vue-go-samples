package searcher

import "github.com/gin-gonic/gin"

func SetupRoute(api *gin.RouterGroup) {
	api.GET("/searcher/search", Search)
}
