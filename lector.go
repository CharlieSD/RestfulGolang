package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Lectura de fichero")

	text, err := ioutil.ReadFile("fichero.txt")
	showError(err)
	fmt.Println(string(text))
}

func showError(e error) {
	if e != nil {
		panic(e)
	}
}
