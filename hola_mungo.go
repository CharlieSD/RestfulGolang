package main

import (
	"fmt"
)

type Gorra struct {
	marca  string
	color  string
	precio float32
	plana  bool
}

func main() {
	var gorraNegra = Gorra{
		marca:  "Nike",
		color:  "Negro",
		precio: 25.50,
		plana:  false}

	var gorraRoja = Gorra{"Adidas", "Roja", 25.50, false}

	var numero1 = 10.0
	var numero2 = 6.0

	fmt.Println(gorraNegra)
	fmt.Println(gorraRoja)

	fmt.Print("La suma es: ")
	fmt.Println(numero1 + numero2)

	fmt.Print("La resta es: ")
	fmt.Println(numero1 - numero2)

	fmt.Print("La multiplicación es: ")
	fmt.Println(numero1 * numero2)

	fmt.Print("La división es: ")
	fmt.Println(numero1 / numero2)
}
