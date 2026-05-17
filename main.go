package main

import (
	"log"
	"rest-api-go/config"
	"rest-api-go/routes"
	"rest-api-go/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InnitDB()
	server := gin.Default()

	err := utils.SeedAll(config.DB)
	if err != nil {
		log.Fatal("Failed to seed data.")
	}

	routes.ProductRoute(server)
	routes.UserRoute(server)
	routes.TransactionRoute(server)

	server.Run(":8080")
}
