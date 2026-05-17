package models

type District struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	RegencyID string `json:"regency_id"`
}
