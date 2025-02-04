package vehicles

import "fmt"

type Vehicle interface { // Interface
	Distance() float64
}

const ( // Constantes
	CarVehicle   = "Car"
	TruckVehicle = "Truck"
)

func New(v string, time int) (Vehicle, error) { // Factory para polimorfismo
	switch v {
	case CarVehicle:
		return &Car{Time: time}, nil
	case TruckVehicle:
		return &Truck{Time: time}, nil
	default:
		return nil, fmt.Errorf("Vehicle %s not found", v)
	}

}

type Car struct { // Structs
	Time int
}

type Truck struct { // Structs
	Time int
}

// Las funciones se llaman igual pero pertenecen a diferentes Structs
func (c *Car) Distance() float64 { // Struct method de Car
	return 100 * (float64(c.Time) / 60)
}

func (t *Truck) Distance() float64 { // Struct method de Truck
	return 70 * (float64(t.Time) / 60)
}
