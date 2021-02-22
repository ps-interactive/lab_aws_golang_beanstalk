package handlers

import (
	"net/http"
)

// RootHandler returns a hello world message
func RootHandler(res http.ResponseWriter, req *http.Request) {
	message := "Hello from Code Pipeline"
	res.Write([]byte(message))
}
