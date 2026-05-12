package routes

import (
	"rest-api-go/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	userRoutes := router.Group("/v1/api/user")
	userRoutes.POST("/register", handlers.UserCreate)
	userRoutes.POST("/login", handlers.UserLogin)
	// userRoutes.POST("/create", handlers.ProductCreate)
	// userRoutes.PUT("/update", handlers.ProductUpdate)
	// userRoutes.DELETE("/delete", handlers.ProductDelete)
}
