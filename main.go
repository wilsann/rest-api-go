package main

import (
	"rest-api-go/config"
	"rest-api-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InnitDB()
	server := gin.Default()

	routes.ProductRoute(server)
	routes.UserRoute(server)

	server.Run(":8080")
}
