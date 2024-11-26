package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	idParam := queryParams.Get("id")

	if idParam != "" {
		id, err := strconv.Atoi(idParam)
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		product, err := GetProductById(id)
		if err != nil {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(product); err != nil {
			http.Error(w, "Failed to encode products to JSON", http.StatusInternalServerError)
			return
		}
		return
	}

	products, err := GetAllProducts()
	if err != nil {
		http.Error(w, "Unable to retrieve products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "Failed to encode products to JSON", http.StatusInternalServerError)
		return
	}
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	id, err := AddProduct(product.Name, product.Price)
	if err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	product.Id = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(product); err != nil {
		http.Error(w, "Failed to encode products to JSON", http.StatusInternalServerError)
		return
	}
}

func EditProductHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	idParam := queryParams.Get("id")

	if idParam == "" {
		http.Error(w, "Missing product ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = EditProduct(id, product.Name, product.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("Product updated successfully")); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	idParam := queryParams.Get("id")

	if idParam == "" {
		http.Error(w, "Missing product ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	err = DeleteProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("Product deleted successfully")); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}
