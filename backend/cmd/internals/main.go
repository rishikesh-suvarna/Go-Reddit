package main

import (
	"log"
	"net/http"

	"github.com/rishikesh-suvarna/go-reddit/db"
	"github.com/rishikesh-suvarna/go-reddit/web"
)

func main() {
	store, err := db.NewStore("postgres://rishikeshsuvarna@localhost:5432/go-reddit?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	handler := web.NewHandler(*store)

	http.ListenAndServe(":8000", handler)
}
