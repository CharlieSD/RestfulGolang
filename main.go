package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola mundo desde mi servidor Web em Go")
	})

	server := http.ListenAndServe(":8080", nil)

	log.Fatal(server)
}
