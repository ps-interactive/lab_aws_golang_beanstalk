package handlers

import (
	"net/http"
)

// RootHandler returns a hello world message
func RootHandler(res http.ResponseWriter, req *http.Request) {
	message := "Hello from PS Labs"
	res.Write([]byte(message))
}
