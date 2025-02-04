package main

import "fmt"

func main() {
	sum := 0

	// Similar al for de C, pero sin parentesis
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// For como un while
	sum = 1

	for sum < 1000 {
		sum++
	}
	fmt.Println(sum)

	// For infinito
	// for {
	// }

	// For con range
	// El range produce un par de valores: el indice y el valor
	// Si solo se necesita el valor, se puede omitir el indice
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow { // i es el indice, v es el valor... se puede omitir el indice cambiandolo por un _
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// Recorriendo un map
	myMap := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range myMap {
		fmt.Printf("%s -> %s\n", k, v)
	}

}
