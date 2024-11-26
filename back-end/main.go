package main

import (
	_ ""
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	initDb()
	defer closeDb()

	r := mux.NewRouter()
	r.HandleFunc("/products", getProductsHandler).Methods("GET")
	r.HandleFunc("/products", CreateProductHandler).Methods("POST")

	port := "8080"
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
