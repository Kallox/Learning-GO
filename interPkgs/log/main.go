// Este paquete es fundamental para mostrar mensajes en la consola.
// Permite guardar mensajes con timestamp en un archivo de texto o en la consola.
// Fundamental para mostrar errores y mensajes de depuración.

package main

import (
	"log"
	"os"
)

func main() {
	log.Println("Hello, World!") // Mensaje en consola
	log.Print("Hello, World!")   // Mensaje en consola sin salto de línea
	// log.Fatal("Goodbye, World!") // Mensaje en consola y termina la ejecución del programa exit status 1
	// log.Panic("Panic, World!")   // Mensaje en consola y termina la ejecución del programa con un panic

	file, err := os.Create("logs.log") // Crear un archivo de texto
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(file, "Success: ", log.LstdFlags) // Logger que guarda todo en el archivo file, con el prefijo "Success: " y la fecha y hora
	logger.Println("Hello, World!")                     // Mensaje en archivo de texto

	logger.SetPrefix("Error: ") // Cambiar el prefijo
	logger.Println("Ha ocurrido un error D:")
}
