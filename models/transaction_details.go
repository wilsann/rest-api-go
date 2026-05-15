package models

import (
	"time"

	"gorm.io/gorm"
)

type TransactionDetail struct {
	ID              int64          `json:"id"`
	TransactionID   int64          `json:"transacrion_id"`
	ProductID       int64          `json:"product_id"`
	Quantity        int64          `json:"quantity"`
	PriceAtPurchase float64        `json:"price_at_purchase"`
	SubTotal        float64        `json:"sub_total"`
	CreatedAt       time.Time      `json:"created_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`

	Transaction Transaction `gorm:"foreignKey:TransactionID" json:"transaction"`
	Product     Product     `gorm:"foreignKey:ProductID" json:"product"`
}
