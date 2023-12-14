package models

import (
	"gorm.io/gorm"
	"time"
)

type Payment struct {
	gorm.Model
	TransactionID uint
	Transaction Transaction `gorm:"foreignKey:TransactionID"`
	Date time.Time
	Amount int
	Method string
	Status string
}