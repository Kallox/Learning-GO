package main

import (
	"errors"
	"fmt"

	"github.com/kallox/learning-go/errors/operator"
)

func main() {
	var err error                        // Su valor por defecto es nil
	err = errors.New("Error de ejemplo") // Se le asigna un error
	fmt.Println(err)
	fmt.Println(err.Error()) // El metodo Error retorna el mensaje del error en tipo string

	err2 := fmt.Errorf("Error con formato, string: %s, int: %d", "ERROR", 10) // Crea un error con formato
	fmt.Println(err2)
	fmt.Println(err2.Error())

	x, y := 4, 1

	defer func() { // defer se ejecuta al final de la funcion
		fmt.Println("In main defer")
		r := recover() // recover se usa para recuperar el control de la ejecucion de un programa
		if r != nil {
			fmt.Println("Error en main:", r)
		}
	}() // Se llama sola

	fmt.Println("Dividiendo x e y:", x/y)

	fmt.Println(operator.Div(4, 0)) // panic: Division por cero
}
