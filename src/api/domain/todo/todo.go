package todo

import (
	"time"
)

// Todo
type Todo struct {
	Id         uint64     `json:"id"`
	UserId     uint64     `json:"user_id"`
	Title      string     `json:"title"`
	Memo       string     `json:"memo"`
	LimitDate  time.Time  `json:"limit_date"`
	FinishedAt *time.Time `json:"finished_at"`
	CreatedAt  *time.Time `json:"created_at" sql:"DEFAULT:current_timestamp"`
	UpdatedAt  *time.Time `json:"updated_at" sql:"DEFAULT:current_timestamp on update current_timestamp"`
}

// Finished はTodoを完了済みにします
func (t *Todo) Finished(date time.Time) {
	t.FinishedAt = &date
}

// UnFinished はTodoを未完了にします
func (t *Todo) UnFinished() {
	t.FinishedAt = nil
}
