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
}
