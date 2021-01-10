package core

import (
	"time"

	"gorm.io/gorm"
)

// Model - General struct for models
type Model struct {
	ID        uint64         `form:"id" gorm:"primary_key"`
	CreatedAt time.Time      `binding:"-" form:"-"`
	UpdatedAt time.Time      `binding:"-" form:"-"`
	DeletedAt gorm.DeletedAt `binding:"-" form:"-" gorm:"index"`
}
