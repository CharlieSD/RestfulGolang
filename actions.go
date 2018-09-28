// Todas las acciones
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/gorilla/mux"
)

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return session
}

var collection = getSession().DB("go").C("movies")

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
	json.NewEncoder(w).Encode(movies)
}

// MovieShow ...
func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieID := params["id"]

	fmt.Fprintf(w, "Has cargado la pelicula numero %s", movieID)
}

// MovieAdd ...
func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var movieData Movie
	err := decoder.Decode(&movieData)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	err = collection.Insert(movieData)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	json.NewEncoder(w).Encode(movieData)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}
