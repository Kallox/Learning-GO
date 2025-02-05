package main

import (
	"fmt"

	"github.com/google/uuid" // Importacion del paquete uuid
)

func main() {
	id1 := uuid.New().String()

	fmt.Println(id1)
}
