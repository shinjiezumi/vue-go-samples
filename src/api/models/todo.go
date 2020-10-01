package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/shinjiezumi/vue-go-samples/src/api/database"
	"time"
)

type Todo struct {
	Id         uint64  `json:"id"`
	UserId     uint64  `json:"user_id"`
	Title      string  `json:"title"`
	Memo       string  `json:"memo"`
	LimitDate  string  `json:"limit_date"`
	FinishedAt *string `json:"finished_at"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

func FindTodos(userId uint64, isShowFinished string) *[]Todo {
	var query string
	if isShowFinished == "true" {
		query = "user_id = ?"
	} else {
		query = "user_id = ? AND finished_at IS NULL"
	}

	var todos []Todo
	database.Conn.Order("limit_date").Where(query, userId).Find(&todos)

	return &todos
}

func StoreTodo(userId uint64, title, memo, limitDate string) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	todo := Todo{UserId: userId, Title: title, Memo: memo, LimitDate: limitDate, CreatedAt: now, UpdatedAt: now}
	result := database.Conn.Create(&todo)

	return result.Error
}

func UpdateTodo(id uint64, userId uint64, title, memo, limitDate string, finishedAt *string) error {
	// 取得
	todo := Todo{}
	result := database.Conn.Where("id = ? AND user_id = ?", id, userId).First(&todo)
	if result.Error != nil {
		return result.Error
	}

	// 更新
	now := time.Now().Format("2006-01-02 15:04:05")
	// TODO 雑なので要見直し
	if finishedAt != nil {
		todo.FinishedAt = finishedAt
	} else if title == "" {
		todo.FinishedAt = nil
	} else {
		todo.Title = title
		todo.Memo = memo
		todo.LimitDate = limitDate
	}
	todo.UpdatedAt = now
	result = database.Conn.Save(todo)

	return result.Error
}

func DeleteTodo(id uint64, userId uint64) error {
	// 取得
	todo := Todo{}
	result := database.Conn.Where("id = ? AND user_id = ?", id, userId).First(&todo)
	if result.Error != nil {
		return result.Error
	}

	// 削除
	result = database.Conn.Delete(todo)

	return result.Error
}
