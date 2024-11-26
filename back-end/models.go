package main

import (
	"context"
	"errors"
	"log"
	"sort"
)

type Product struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func GetAllProducts() ([]Product, error) {
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

	sort.Slice(products, func(i, j int) bool {
		return products[i].Id < products[j].Id
	})

	return products, nil
}

func GetProductById(id int) (*Product, error) {
	product := &Product{}
	query := "SELECT id, name, price FROM products WHERE id=$1"
	err := db.QueryRow(context.Background(), query, id).Scan(&product.Id, &product.Name, &product.Price)
	if err != nil {
		log.Printf("Error fetching product by Id: %v", err)
		return nil, errors.New("product not found")
	}
	return product, nil
}

func AddProduct(name string, price int) (int, error) {
	var id int
	err := db.QueryRow(
		context.Background(), "INSERT INTO products(name, price) VALUES ($1, $2) RETURNING id",
		name, price,
	).Scan(&id)

	return id, err
}

func EditProduct(id int, name string, price int) error {
	query := "UPDATE products SET name=$1, price=$2 WHERE id=$3"
	cmd, err := db.Exec(context.Background(), query, name, price, id)
	if err != nil {
		log.Printf("Error updating product: %v", err)
		return err
	}
	if cmd.RowsAffected() == 0 {
		return errors.New("product not found")
	}
	return nil
}

func DeleteProduct(id int) error {
	query := "DELETE FROM products WHERE id=$1"
	cmd, err := db.Exec(context.Background(), query, id)
	if err != nil {
		log.Printf("Error deleting product: %v", err)
		return err
	}
	if cmd.RowsAffected() == 0 {
		return errors.New("product not found")
	}
	return nil
}
