package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println // alias

	now := time.Now() // current time
	p(now)            // Uso del alias

	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC) // Creación de una fecha
	p(then)
	then = then.Add(time.Hour * 1) // Suma de una hora... cambiar el 1 por N para sumar N horas. Puede ser Negativo para restar
	p(then)
	// Se pueden obtener las partes de la fecha
	p(then.Year())       // Año
	p(then.Month())      // Mes
	p(then.Day())        // Día
	p(then.Hour())       // Hora
	p(then.Minute())     // Minuto
	p(then.Second())     // Segundo
	p(then.Nanosecond()) // Nanosegundo
	p(then.Location())   // Zona horaria
	p(then.Weekday())    // Día de la semana

	// Comparación de fechas
	p(then.Before(now)) // then < now
	p(then.After(now))  // then > now
	p(then.Equal(now))  // then == now

	// Diferencia de fechas
	diff := now.Sub(then) // now - then
	p(diff)               // Puedes extraer los valores de la diferencia de horas, minutos o segundos como en el ejemplo de la línea 22 a 25
}
