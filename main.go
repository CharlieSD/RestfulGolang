package main

import (
	"log"
	"net/http"
)

func main() {
	router := newRouter()
	server := http.ListenAndServe(":8080", router)

	log.Fatal(server)
}
