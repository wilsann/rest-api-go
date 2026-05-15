package response

import "time"

type TransactionListResponse struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	TotalAmount float64   `json:"total_amount"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type TransactionDetailResponse struct {
	ID                int64                           `json:"id"`
	UserID            int64                           `json:"user_id"`
	TotalAmount       float64                         `json:"total_amount"`
	Status            string                          `json:"status"`
	TransactionDetail []TransactionDetailDataResponse `json:"transaction_detail"`
}

type TransactionDetailDataResponse struct {
	ID              int64   `json:"id"`
	TransactionID   int64   `json:"transacrion_id"`
	ProductID       int64   `json:"product_id"`
	Quantity        int64   `json:"quantity"`
	PriceAtPurchase float64 `json:"price_at_purchase"`
	SubTotal        float64 `json:"sub_total"`
}
