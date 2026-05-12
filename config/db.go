package config

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InnitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "kios_warga.db")
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		phone TEXT NOT NULL,
		password TEXT NOT NULL
	)`
	_, err := DB.Exec(createUserTable)

	createProductTable := `
	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT,
		price TEXT NOT NULL,
		image_url TEXT
	);
	`
	_, err = DB.Exec(createProductTable)
	if err != nil {
		panic("Failed to create products table: " + err.Error())
	}
}
