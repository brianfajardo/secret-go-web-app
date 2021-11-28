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

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		landingHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		var errCode = http.StatusNotFound
		var errMsg = http.StatusText(errCode)
		http.Error(w, errMsg, errCode)
	}
}

func main() {
	var router Router

	fmt.Println("Starting server on port :3000")
	http.ListenAndServe(":3000", router)
}

// WIP - pick up on lesson 3.4!
