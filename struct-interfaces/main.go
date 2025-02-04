package main

import (
	"encoding/json"
	"fmt"

	"github.com/kallox/learning-go/struct-interfaces/structs"
	"github.com/kallox/learning-go/struct-interfaces/vehicles"
)

func main() {
	var p1 structs.Product // Se crea una instancia de Product
	p1.ID = 1              // Se asignan valores a los campos
	p1.Name = "Product 1"
	p1.Type = structs.Type{ID: 1, Code: "T1", Desc: "Type 1"} // Se asignan valores al Struct anidado
	p1.Price = 10.5
	p1.Tags = []string{"tag1", "tag2"} // Se asignan valores al Slice
	fmt.Println(p1)
	v, err := json.Marshal(p1)  // Se convierte a JSON
	fmt.Println(string(v), err) // Se imprime el JSON
	fmt.Println("Total price:", p1.TotalPrice())
	p1.SetName("Product 1 Updated")
	fmt.Println(p1.Name)

	fmt.Println("------Vehicles------")

	carV := vehicles.Car{Time: 120} // Se crea una instancia de Car
	fmt.Println(carV.Distance())    // Se llama al método de Car

	truckV, err := vehicles.New("Truck", 100) // Se crea una instancia de Truck
	fmt.Println(truckV.Distance())            // Se llama al método de Truck
}
