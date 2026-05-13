package models

import (
	"errors"
	"rest-api-go/config"
	"rest-api-go/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Name     string `binding:"required" json:"name"`
	Email    string `binding:"required" json:"email"`
	Phone    string `json:"phone"`
	Password string `binding:"required" json:"password"`
	Status   string `json:"status"`
}

func (u *User) ValidateCredential() error {
	inputPassword := u.Password
	result := config.DB.Select("id", "password").Where("email = ?", u.Email).Find(&u)
	if result.Error != nil {
		return errors.New("Failed to fetch user by email")
	}

	validPassword := utils.CheckPasswordHash(inputPassword, u.Password)
	if !validPassword {
		return errors.New("Invalid Credentials")
	}

	return nil
}
