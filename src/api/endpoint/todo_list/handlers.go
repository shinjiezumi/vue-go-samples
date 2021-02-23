package todo_list

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/shinjiezumi/vue-go-samples/src/api/common"
	"github.com/shinjiezumi/vue-go-samples/src/api/endpoint/auth"
	"github.com/shinjiezumi/vue-go-samples/src/api/usecase/todo_list"
)

// GetTodoList はTodo一覧を返す
func GetTodoList(c *gin.Context) {
	u := auth.GetLoginUser(c)
	isShowFinished := c.DefaultQuery("is_show_finished", "false") == "true"

	res := todo_list.NewGetTodoListUseCase().Execute(u, isShowFinished)

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

// Create はTodoを作成します
func Create(c *gin.Context) {
	var r todo_list.TodoRequest
	if err := c.ShouldBindBodyWith(&r, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// TODO エラーハンドリング見直し(コンテキスト渡さない)
	errOccurred := todo_list.NewCreateTodoUseCase().Execute(c, auth.GetLoginUser(c).Id, r)
	if errOccurred {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": common.Stored,
	})
}

// Update はTodoを更新します
func Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var r todo_list.TodoRequest
	if err := c.ShouldBindBodyWith(&r, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// TODO エラーハンドリング見直し(コンテキスト渡さない)
	errOccurred := todo_list.NewUpdateTodoUseCase().Execute(c, auth.GetLoginUser(c).Id, id, r)
	if errOccurred {
		return
	}

	c.JSON(200, gin.H{
		"message": common.Modified,
	})
}

// Delete はTodoを削除します
func Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// TODO エラーハンドリング見直し(コンテキスト渡さない)
	errOccurred := todo_list.NewDeleteTodoUseCase().Execute(c, auth.GetLoginUser(c).Id, id)
	if errOccurred {
		return
	}

	c.JSON(200, gin.H{
		"message": common.Deleted,
	})
}

// Finished はTodoを完了済みにします
func Finished(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// TODO エラーハンドリング見直し(コンテキスト渡さない)
	errOccurred := todo_list.NewFinishTodoUseCase().Execute(c, auth.GetLoginUser(c).Id, id)
	if errOccurred {
		return
	}

	c.JSON(200, gin.H{
		"message": common.Modified,
	})
}

// UnFinished はTodoを未完了にします
func UnFinished(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// TODO エラーハンドリング見直し(コンテキスト渡さない)
	errOccurred := todo_list.NewUnFinishTodoUseCase().Execute(c, auth.GetLoginUser(c).Id, id)
	if errOccurred {
		return
	}

	c.JSON(200, gin.H{
		"message": common.Modified,
	})
}
