package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InnitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	var errDB error
	DB, errDB = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errDB != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	fmt.Println("Successfully connected to the database!")
}

// func createTables() {
// 	createUserTable := `
// 	CREATE TABLE IF NOT EXISTS users (
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		name TEXT NOT NULL,
// 		email TEXT NOT NULL UNIQUE,
// 		phone TEXT NOT NULL,
// 		password TEXT NOT NULL
// 	)`
// 	_, err := DB.Exec(createUserTable)

// 	createProductTable := `
// 	CREATE TABLE IF NOT EXISTS products (
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		name TEXT NOT NULL,
// 		description TEXT,
// 		price TEXT NOT NULL,
// 		image_url TEXT,
// 		created_at TIMESTAMP,
// 		created_by INTEGER
// 	);
// 	`
// 	_, err = DB.Exec(createProductTable)
// 	if err != nil {
// 		panic("Failed to create products table: " + err.Error())
// 	}
// }
