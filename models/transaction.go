package models

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	gorm.Model
	UserID uint
	User User `gorm:"foreignKey:UserID"`
	Date time.Time
	Status string
	Nominal int
}