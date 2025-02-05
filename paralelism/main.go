package main

import (
	"fmt"
	"time"
)

func main() {
	go myProcess("A") // Palabra reservada para hacer que se ejecute en otro hilo
	go myProcess("B")

	time.Sleep(20 * time.Second) // Espera 20 segundos para que los procesos terminen

	myFirstChannel := make(chan string) // Crear un canal de tipo string

	go myProcessWithChannel("C", myFirstChannel) // Ejecutar un proceso con canal

	result := <-myFirstChannel // Recibir un mensaje del canal

	fmt.Println("Result: ", result)

	close(myFirstChannel) // Cerrar el canal

}

func myProcessWithChannel(p string, c chan string) {
	i := 0

	for i < 15 {
		time.Sleep(1 * time.Second)
		i++
		fmt.Printf("Process: %s - num: %d\n", p, i)
	}

	c <- p   // Enviar un mensaje al canal
	close(c) // Cerrar el canal
}

func myProcess(p string) {
	i := 0

	for i < 15 {
		time.Sleep(1 * time.Second)
		i++
		fmt.Printf("Process: %s - num: %d\n", p, i)
	}
}
