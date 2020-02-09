package todo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleGetTodoList(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{
		"message": "Vue Go Samples",
	})
}
