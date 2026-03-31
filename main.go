package main

import (
	"fmt"
	"net/http"
	"os"
	"encoding/json"

	"github.com/go-chi/chi/v5"
)

type Movie struct{
	ID uint `json:"ID"`
	Title string `json:"title"`
	Genre string `json:"genre"`
	Watched bool `json:"watched"`
}

var movies []Movie

var nextID uint = 1

func main() {
	loadMovies()

	for _, movie := range movies{
		if movie.ID > nextID{
			nextID = movie.ID
		}
	}
	nextID++ 

	r := chi.NewRouter()

	r.Get("/", handleWelcomePage)
	r.Get("/movies", handleGetMovies)
	r.Post("/movies", handlePostMovie)

	fmt.Println("Server listening on localhost:8080!")
	http.ListenAndServe(":8080", r)
}

func handleWelcomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the best watchlist!")
}

func handleGetMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func handlePostMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie

	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil{
		http.Error(w, "Invalid JSON" ,http.StatusBadRequest)
	}

}

func loadMovies(){
	data, err := os.ReadFile("movies.json")
	if err != nil{
		fmt.Println("Coudlnt Read file")
		return
	}
	json.Unmarshal(data, &movies)	
}
