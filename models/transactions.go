package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID          int64          `json:"id"`
	UserID      int64          `json:"user_id"`
	TotalAmount float64        `json:"total_amount"`
	Status      string         `json:"status"`
	CreatedAt   time.Time      `json:"created_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`

	User User `gorm:"foreignKey:UserID" json:"user"`
}
