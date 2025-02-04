package main

import (
	"fmt"

	"github.com/kallox/learning-go/funciones/function"
)

func main() {
	sum := function.Add(1, 2)
	fmt.Println(sum)

	v, err := function.Calc(function.DIV, 1, 1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(v)
	}

	fmt.Println(function.Split(100))

	fmt.Println(function.MSum(1, 2, 3, 4, 5))

}
