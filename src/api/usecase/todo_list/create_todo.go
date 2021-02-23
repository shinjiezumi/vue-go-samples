package todo_list

import (
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/shinjiezumi/vue-go-samples/src/api/common"
	"github.com/shinjiezumi/vue-go-samples/src/api/database"
	"github.com/shinjiezumi/vue-go-samples/src/api/domain/todo"
)

type TodoRequest struct {
	Title     string `json:"title" validate:"required,gte=1,lte=128"`
	Memo      string `json:"memo" validate:"lte=255"`
	LimitDate string `json:"limit_date" validate:"required,limitDate"`
}

type createTodoUseCase struct{}

func NewCreateTodoUseCase() *createTodoUseCase {
	return &createTodoUseCase{}
}

func (s *createTodoUseCase) Execute(c *gin.Context, userID uint64, r TodoRequest) bool {
	// バリデーション
	v := validator.New()
	_ = v.RegisterValidation("limitDate", validateLimitDate)
	if err := v.Struct(r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": common.ExtractValidationErrorMsg(err)})
		return true
	}

	// 保存
	limitDate, _ := time.Parse(common.DateFormat, r.LimitDate)
	t := todo.Todo{
		UserId:    userID,
		Title:     r.Title,
		Memo:      r.Memo,
		LimitDate: limitDate,
	}
	todo.NewRepository(database.Conn).Create(&t)

	return false
}

func validateLimitDate(fl validator.FieldLevel) bool {
	m := regexp.MustCompile("^[0-9]{4}-[0-9]{2}-[0-9]{2}$")
	return m.MatchString(fl.Field().String())
}
