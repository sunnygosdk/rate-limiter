package main

import (
	"fmt"
	"net/http"
)

// GetRouter returns the router
func GetRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Rate Limiter!")
	})

	return router
}
