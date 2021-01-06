package models

import (
	"time"

	"github.com/jinzhu/gorm"
	// Use postgres dialect for GORM
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Model - General struct for models
type Model struct {
	ID        uint64     `form:"id" gorm:"primary_key"`
	CreatedAt time.Time  `binding:"-" form:"-"`
	UpdatedAt time.Time  `binding:"-" form:"-"`
	DeletedAt *time.Time `binding:"-" form:"-"`
}

var db *gorm.DB

// SetDB Configures DB
func SetDB(connection string) {
	var err error
	db, err = gorm.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
}

// GetDB Returns DB Instance
func GetDB() *gorm.DB {
	return db
}

// AutoMigrate migrates models
func AutoMigrate() {
	db.AutoMigrate(&User{}, &Post{})
}
