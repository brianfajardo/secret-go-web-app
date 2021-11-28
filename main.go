package main

import (
	"fmt"
	"net/http"
)

func landingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Landing Page</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1>")
}

func routesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		landingHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Page not found.", 404)
	}
}

func main() {
	http.HandleFunc("/", routesHandler)

	fmt.Println("Starting server on port :3000")
	http.ListenAndServe(":3000", nil)
}
