package main

import "fmt"

func main() {
	var v = 4  // declaracion de una variable
	var p *int //Declaracion de variable
	p = &v     //atribuir a direccion de memoria de la variable v
	//Mostrammos una variable y su direccion de memoria usando el caracter &
	fmt.Printf("el valor de la variable v es %v y su direccion de memoria es %v\n", v, &v)
	//Mostrammos el valor  de la  variable usando el * y su direccion de memoria como es un puntero no necesitamos ningun caracter
	fmt.Printf("el valor de la variable p es %v y su direccion de memoria es %v\n", *p, p)

	//usando el puntero podemos cambiar el valor de una variable

	*p = 8 // cambiamos el valor del puntero a 8
	//Mostrammos una variable y su direccion de memoria usando el caracter &
	fmt.Printf("el valor de la variable v es %v y su direccion de memoria es %v\n", v, &v)

}
