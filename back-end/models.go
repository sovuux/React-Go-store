package main

import (
	"context"
	_ "github.com/swaggo/http-swagger"
	"log"
)

type Product struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func getAllProducts() ([]Product, error) {
	rows, err := db.Query(context.Background(), "SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Price); err != nil {
			log.Println("Error scanning rows:", err)
			continue
		}
		products = append(products, product)
	}
	return products, nil
}

func AddProduct(name string, price int) (int, error) {
	var id int
	err := db.QueryRow(
		context.Background(), "INSERT INTO products(name, price) VALUES ($1, $2) RETURNING id",
		name, price,
	).Scan(&id)

	return id, err
}
