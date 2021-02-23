package todo

import (
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

func (r Repository) GetById(id uint64) *Todo {
	var t Todo

	if err := r.db.First(&t, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		panic(err)
	}

	return &t
}

func (r Repository) GetByUserId(userId uint64, isShowFinished bool) []*Todo {
	var todos []*Todo

	q := r.db.Order("limit_date").Where("user_id = ?", userId)
	if !isShowFinished {
		q = q.Where("finished_at IS NULL")
	}
	if err := q.Find(&todos).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return todos
		}
		panic(err)
	}

	return todos
}

func (r Repository) Create(t *Todo) {
	if err := r.db.Create(t).Error; err != nil {
		panic(err)
	}
}

func (r Repository) Save(t *Todo) {
	if err := r.db.Save(t).Error; err != nil {
		panic(err)
	}
}

func (r Repository) Delete(id uint64) {
	if err := r.db.Delete(Todo{}, id).Error; err != nil {
		panic(err)
	}
}
