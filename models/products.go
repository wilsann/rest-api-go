package models

import (
	"rest-api-go/config"
)

type Product struct {
	ID          int64   `json:"id"`
	Name        string  `binding:"required" json:"name"`
	Description string  `json:"description"`
	Price       float64 `binding:"required" json:"price"`
	ImageURL    string  `json:"image_url"`
}

var products = []Product{}

func GetList() ([]Product, error) {
	query := "SELECT * FROM products"
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.ImageURL)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func GetByID(id int64) (*Product, error) {
	query := "SELECT * FROM products WHERE id = ?"
	row := config.DB.QueryRow(query, id)

	var p Product
	err := row.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.ImageURL)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (p *Product) Create() error {
	query := `INSERT INTO products (name, description, price, image_url) 
	VALUES (?, ?, ?, ?)`
	stmt, err := config.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(p.Name, p.Description, p.Price, p.ImageURL)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	p.ID = id
	return err
}

func (p *Product) Update() error {
	query := `UPDATE products 
	SET name = ?, description = ?, price = ?, image_url = ?
	WHERE id = ?`
	stmt, err := config.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(p.Name, p.Description, p.Price, p.ImageURL, p.ID)
	return err
}

func (p Product) Delete() error {
	query := "DELETE FROM products WHERE id = ?"
	stmt, err := config.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(p.ID)
	return err
}
