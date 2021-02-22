package main

import (
	"net/http"
	"os"
	"web-api/handlers"
)

var beanstalkDefaultPort = "5000"

func main() {
	port := os.Getenv("PORT")
	// If not SET via ENV, then 
	// use Beanstalk default port
	if len(port) == 0 {
		port = beanstalkDefaultPort
	}

	http.HandleFunc("/", handlers.RootHandler)

	http.ListenAndServe(":"+port, nil)
}