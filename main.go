package main

import "fmt"

func main() {
	var v = 4 // declaracion de una variable
	//Mostrammos una variable y su direccion de memoria usando el caracter &
	fmt.Printf("el valor de la variable v es %v y su direccion de memoria es %v\n", v, &v)
}
