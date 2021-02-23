package todo_list

import (
	"github.com/gin-gonic/gin"
)

func SetupRoute(api *gin.RouterGroup) {
	api.GET("/todos", GetTodoList)
	api.POST("/todos", Create)
	api.PUT("/todos/:id", Update)
	api.DELETE("/todos/:id", Delete)
	api.PUT("/todos/:id/finished", Finished)
	api.PUT("/todos/:id/unfinished", UnFinished)
}
