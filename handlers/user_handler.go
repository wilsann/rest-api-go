package handlers

import (
	"net/http"
	"rest-api-go/domain/request"
	"rest-api-go/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type UserHandlerInteractor struct {
	usecase usecase.UserUsecase
}

func UserHandlerImpl(userUsecase usecase.UserUsecase) UserHandler {
	return &UserHandlerInteractor{
		usecase: userUsecase,
	}
}

func (u *UserHandlerInteractor) Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	result, err := u.usecase.Login(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Login Failed. Try again later.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"token":   result,
	})
}

func (u *UserHandlerInteractor) Register(c *gin.Context) {
	var req request.RegistrationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}
	err := u.usecase.Register(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed create user.",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}
