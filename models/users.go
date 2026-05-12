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
}

func (u *User) ValidateCredential() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := config.DB.QueryRow(query, u.Email)

	var retreivedPassword string
	err := row.Scan(&u.ID, &retreivedPassword)
	if err != nil {
		return errors.New("Invalid Credentials")
	}

	validPassword := utils.CheckPasswordHash(u.Password, retreivedPassword)
	if !validPassword {
		return errors.New("Invalid Credentials")
	}

	return nil
}

func (u *User) Create() error {
	query := `INSERT INTO users (name, email, phone, password) 
	VALUES (?,?,?,?)`
	stmt, err := config.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashed, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Name, u.Email, u.Phone, hashed)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = userId
	return err
}
