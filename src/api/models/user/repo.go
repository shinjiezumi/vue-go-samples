package user

import (
	"github.com/shinjiezumi/vue-go-samples/src/api/common"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r Repository) GetUserByEmail(email string) *User {
	var u User

	if err := r.db.First(&u, "email = ?", email).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		panic(err)
	}

	return &u
}

func (r Repository) Create(u *User) {
	if err := r.db.Create(u).Error; err != nil {
		panic(err)
	}
}

func (r Repository) FindUser(email, password string) *User {
	user := User{}
	r.db.First(&user, "email = ?", email)
	if err := common.ComparePassword(user.Password, password); err != nil {
		return nil
	}

	return &user
}
