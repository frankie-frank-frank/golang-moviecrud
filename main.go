package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

// struct movie
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie
const contentType = "Content-Type"
const applicationJSON = "application/json"

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set(contentType, applicationJSON)
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set(contentType, applicationJSON)
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set(contentType, applicationJSON)
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func main(){
	r := mux.NewRouter()
	const movieId = "/movies/{id}"

	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Avengers Endgame", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "448227", Title: "Avengers Infinity Wars", Director: &Director{Firstname: "Johnny", Lastname: "Depp"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc(movieId, getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc(movieId, updateMovie).Methods("PUT")
	r.HandleFunc(movieId, deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}