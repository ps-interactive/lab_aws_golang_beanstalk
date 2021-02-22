package main

import (
	"log"
	"net/http"
	"os"
	"web-api/handlers"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", handlers.RootHandler)

	http.ListenAndServe(":"+port, nil)
}