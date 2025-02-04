// Reflect permite inspeccionar tipos, valores y estructuras de datos en tiempo de ejecución.

package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID        int64
	Email     string
	FirstName string
	LastName  string
}

func main() {
	myInt := 42
	myPointer := &myInt
	Examine(myInt)
	fmt.Println("-----")
	Examine(myPointer)
	fmt.Println("-----")
	user := User{1, "example@example.com", "Example", "LastName"}
	Examine(user)
	fmt.Println("-----")
}

func Examine(data interface{}) {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	k := v.Kind()
	fmt.Println("Type:", t)
	fmt.Println("Value:", v)
	fmt.Println("Kind:", k)

	if k == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			value := v.Field(i)
			fmt.Printf("Field Name: %s, Field Type: %s, Field Value: %v\n", field.Name, field.Type, value)
		}
	}
}
