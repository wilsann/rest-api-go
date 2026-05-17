package utils

import (
	"errors"
	"rest-api-go/config"
	"rest-api-go/models"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ValidateCredential(u *models.User) error {
	inputPassword := u.Password
	result := config.DB.Select("id", "password").Where("email = ?", u.Email).Find(&u)
	if result.Error != nil {
		return errors.New("Failed to fetch user by email")
	}

	validPassword := CheckPasswordHash(inputPassword, u.Password)
	if !validPassword {
		return errors.New("Invalid Credentials")
	}

	return nil
}
