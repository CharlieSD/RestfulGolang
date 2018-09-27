package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/contacto", Contact)

	server := http.ListenAndServe(":8080", router)

	log.Fatal(server)
}

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor Web em Go")
}

// Contact ...
func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Esta es la pagina de contacto")
}
