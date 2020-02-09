package lib

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Connect() (*gorm.DB, error) {
	DBMS := "mysql"
	USER := "user"
	PASS := "pass"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "mysql"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		return nil, err
	}
	return db, nil
}
