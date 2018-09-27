package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Hola " + os.Args[1] + " Bienvenido al programa de Go")

	edad, error := strconv.Atoi(os.Args[2])

	if error == nil {
		if edad >= 18 {
			fmt.Println("Eres mayor de edad")
		} else {
			fmt.Println("Eres menor de edad")
		}
	} else {
		fmt.Println("Hubo un error de conversión, no podemos determinar tu edad")
	}

	for i := 0; i < 15; i++ {
		fmt.Println("El numero actual es: ", i)
	}

	peliculas := []string{"El club de la pelea", "Batman v Superma", "Interestelar"}

	for _, pelicula := range peliculas {
		fmt.Println(pelicula)
	}
}
