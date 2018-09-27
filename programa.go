package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
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

	hoy := time.Now()
	dia := hoy.Weekday()

	switch dia {
	case 0:
		fmt.Println("Hoy es domingo")
	case 1:
		fmt.Println("Hoy es lunes")
	case 2:
		fmt.Println("Hoy es martes")
	case 3:
		fmt.Println("Hoy es miercoles")
	case 4:
		fmt.Println("Hoy es jueves")
	case 5:
		fmt.Println("Hoy es viernes")
	case 6:
		fmt.Println("Hoy es sábado")
	default:
		fmt.Println("No se puede determinar el día")
	}

}
