package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/shinjiezumi/vue-go-samples/src/api/database"
	"log"
)

type User struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func FindUser(email string, password string) *User {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	user := User{}

	db.First(&user, "email = ? AND password = ?", email, password)

	return &user
}

func StoreUser(name string, email string, password string) error {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	user := User{Name: name, Email: email, Password: password}
	result := db.Create(&user)

	return result.Error
}
