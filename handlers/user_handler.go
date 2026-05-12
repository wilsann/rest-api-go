package handlers

import (
	"net/http"
	"rest-api-go/domain/request"
	"rest-api-go/models"
	"rest-api-go/utils"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	user := models.User{
		Email:    req.Email,
		Password: req.Password,
	}
	err := user.ValidateCredential()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	token, err := utils.GenerateToken(req.Email, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"token":   token,
	})
}

func UserCreate(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	err := user.Create()
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
