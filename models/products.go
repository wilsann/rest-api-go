package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          int64          `json:"id"`
	Name        string         `binding:"required" json:"name" gorm:"column:name"`
	Description string         `json:"description"`
	Price       float64        `binding:"required" json:"price"`
	ImageURL    string         `json:"image_url"`
	CreatedAt   time.Time      `json:"created_at"`
	CreatedBy   int64          `json:"created_by"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
