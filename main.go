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

	byValue(v)
	fmt.Printf("El valor fuera de la funcion byValue es %v\n", v)
	byRef(&v)
	fmt.Printf("El valor fuera de la funcion byValue es %v\n", v)

	// vamos a crear un flujo donde vamos a inserir clientes dentro de uuna lista enlazada
	var firstClient, lastClient *Client

	var name, phoneNumber, continuar string
	keep := true
	for keep {

		fmt.Println("Cual es el nombre del Cliente: ")
		fmt.Scan(&name)
		fmt.Println("Cual es el telefono del Cliente: ")
		fmt.Scan(&phoneNumber)
		var newCustomer Client
		newCustomer.Name = name
		newCustomer.PhoneNumber = phoneNumber
		if firstClient == nil { //Es el primerCliente
			firstClient = &newCustomer
			lastClient = firstClient
		} else {
			lastClient.NextCustomer = &newCustomer
			lastClient = &newCustomer
		}
		fmt.Println("Desea inserir otro Cliente: ")
		fmt.Scan(&continuar)
		if continuar == "no" {
			keep = false
		}

	}
	listarClientes(firstClient)

}

func byValue(num int) {
	num = num + 1
	fmt.Printf("El valor dentro de la funcion byRef es %v\n", num)
}
func byRef(num *int) {
	*num = *num + 1
	fmt.Printf("El valor dentro de la funcion byRef es %v\n", *num)
}

// Ejemplo de punteros usando una structura que sera usada como una lista enlazada
type Client struct {
	Name         string `json:"Name"`
	PhoneNumber  string `json:"PhoneNumber"`
	NextCustomer *Client
}

func listarClientes(inicio *Client) {
	fmt.Println("Presentando la lista de clientes: ")
	currentClient := inicio
	if currentClient == nil {
		return
	}
	for {
		fmt.Printf("Nombre: %v \n", currentClient.Name)
		fmt.Printf("Telefono: %v \n", currentClient.PhoneNumber)
		currentClient = currentClient.NextCustomer
		if currentClient.NextCustomer == nil {
			fmt.Printf("Nombre: %v \n", currentClient.Name)
			fmt.Printf("Telefono: %v \n", currentClient.PhoneNumber)
			return
		}
	}
}
