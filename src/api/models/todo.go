package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/shinjiezumi/vue-go-samples/src/api/database"
	"log"
	"time"
)

// TODO gormをembeddedする
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
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	query := db.Order("limit_date")
	if isShowFinished == "true" {
		query = query.Where("user_id = ?", userId)
	} else {
		query = query.Where("user_id = ? AND finished_at IS NULL", userId)
	}
	var todos []Todo
	query.Find(&todos)

	return &todos
}

func StoreTodo(userId uint64, title, memo, limitDate string) error {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	now := time.Now().Format("2006-01-02 15:04:05")
	todo := Todo{UserId: userId, Title: title, Memo: memo, LimitDate: limitDate, CreatedAt: now, UpdatedAt: now}
	result := db.Create(&todo)

	return result.Error
}

func UpdateTodo(id uint64, userId uint64, title, memo, limitDate string, finishedAt *string) error {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	// 取得
	todo := Todo{}
	result := db.Where("id = ? AND user_id = ?", id, userId).First(&todo)
	if result.Error != nil {
		return result.Error
	}

	// 更新
	now := time.Now().Format("2006-01-02 15:04:05")
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
	result = db.Save(todo)

	return result.Error
}

func DeleteTodo(id uint64, userId uint64) error {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	// 取得
	todo := Todo{}
	result := db.Where("id = ? AND user_id = ?", id, userId).First(&todo)
	if result.Error != nil {
		return result.Error
	}

	// 削除
	result = db.Delete(todo)

	return result.Error
}
