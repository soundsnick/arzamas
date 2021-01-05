package models

import (
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Model struct {
	ID        uint64     `form:"id" gorm:"primary_key"`
	CreatedAt time.Time  `binding:"-" form:"-"`
	UpdatedAt time.Time  `binding:"-" form:"-"`
	DeletedAt *time.Time `binding:"-" form:"-"`
}

var db *gorm.DB

func SetDB(connection string) {
	var err error
	db, err = gorm.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}

func AutoMigrate() {
	db.AutoMigrate(&User{})
}
