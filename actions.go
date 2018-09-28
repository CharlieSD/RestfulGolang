// Todas las acciones
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var movies = Movies{
	Movie{"Sin límites", 2013, "Desconocido"},
	Movie{"Batman Begins", 1999, "Scorsese"},
	Movie{"Rápidos y furiosos", 2005, "Juan Antonio"}}

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor Web em Go")
}

// MovieList ...
func MovieList(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Esta es la pagina de Películas")

	json.NewEncoder(w).Encode(movies)
}

// MovieShow ...
func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieID := params["id"]

	fmt.Fprintf(w, "Has cargado la pelicula numero %s", movieID)
}

func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var movieData Movie
	err := decoder.Decode(&movieData)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	log.Println(movieData)

	movies = append(movies, movieData)
}
