package routes

import (
	"rest-api-go/handlers"
	"rest-api-go/middlewares"

	"github.com/gin-gonic/gin"
)

func ProductRoute(router *gin.Engine) {
	productRoutes := router.Group("/v1/api/product")
	productRoutes.GET("/list", handlers.ProductList)
	productRoutes.GET("/detail", handlers.ProductDetail)
	productRoutes.Use(middlewares.Authenticate)
	productRoutes.POST("/create", handlers.ProductCreate)
	productRoutes.PUT("/update", handlers.ProductUpdate)
	productRoutes.DELETE("/delete", handlers.ProductDelete)
}
