package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

var db *pgxpool.Pool

func initDb() {
	dsn := "postgres://postgres:admin@localhost:5432/react_go_store?sslmode=disable"

	var err error
	db, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v/\n\n", err)
	}

	log.Println("Connected to database")
}

func closeDb() {
	if db != nil {
		db.Close()
	}
}
