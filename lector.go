package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Lectura de fichero")

	nuevoTexto := os.Args[1] + "\n"

	//nuevoFichero := ioutil.WriteFile("elotrofichero.txt", nuevoTexto, 0777)

	fichero, err := os.OpenFile("elotrofichero.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	showError(err)

	tamaño, err := fichero.WriteString(nuevoTexto)
	fmt.Println(tamaño)
	showError(err)

	fichero.Close()

	text, err := ioutil.ReadFile("elotrofichero.txt")
	showError(err)
	fmt.Println(string(text))
}

func showError(e error) {
	if e != nil {
		panic(e)
	}
}
