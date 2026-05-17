package usecase

import (
	"errors"
	"rest-api-go/domain/request"
	"rest-api-go/domain/response"
	"rest-api-go/models"
	"rest-api-go/repositories"
	"rest-api-go/utils"
)

type UserUsecase interface {
	Register(req request.RegistrationRequest) error
	Login(req request.LoginRequest) (*response.LoginResponse, error)
}

type UserUsecaseInteractor struct {
	UserRepository repositories.UserRepository
}

func UserUsecaseImpl(
	userRepository repositories.UserRepository,
) UserUsecase {
	return &UserUsecaseInteractor{
		UserRepository: userRepository,
	}
}

func (u *UserUsecaseInteractor) Register(req request.RegistrationRequest) error {
	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	var status string
	if req.Status == "" {
		status = "ACTIVE"
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: hashed,
		Status:   status,
	}
	_, err = u.UserRepository.Create(&user)
	if err != nil {
		return errors.New("Failed register. Try again later.")
	}

	return nil
}

func (u *UserUsecaseInteractor) Login(req request.LoginRequest) (*response.LoginResponse, error) {
	user := models.User{
		Email:    req.Email,
		Password: req.Password,
	}
	err := utils.ValidateCredential(&user)
	if err != nil {
		return nil, errors.New("Failed validate credentials.")
	}

	token, err := utils.GenerateToken(req.Email, user.ID)
	if err != nil {
		return nil, errors.New("Failed generate token.")
	}

	return &response.LoginResponse{
		Token: token,
	}, nil
}
