package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	FromAccountID uint
	ToAccountID   uint
	Amount        float64
	Description   string
}