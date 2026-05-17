package models

type User struct {
	ID        int64  `json:"id"`
	Name      string `binding:"required" json:"name"`
	Email     string `binding:"required" json:"email"`
	Phone     string `json:"phone"`
	Password  string `binding:"required" json:"password"`
	Status    string `json:"status"`
	AddressID string `json:"address_id"`

	Address Address `gorm:"foreignKey:AddressID" json:"address"`
}
