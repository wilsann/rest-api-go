package routes

import (
	"rest-api-go/config"
	"rest-api-go/handlers"
	"rest-api-go/repositories"
	"rest-api-go/usecase"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	userRepository := repositories.UserRepositoryImpl(config.DB)
	userUsecase := usecase.UserUsecaseImpl(*userRepository)
	userHandler := handlers.UserHandlerImpl(userUsecase)

	userRoutes := router.Group("/v1/api/user")
	userRoutes.POST("/register", userHandler.Register)
	userRoutes.POST("/login", userHandler.Login)
	// userRoutes.POST("/create", userHandler.UserCreate)
	// userRoutes.PUT("/update", userHandler.UserUpdate)
	// userRoutes.DELETE("/delete", userHandler.UserDelete)
}
