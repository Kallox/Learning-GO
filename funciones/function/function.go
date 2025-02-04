package function

import (
	"errors"
)

type Operation int // Se crea un tipo de dato Operation que es un int

const (
	SUM Operation = iota // Se le asigna un valor a cada constante... iota es un generador de numeros secuenciales
	SUB
	DIV
	MUL
)

func Add(a, b int) int { // Es publico porque empieza con mayuscula
	return a + b
}

func Calc(op Operation, x, y float64) (float64, error) {
	switch op {
	case SUM:
		return x + y, nil // nil es como null o None
	case SUB:
		return x - y, nil
	case DIV:
		if y == 0 {
			return 0, errors.New("division by zero") // errors.New crea un error
		}
		return x / y, nil
	case MUL:
		return x * y, nil
	default:
		return 0, errors.New("invalid operation")
	}
}

func Split(v int) (x int, y int) { // Se pueden retornar multiples valores. Si se retorna mas de uno se deben poner entre parentesis
	x = v * 4 / 9
	y = v - x
	return // En la declaracion de la funcion se indico que x e y se retornan, por lo que no es necesario ponerlos en el return
}

// Se le pueden pasar multiples parametros de un mismo tipo
func MSum(values ...float64) float64 { // Cuando se ocupa este tipo de parametros se le conoce como variadic, siempre debe ser el ultimo parametro y unico
	var sum float64
	for _, v := range values {
		sum += v
	}
	return sum
}
