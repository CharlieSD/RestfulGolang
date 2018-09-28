// Todas las acciones
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

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

func responseMovie(w http.ResponseWriter, status int, results Movie) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}

func responseMovies(w http.ResponseWriter, status int, results []Movie) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}

var collection = getSession().DB("go").C("movies")

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor Web em Go")
}

// MovieList ...
func MovieList(w http.ResponseWriter, r *http.Request) {
	var results []Movie
	err := collection.Find(nil).Sort("-_id").All(&results)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Resultados", results)
	}

	responseMovies(w, 200, results)
}

// MovieShow ...
func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieID := params["id"]

	if !bson.IsObjectIdHex(movieID) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movieID)

	results := Movie{}

	err := collection.FindId(oid).One(&results)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	responseMovie(w, 200, results)
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

	responseMovie(w, 200, movieData)
}

// MovieUpdate ...
func MovieUpdate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieID := params["id"]

	if !bson.IsObjectIdHex(movieID) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movieID)
	decoder := json.NewDecoder(r.Body)

	var movieData Movie
	err := decoder.Decode(&movieData)

	if err != nil {
		panic(err)
		w.WriteHeader(501)
		return
	}

	defer r.Body.Close()

	document := bson.M{"_id": oid}
	change := bson.M{"$set": movieData}
	err = collection.Update(document, change)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	responseMovie(w, 200, movieData)
}

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// MovieRemove ...
func MovieRemove(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieID := params["id"]

	if !bson.IsObjectIdHex(movieID) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movieID)
	err := collection.RemoveId(oid)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	results := Message{"success", "La pelicula con ID " + movieID + " ha sido borrada correctamente"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)

}
