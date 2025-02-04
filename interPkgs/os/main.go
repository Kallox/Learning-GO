package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil { // Si hay un error
		fmt.Println(err)
	}
	data := make([]byte, 100) // Crear un slice de bytes de 100 elementos
	c, err := file.Read(data) // Leer el archivo y guardar la cantidad de bytes leídos en c
	if err != nil {           // Si hay un error
		fmt.Println(err)
		os.Exit(1) // Salir del programa con codigo 1
	}
	fmt.Printf("Leyendo %d bytes: %q\n", c, data[:c]) // Imprimir la cantidad de bytes leídos y los bytes leídos

	// Variables de entorno
	env1 := os.Getenv("HOME")       // Obtener el valor de la variable de entorno HOME
	fmt.Println(env1)               // Imprimir el valor de la variable de entorno HOME
	os.Setenv("HOME", "/home/user") // Cambiar el valor de la variable de entorno HOME
	env2 := os.Getenv("HOME")       // Obtener el valor de la variable de entorno HOME
	fmt.Println(env2)               // Imprimir el valor de la variable de entorno HOME

	env, ok := os.LookupEnv("NOT_EXISTS") // Retorna el valor de la variable de entorno HOME y si existe
	fmt.Println(env, ok)                  // Imprimir el valor de la variable de entorno HOME y si existe

	expandedVar := os.ExpandEnv("HOME${HOME}FueExpandida") // Ocupa el valor de la variable de entorno HOME y lo expande
	fmt.Println(expandedVar)                               // Imprimir la variable de entorno expandida
}
