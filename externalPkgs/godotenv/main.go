package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
	fmt.Println(os.Getenv("MY_ENV1"))
	fmt.Println(os.Getenv("MY_ENV2"))
	fmt.Println(os.Getenv("MY_ENV3"))

	myEnv, err := godotenv.Read() // Lee las variables y las guarda en un mapa
	if err != nil {
		panic("Error reading .env file")
	}
	fmt.Println(myEnv)
}
