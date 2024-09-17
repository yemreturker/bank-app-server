package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	AccountNumber string  `gorm:"unique"`
	Balance       float64
	UserID        uint
}