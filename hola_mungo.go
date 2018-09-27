package main

import (
	"fmt"
)

func main() {
	/*var num1 float32 = 10.0
	var num2 float32 = 6.0
	holaMundo()
	fmt.Println("Calculadora 1")
	calculadora(num1, num2)
	fmt.Println("-------------")
	var num3 float32 = 44
	var num4 float32 = 54
	fmt.Println("Calculadora 2")
	calculadora(num3, num4)
	fmt.Println("-------------")
	fmt.Println(devolverDatos())
	fmt.Println("-------------")
	fmt.Print("Pedido 1 ---->")
	fmt.Println(gorras(45, "€"))
	fmt.Print("Pedido 2 ---->")
	fmt.Println(gorras(24, "$"))

	pantalon("mezclilla", "azul", "largo", "Sin bolsillos", "Levis")*/

	/* var peliculas [3]string
	peliculas[0] = "Batman v Superman"
	peliculas[1] = "Avengers 3"
	peliculas[2] = "Interestelar"

	peliculas2 := [3]string{"Justice League", "Lego Batman", "Batman return"}

	fmt.Println(peliculas)
	fmt.Println(peliculas2) */
	columnas := [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	filas := [8]string{"1", "2", "3", "4", "5", "6", "7", "8"}

	var tablero [8][8]string
	for i, col := range columnas {
		for j, fila := range filas {
			tablero[i][j] = col + fila
		}
	}

	for i := range columnas {
		for j := range filas {
			fmt.Print(tablero[i][j] + " ")
		}
		fmt.Println()
	}
}

func pantalon(caracteristicas ...string) {
	for _, caracteristica := range caracteristicas {
		fmt.Println(caracteristica)
	}
}

func gorras(pedido float32, moneda string) (string, float32, string) {
	precio := func() float32 {
		return pedido * 7
	}

	return "El precio del pedido: ", precio(), moneda
}

func holaMundo() {
	fmt.Println("Hola Mundo!!")
}

func devolverDatos() (dato1 string, dato2 int) {
	dato1 = "Carlos"
	dato2 = 25

	return
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
