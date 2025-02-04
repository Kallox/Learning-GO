package operator

func Div(x, y float64) float64 {
	defer func() {
		if y <= 0 {
			panic("Division por cero") // panic se usa para detener la ejecucion de un programa y mostrar un mensaje de error
		}
	}()
	return x / y
}
