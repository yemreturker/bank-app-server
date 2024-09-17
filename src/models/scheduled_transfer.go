package models

import (
	"time"

	"gorm.io/gorm"
)

type ScheduledTransfer struct {
	gorm.Model
	FromAccountID uint
	ToAccountID   uint
	Amount        float64
	ScheduledDate time.Time
	Executed      bool
}