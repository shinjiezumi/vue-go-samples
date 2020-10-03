package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/shinjiezumi/vue-go-samples/src/api/auth"
	"github.com/shinjiezumi/vue-go-samples/src/api/common"
	"github.com/shinjiezumi/vue-go-samples/src/api/database"
	"github.com/shinjiezumi/vue-go-samples/src/api/messages"
	"github.com/shinjiezumi/vue-go-samples/src/api/models/todo"
	"net/http"
	"strconv"
	"time"
)

type todoParams struct {
	Title      string  `form:"title" json:"title"`
	Memo       string  `form:"memo" json:"memo"`
	LimitDate  string  `form:"limit_date" json:"limit_date"`
	FinishedAt *string `form:"finished_at" json:"finished_at"`
}

type todoResponse struct {
	Id         uint64 `json:"id"`
	UserId     uint64 `json:"user_id"`
	Title      string `json:"title"`
	Memo       string `json:"memo"`
	LimitDate  string `json:"limit_date"`
	FinishedAt string `form:"finished_at" json:"finished_at"`
}

// GetList はTodo一覧を取得します
func GetList(c *gin.Context) {
	user := auth.GetLoginUser(c)
	isShowFinished := c.DefaultQuery("is_show_finished", "false") == "true"

	todos := todo.NewRepository(database.Conn).GetByUserId(user.Id, isShowFinished)
	var res []todoResponse
	for _, t := range todos {
		var finishedAt string
		if t.FinishedAt != nil {
			finishedAt = t.FinishedAt.Format(common.DateFormat)
		}
		res = append(res, todoResponse{
			Id:         t.Id,
			UserId:     t.UserId,
			Title:      t.Title,
			Memo:       t.Memo,
			LimitDate:  t.LimitDate.Format(common.DateFormat),
			FinishedAt: finishedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

// Create はTodoを作成します
func Create(c *gin.Context) {
	var params todoParams
	if err := c.ShouldBindBodyWith(&params, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": messages.RequiredError,
		})
		return
	}

	// TODO バリデーション＋CSRF

	u := auth.GetLoginUser(c)
	limitDate, _ := time.Parse(common.DateFormat, params.LimitDate)
	t := todo.Todo{
		UserId:    u.Id,
		Title:     params.Title,
		Memo:      params.Memo,
		LimitDate: limitDate,
	}
	todo.NewRepository(database.Conn).Create(&t)

	c.JSON(http.StatusOK, gin.H{
		"message": messages.Stored,
	})
}

// Update はTodoを更新します
func Update(c *gin.Context) {
	var params todoParams
	if err := c.ShouldBindBodyWith(&params, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": messages.RequiredError,
		})
		return
	}
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	// TODO バリデーション＋CSRF

	u := auth.GetLoginUser(c)

	repo := todo.NewRepository(database.Conn)
	t := repo.GetById(id)
	if t == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": messages.NotFound,
		})
		return
	} else if u.Id != t.UserId {
		c.JSON(http.StatusForbidden, gin.H{
			"message": messages.Forbidden,
		})
		return
	}

	t.Title = params.Title
	t.Memo = params.Memo
	limitDate, _ := time.Parse(common.DateFormat, params.LimitDate)
	t.LimitDate = limitDate

	repo.Save(t)

	c.JSON(200, gin.H{
		"message": messages.Modified,
	})
}

// Delete はTodoを削除します
func Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	u := auth.GetLoginUser(c)

	// TODO CSRF

	repo := todo.NewRepository(database.Conn)
	t := repo.GetById(id)
	if t == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": messages.NotFound,
		})
		return
	} else if u.Id != t.UserId {
		c.JSON(http.StatusForbidden, gin.H{
			"message": messages.Forbidden,
		})
		return
	}

	repo.Delete(t.Id)

	c.JSON(200, gin.H{
		"message": messages.Deleted,
	})
}

// Finished はTodoを完了済みにします
func Finished(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	u := auth.GetLoginUser(c)

	// TODO CSRF

	repo := todo.NewRepository(database.Conn)
	t := repo.GetById(id)
	if t == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": messages.NotFound,
		})
		return
	} else if u.Id != t.UserId {
		c.JSON(http.StatusForbidden, gin.H{
			"message": messages.Forbidden,
		})
		return
	}

	t.Finished(time.Now())
	repo.Save(t)

	c.JSON(200, gin.H{
		"message": messages.Modified,
	})
}

// UnFinished はTodoを未完了にします
func UnFinished(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	u := auth.GetLoginUser(c)

	// TODO CSRF

	repo := todo.NewRepository(database.Conn)
	t := repo.GetById(id)
	if t == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": messages.NotFound,
		})
		return
	} else if u.Id != t.UserId {
		c.JSON(http.StatusForbidden, gin.H{
			"message": messages.Forbidden,
		})
		return
	}

	t.UnFinished()
	repo.Save(t)

	c.JSON(200, gin.H{
		"message": messages.Modified,
	})
}
