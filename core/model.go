package core

import (
	"time"

	"gorm.io/gorm"
)

// Model - General struct for models
type Model struct {
	ID        uint64         `form:"id" gorm:"primary_key" json:"id"`
	CreatedAt time.Time      `binding:"-" form:"-" json:"created_at"`
	UpdatedAt time.Time      `binding:"-" form:"-" json:"updated_at"`
	DeletedAt gorm.DeletedAt `binding:"-" form:"-" gorm:"index" json:"deleted_at"`
}
