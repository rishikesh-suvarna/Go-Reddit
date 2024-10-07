package main

import (
	"log"
	"net/http"

	"github.com/rishikesh-suvarna/go-reddit/db"
	"github.com/rishikesh-suvarna/go-reddit/routes"
)

func main() {
	store, err := db.NewStore("postgres://rishikeshsuvarna@localhost:5432/go-reddit?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	handler := routes.NewHandler(*store)

	http.ListenAndServe(":8000", handler)
}
