package main

import "fmt"

/* Tipos de datos:
- string: Cadenas de texto
- int: Números enteros
- uint: Números enteros positivos
- bool: Valores booleanos (true o false)
- float32: Números decimales
*/

func main() {
	var x string = "Hello, World!"
	fmt.Println(x)

	// Inferencia de tipos... Go lo asigna automáticamente
	y := "Hello, World!"
	fmt.Println(y)

	// Constantes
	const z string = "Hello, World!"
	const a = 10 // Go asigna el tipo automáticamente, solo para constantes

	fmt.Println(z)
	fmt.Println(a)

	prueba(true, 19)
	prueba(false, 19)
	prueba(true, 15)
}

func prueba(has_license bool, current_age int) {
	license := has_license
	age := current_age

	if license && age > 15 {
		fmt.Println("Puede seguir circulando")
	} else {
		fmt.Println("No puede seguir circulando")
	}
}
