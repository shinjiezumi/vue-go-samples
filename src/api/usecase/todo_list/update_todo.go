package todo_list

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/shinjiezumi/vue-go-samples/src/api/common"
	"github.com/shinjiezumi/vue-go-samples/src/api/database"
	"github.com/shinjiezumi/vue-go-samples/src/api/domain/todo"
)

type updateTodoUseCase struct{}

func NewUpdateTodoUseCase() *updateTodoUseCase {
	return &updateTodoUseCase{}
}

func (s *updateTodoUseCase) Execute(c *gin.Context, userID, id uint64, r TodoRequest) bool {
	// バリデーション
	v := validator.New()
	_ = v.RegisterValidation("limitDate", validateLimitDate)
	if err := v.Struct(r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": common.ExtractValidationErrorMsg(err)})
		return true
	}

	// 更新
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

	t.Title = r.Title
	t.Memo = r.Memo
	limitDate, _ := time.Parse(common.DateFormat, r.LimitDate)
	t.LimitDate = limitDate

	repo.Save(t)

	return false
}
