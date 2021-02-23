package todo_list

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/shinjiezumi/vue-go-samples/src/api/common"
	"github.com/shinjiezumi/vue-go-samples/src/api/database"
	"github.com/shinjiezumi/vue-go-samples/src/api/domain/todo"
)

type deleteTodoUseCase struct{}

func NewDeleteTodoUseCase() *deleteTodoUseCase {
	return &deleteTodoUseCase{}
}

func (s *deleteTodoUseCase) Execute(c *gin.Context, userID, id uint64) bool {
	repo := todo.NewRepository(database.Conn)
	t := repo.GetById(id)
	if t == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": common.NotFound,
		})
		return true
	} else if userID != t.UserId {
		c.JSON(http.StatusForbidden, gin.H{
			"message": common.Forbidden,
		})
		return true
	}

	repo.Delete(t.Id)

	return false
}
