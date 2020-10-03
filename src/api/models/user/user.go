package user

import (
	"time"
)

// User
type User struct {
	Id        uint64     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"created_at" sql:"DEFAULT:current_timestamp"`
	UpdatedAt *time.Time `json:"updated_at" sql:"DEFAULT:current_timestamp on update current_timestamp"`
}
