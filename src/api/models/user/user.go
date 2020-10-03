package user

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/shinjiezumi/vue-go-samples/src/api/database"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type User struct {
	Id        uint64 `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func FindUser(email, password string) *User {
	user := User{}
	database.Conn.First(&user, "email = ?", email)
	if err := compare(user.Password, password); err != nil {
		return &User{}
	}

	return &user
}

func StoreUser(name, email, password string) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	user := User{Name: name, Email: email, Password: hash(password), CreatedAt: now, UpdatedAt: now}
	result := database.Conn.Create(&user)

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
