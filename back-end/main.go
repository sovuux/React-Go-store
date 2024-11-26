package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	InitDB()
	defer CloseDB()

	r := mux.NewRouter()
	r.HandleFunc("/products", GetProductHandler).Methods("GET")
	r.HandleFunc("/products", CreateProductHandler).Methods("POST")
	r.HandleFunc("/products", EditProductHandler).Methods("PUT")
	r.HandleFunc("/products", DeleteProductHandler).Methods("DELETE")

	port := "8080"
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
