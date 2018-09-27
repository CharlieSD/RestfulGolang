package main

import (
	"fmt"
)

func main() {
	var num1 float32 = 10.0
	var num2 float32 = 6.0
	holaMundo()
	fmt.Println("Calculadora 1")
	calculadora(num1, num2)
	fmt.Println("-------------")
	var num3 float32 = 44
	var num4 float32 = 54
	fmt.Println("Calculadora 2")
	calculadora(num3, num4)
}

func holaMundo() {
	fmt.Println("Hola Mundo!!")
}

func calculadora(n1 float32, n2 float32) {
	fmt.Print("La suma es: ")
	fmt.Println(operacion(n1, n2, "+"))

	fmt.Print("La resta es: ")
	fmt.Println(operacion(n1, n2, "-"))

	fmt.Print("La multiplicación es: ")
	fmt.Println(operacion(n1, n2, "*"))

	fmt.Print("La división es: ")
	fmt.Println(operacion(n1, n2, "/"))
}

func operacion(n1 float32, n2 float32, op string) float32 {
	var res float32
	if op == "+" {
		res = n1 + n2
	}
	if op == "-" {
		res = n1 - n2
	}
	if op == "*" {
		res = n1 * n2
	}
	if op == "/" {
		res = n1 / n2
	}
	return res
}
