package models

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	TodoText  string
}
