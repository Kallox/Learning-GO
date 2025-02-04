package main

import "fmt"

func main() {
	myMap := make(map[int]string) // Crear un mapa con clave int y valor string
	myMap[1] = "One"              // Asignar un valor a la clave 1
	myMap[2] = "Two"
	myMap[3] = "Three"

	fmt.Println(myMap[3]) // No es la posicion 3, es la clave 3
	fmt.Println(myMap[4]) // No existe, devuelve un string vacio porque es el valor por defecto para el tipo String

	v, ok := myMap[3] // Devuelve el valor y true porque la clave 3 existe

	fmt.Println(v, ok) // Devuelve un string vacio y false porque no existe la clave 4
}
