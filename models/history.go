package models

import (
	"gorm.io/gorm"
	"time"
)

type History struct {
	gorm.Model
	UserID uint
	User User `gorm:"foreignKey:UserID"`
	Date time.Time
	Action string
}