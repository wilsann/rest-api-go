package routes

import (
	"rest-api-go/config"
	"rest-api-go/handlers"
	"rest-api-go/middlewares"
	"rest-api-go/repositories"
	"rest-api-go/usecase"

	"github.com/gin-gonic/gin"
)

func ProductRoute(router *gin.Engine) {
	productRepository := repositories.ProductRepositoryImpl(config.DB)
	productUsecase := usecase.ProductUsecaseImpl(*productRepository)
	productHandler := handlers.ProductHandlerImpl(productUsecase)

	productRoutes := router.Group("/v1/api/product")
	productRoutes.GET("/list", productHandler.ProductList)
	productRoutes.GET("/detail", productHandler.ProductDetail)
	productRoutes.Use(middlewares.Authenticate)
	productRoutes.POST("/create", productHandler.ProductCreate)
	productRoutes.PUT("/update", productHandler.ProductUpdate)
	productRoutes.DELETE("/delete", productHandler.ProductDelete)
}
