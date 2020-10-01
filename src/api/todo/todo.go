package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/shinjiezumi/vue-go-samples/src/api/auth"
	"github.com/shinjiezumi/vue-go-samples/src/api/messages"
	"github.com/shinjiezumi/vue-go-samples/src/api/models"
	"net/http"
	"strconv"
)

type todoParams struct {
	Title      string  `form:"title" json:"title"`
	Memo       string  `form:"memo" json:"memo"`
	LimitDate  string  `form:"limit_date" json:"limit_date"`
	FinishedAt *string `form:"finished_at" json:"finished_at"`
}

func GetList(c *gin.Context) {
	user := auth.GetLoginUser(c)
	todoList := models.FindTodos(user.Id, c.DefaultQuery("is_show_finished", "false"))

	c.JSON(http.StatusOK, gin.H{
		"data": todoList,
	})
}

func Store(c *gin.Context) {
	var params todoParams
	if err := c.ShouldBindBodyWith(&params, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": messages.MESSAGE_REQUIRED_ERROR,
		})
		return
	}

	// TODO バリデーション＋CSRF

	user := auth.GetLoginUser(c)

	err := models.StoreTodo(user.Id, params.Title, params.Memo, params.LimitDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": messages.MESSAGE_GENERAL_ERROR,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": messages.MESSAGE_STORED,
	})
}

func Modify(c *gin.Context) {
	var params todoParams
	if err := c.ShouldBindBodyWith(&params, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": messages.MESSAGE_REQUIRED_ERROR,
		})
		return
	}
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	// TODO バリデーション＋CSRF

	user := auth.GetLoginUser(c)

	err := models.UpdateTodo(id, user.Id, params.Title, params.Memo, params.LimitDate, params.FinishedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": messages.MESSAGE_GENERAL_ERROR,
		})
	}

	c.JSON(200, gin.H{
		"message": messages.MESSAGE_MODIFIED,
	})
}

func Remove(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	user := auth.GetLoginUser(c)

	// TODO CSRF

	err := models.DeleteTodo(id, user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": messages.MESSAGE_GENERAL_ERROR,
		})
	}

	c.JSON(200, gin.H{
		"message": messages.MESSAGE_DELETED,
	})
}
