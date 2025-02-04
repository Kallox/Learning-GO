package main

import "fmt"

func main() {
	tarea1()
	tarea2()
}

func tarea1() {
	// Dentro de nuestro código ya tenemos definido un array, debemos reccorrerlo e incrementar todos sus valores en 20.
	// Al finalizar el programa se debera visualizar el array con los valores incrementados

	array := [5]int{4, 2, 5, 6, 7}

	// realizar la funcionalidad
	for i, v := range array {
		array[i] = v + 20
	}
	//

	fmt.Println("Los valores del array son: ", array)
}

func tarea2() {
	// Debemos implementar un programa en go para ir ingresando valores por consola hasta que se agrega cero y finaliza el programa.
	// Todos los valores ingresados por consola se deben agregar a un array e imprimirlo por pantalla al finalizar.

	fmt.Println("Ingresa tus valores para agregarlos al arreglo... ingresa 0 para deterner")
	array := []int{}
	for {
		var valor int
		fmt.Scanln(&valor)
		if valor == 0 {
			break
		}
		// Agregar el valor al arreglo
		array = append(array, valor)
	}
	fmt.Println("Los valores del array son: ", array)
}
