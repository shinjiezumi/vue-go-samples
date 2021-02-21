package todo

import (
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"vgs/auth"
	"vgs/common"
	"vgs/database"
	"vgs/models/todo"
)

type todoRequest struct {
	Title     string `json:"title" validate:"required,gte=1,lte=128"`
	Memo      string `json:"memo" validate:"lte=255"`
	LimitDate string `json:"limit_date" validate:"required,limitDate"`
}

type todoResponse struct {
	Id         uint64 `json:"id"`
	UserId     uint64 `json:"user_id"`
	Title      string `json:"title"`
	Memo       string `json:"memo"`
	LimitDate  string `json:"limit_date"`
	FinishedAt string `json:"finished_at"`
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
	var r todoRequest
	if err := c.ShouldBindBodyWith(&r, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	v := validator.New()
	_ = v.RegisterValidation("limitDate", validateLimitDate)
	if err := v.Struct(r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": common.ExtractValidationErrorMsg(err)})
		return
	}

	limitDate, _ := time.Parse(common.DateFormat, r.LimitDate)
	t := todo.Todo{
		UserId:    auth.GetLoginUser(c).Id,
		Title:     r.Title,
		Memo:      r.Memo,
		LimitDate: limitDate,
	}
	todo.NewRepository(database.Conn).Create(&t)

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

	var r todoRequest
	if err := c.ShouldBindBodyWith(&r, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	v := validator.New()
	_ = v.RegisterValidation("limitDate", validateLimitDate)
	if err := v.Struct(r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": common.ExtractValidationErrorMsg(err)})
		return
	}

	repo := todo.NewRepository(database.Conn)
	t := repo.GetById(id)
	if t == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": common.NotFound,
		})
		return
	} else if auth.GetLoginUser(c).Id != t.UserId {
		c.JSON(http.StatusForbidden, gin.H{
			"message": common.Forbidden,
		})
		return
	}

	t.Title = r.Title
	t.Memo = r.Memo
	limitDate, _ := time.Parse(common.DateFormat, r.LimitDate)
	t.LimitDate = limitDate

	repo.Save(t)

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

	repo := todo.NewRepository(database.Conn)
	t := repo.GetById(id)
	if t == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": common.NotFound,
		})
		return
	} else if auth.GetLoginUser(c).Id != t.UserId {
		c.JSON(http.StatusForbidden, gin.H{
			"message": common.Forbidden,
		})
		return
	}

	repo.Delete(t.Id)

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

	repo := todo.NewRepository(database.Conn)
	t := repo.GetById(id)
	if t == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": common.NotFound,
		})
		return
	} else if auth.GetLoginUser(c).Id != t.UserId {
		c.JSON(http.StatusForbidden, gin.H{
			"message": common.Forbidden,
		})
		return
	}

	t.Finished(time.Now())
	repo.Save(t)

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

	repo := todo.NewRepository(database.Conn)
	t := repo.GetById(id)
	if t == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": common.NotFound,
		})
		return
	} else if auth.GetLoginUser(c).Id != t.UserId {
		c.JSON(http.StatusForbidden, gin.H{
			"message": common.Forbidden,
		})
		return
	}

	t.UnFinished()
	repo.Save(t)

	c.JSON(200, gin.H{
		"message": common.Modified,
	})
}

func validateLimitDate(fl validator.FieldLevel) bool {
	m := regexp.MustCompile("^[0-9]{4}-[0-9]{2}-[0-9]{2}$")
	return m.MatchString(fl.Field().String())
}
