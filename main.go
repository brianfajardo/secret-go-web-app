package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func landingPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Landing Page</h1>")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1>")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", landingPage)
	r.HandleFunc("/contact", contactPage)

	fmt.Println("Starting server on port :3000")
	http.ListenAndServe(":3000", r)
}
