package handlers

import (
	"net/http"
)

// RootHandler returns a hello world message
func RootHandler(res http.ResponseWriter, req *http.Request) {
	message := "Hello from Web API"
	res.Write([]byte(message))
}
