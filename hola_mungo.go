package main

import (
	"fmt"
)

func main() {
	var suma = 8 + 9
	var resta int = 6 - 4
	var nombre string = "Carlos "
	var apellido string = "Espinoza" // Declaración tipada

	pais := "México" // Asignación inferida

	fmt.Println("Hola Mundo desde Go con " + nombre + apellido + " desde " + pais)
	fmt.Println(suma)
	fmt.Println(resta)
	fmt.Println(nombre)
}

/*
	En terminal comando:
	- godoc: muestra la documentación de alguna biblioteca
	- go fmt [fichero]: da correcto formato del archivo
	- go build [fichero]: crea el compilado
*/
