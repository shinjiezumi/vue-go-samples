package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/shinjiezumi/vue-go-samples/src/api/database"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func FindUser(email, password string) *User {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	user := User{}
	db.First(&user, "email = ?", email)
	if err = compare(user.Password, password); err != nil {
		return &User{}
	}

	return &user
}

func StoreUser(name, email, password string) error {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	user := User{Name: name, Email: email, Password: hash(password)}
	result := db.Create(&user)

	return result.Error
}

func hash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err.Error())
	}

	return string(hash)
}

func compare(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
