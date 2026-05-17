package models

import "time"

type Address struct {
	ID           int64  `json:"id"`
	UserID       int64  `json:"user_id"`
	ReceiverName string `json:"receiver_name"`
	Phone        string `json:"phone"`
	AddressLine  string `json:"address_line"`
	Label        string `json:"label"`

	PostalCode int64 `json:"postal_code"`
	VillageID  int64 `json:"village_id"`
	DistrictID int64 `json:"district_id"`
	RegencyID  int64 `json:"regency_id"`
	ProvinceID int64 `json:"province_id"`

	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`

	IsDefault bool `json:"is_default"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
