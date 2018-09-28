package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/peliculas", MovieList)
	router.HandleFunc("/peliculas/{id}", MovieShow)

	server := http.ListenAndServe(":8080", router)

	log.Fatal(server)
}

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor Web em Go")
}

// MovieList ...
func MovieList(w http.ResponseWriter, r *http.Request) {
	movies := Movies{
		Movie{"Sin límites", 2013, "Desconocido"},
		Movie{"Batman Begins", 1999, "Scorsese"},
		Movie{"Rápidos y furiosos", 2005, "Juan Antonio"}}
	//fmt.Fprintf(w, "Esta es la pagina de Películas")

	json.NewEncoder(w).Encode(movies)
}

// MovieShow ...
func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieID := params["id"]

	fmt.Fprintf(w, "Has cargado la pelicula numero %s", movieID)
}
