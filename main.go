package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", handleWelcomePage)

	fmt.Println("Server listening on localhost:8080!")
	http.ListenAndServe(":8080", r)
}

func handleWelcomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the best watchlist!")
}
