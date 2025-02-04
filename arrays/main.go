package main

import (
	"fmt"
	"unsafe"
)

func main() {
	learningArrays()
	learningSlices()
}

func learningArrays() {
	var myVector [6]int                   // Array de 6 elementos
	myVector2 := [6]int{1, 2, 3, 4, 5, 6} // Inicializacion

	myVector[0] = 1
	myVector[1] = 2
	myVector[2] = 3
	myVector[3] = 4
	myVector[4] = 5
	myVector[5] = 6

	fmt.Println(myVector)
	fmt.Println(myVector2)

	fmt.Println("Length of myVector: ", len(myVector))                     // Len devuelve el tamaño del array
	fmt.Printf("type: %T, bytes: %d\n", myVector, unsafe.Sizeof(myVector)) // unsafe.Sizeof devuelve el tamaño en bytes del array
}

func learningSlices() {
	var mySlice []int            // Similar a un vector pero no especifica el tamano
	mySlice = append(mySlice, 1) // Append agrega uno/muchos elemento al slice
	mySlice = append(mySlice, 2, 3, 4, 5)

	fmt.Println(mySlice)
	fmt.Println(mySlice[1:3])

	// Slices are references to arrays
	myArray := [4]int{1, 3, 5, 7}
	mySlice2 := myArray[1:3]                // El slice contiene los elementos 3 y 5
	myArray[1] = 13                         // Cambiar el valor de myArray[1] cambia el valor de mySlice2[0]
	mySlice2 = append(mySlice2, myArray[2]) // Al ocupar append no se modifica el valor de myArray[2]
	myArray[2] = 15                         // Esto no cambia el valor de mySlice2[1]
	fmt.Println(mySlice2)

}
