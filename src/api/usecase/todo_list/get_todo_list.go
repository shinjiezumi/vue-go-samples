package todo_list

import (
	"github.com/shinjiezumi/vue-go-samples/src/api/common"
	"github.com/shinjiezumi/vue-go-samples/src/api/database"
	"github.com/shinjiezumi/vue-go-samples/src/api/domain/todo"
	"github.com/shinjiezumi/vue-go-samples/src/api/domain/user"
)

type TodoResponse struct {
	Id         uint64 `json:"id"`
	UserId     uint64 `json:"user_id"`
	Title      string `json:"title"`
	Memo       string `json:"memo"`
	LimitDate  string `json:"limit_date"`
	FinishedAt string `json:"finished_at"`
}

type getTodoListUseCase struct{}

func NewGetTodoListUseCase() *getTodoListUseCase {
	return &getTodoListUseCase{}
}

func (s *getTodoListUseCase) Execute(user *user.User, isShowFinished bool) []TodoResponse {
	todos := todo.NewRepository(database.Conn).GetByUserId(user.Id, isShowFinished)

	var res []TodoResponse
	for _, t := range todos {
		var finishedAt string
		if t.FinishedAt != nil {
			finishedAt = t.FinishedAt.Format(common.DateFormat)
		}
		res = append(res, TodoResponse{
			Id:         t.Id,
			UserId:     t.UserId,
			Title:      t.Title,
			Memo:       t.Memo,
			LimitDate:  t.LimitDate.Format(common.DateFormat),
			FinishedAt: finishedAt,
		})
	}

	return res
}
