package routes

import (
	"rest-api-go/config"
	"rest-api-go/handlers"
	"rest-api-go/middlewares"
	"rest-api-go/repositories"
	"rest-api-go/usecase"

	"github.com/gin-gonic/gin"
)

func TransactionRoute(router *gin.Engine) {
	transactionRepository := repositories.TransactionRepositoryImpl(config.DB)
	transactionDetailRepository := repositories.TransactionDetailRepositoryImpl(config.DB)
	productRepository := repositories.ProductRepositoryImpl(config.DB)
	transactionUsecase := usecase.TransactionUsecaseImpl(*transactionRepository, *transactionDetailRepository, *productRepository)
	transactionHandler := handlers.TransactionHandlerImpl(transactionUsecase)

	transactionRoutes := router.Group("/v1/api/transaction")
	transactionRoutes.Use(middlewares.Authenticate)
	transactionRoutes.GET("/list", transactionHandler.TransactionList)
	transactionRoutes.GET("/detail", transactionHandler.TransactionDetail)
	transactionRoutes.POST("/checkout", transactionHandler.TransactionCreate)
	transactionRoutes.PUT("/update", transactionHandler.TransactionUpdateStatus)
	transactionRoutes.DELETE("/delete", transactionHandler.TransactionDelete)
}
