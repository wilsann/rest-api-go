package handlers

import (
	"fmt"
	"net/http"
	"rest-api-go/domain/request"
	"rest-api-go/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionHandler interface {
	TransactionList(c *gin.Context)
	TransactionDetail(c *gin.Context)
	TransactionCreate(c *gin.Context)
	TransactionUpdateStatus(c *gin.Context)
	TransactionDelete(c *gin.Context)
}

type TransactionHandlerInteractor struct {
	usecase usecase.TransactionUsecase
}

func TransactionHandlerImpl(transactionUsecase usecase.TransactionUsecase) TransactionHandler {
	return &TransactionHandlerInteractor{
		usecase: transactionUsecase,
	}
}

func (hi *TransactionHandlerInteractor) TransactionList(c *gin.Context) {
	userId := c.GetInt64("userId")
	transactions, err := hi.usecase.TransactionList(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":      "success",
		"transactions": transactions,
	})
}

func (hi *TransactionHandlerInteractor) TransactionDetail(c *gin.Context) {
	transactionId, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid transaction ID",
		})
		return
	}

	_ = c.GetInt64("userId")

	transaction, err := hi.usecase.TransactionDetail(transactionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    transaction,
	})
}

func (hi *TransactionHandlerInteractor) TransactionCreate(c *gin.Context) {
	var req request.TransactionCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	userId := c.GetInt64("userId")

	data, err := hi.usecase.TransactionCreate(req, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed create transaction.",
		})
		fmt.Println("Error create transaction: ", err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"product": data,
	})
}

func (hi *TransactionHandlerInteractor) TransactionUpdateStatus(c *gin.Context) {
	var req request.TransactionUpdateStatusReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	userId := c.GetInt64("userId")

	data, err := hi.usecase.TransactionUpdateStatus(req, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed update transaction status.",
		})
		fmt.Println("Error update transaction status: ", err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"product": data,
	})
}

func (hi *TransactionHandlerInteractor) TransactionDelete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid transaction ID",
		})
		return
	}

	_ = c.GetInt64("userId")

	err = hi.usecase.TransactionDelete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed delete transaction: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}
