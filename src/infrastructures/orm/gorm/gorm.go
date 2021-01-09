package gorm

import (
	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New(db *sql.DB) (*gorm.DB, error) {
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return gormDB, nil
}
