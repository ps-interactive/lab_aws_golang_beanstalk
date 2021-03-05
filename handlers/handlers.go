package handlers

import (
	"net/http"
)

// RootHandler returns a hello world message
func RootHandler(res http.ResponseWriter, req *http.Request) {
	message := "Hello from Pluralsight Cloud Labs"
	res.Write([]byte(message))
}
