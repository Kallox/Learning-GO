package main

import "fmt"

type MyStruct struct {
	ID   int
	Name string
}

func main() {
	a := 34        // Valor de la variable
	aPointer := &a // Direccion de memoria donde vive la variable

	fmt.Println("The value is:", a)
	fmt.Println("The pointer is:", aPointer)
	fmt.Println("The value of the pointer is:", *aPointer) // Acceder al valor de la variable a traves del puntero

	// Cuando pasas un puntero a una función, la función puede modificar el valor original
	// de la variable que se pasa como argumento, dado que la direccion de memoria es la misma
	myVar := 10
	fmt.Println(&myVar)
	increment(myVar)
	incrementPointer(&myVar)

	// ------Slices---------

	var slice = []int{1, 2, 3}
	fmt.Println(slice)
	updateSlice(slice)
	fmt.Println(slice)
	fmt.Printf("Direccion: %p\n", slice) // Con %p se imprime la direccion de memoria

	// ------Structs---------

	fmt.Println("------Structs------")
	myStruct := MyStruct{ID: 1, Name: "John"} // Se crea un puntero a la estructura
	fmt.Printf("Direccion: %p\n", &myStruct)
	updatesStruct(&myStruct)
	fmt.Println(myStruct)
}

// No se pasa como puntero por lo que varia la direccion de memoria
func increment(val int) {
	fmt.Println(&val)
	val++
}

// Se pasa como puntero por lo que no varia la direccion de memoria
func incrementPointer(val *int) {
	fmt.Println(val)
	*val++
}

func updateSlice(s []int) {
	fmt.Printf("Direccion: %p\n", s) // La direccion de memoria del slice es la misma que la de el primer elemento
	fmt.Printf("Direccion 1: %p, Dirrecion 2: %p, Direccion 3: %p\n", &s[0], &s[1], &s[2])
	s[0] = 12 // Se modifica el valor del slice sin necesidad de retornarlo
}

func updatesStruct(v *MyStruct) {
	fmt.Printf("Direccion: %p\n", v) // Modificar la estructura a traves del puntero
	v.ID = 2
	v.Name = "Jane"
}
