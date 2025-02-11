package main

import (
	"github.com/kallox/learning-go/poo/book"
)

func main() {
	book1 := book.Book{
		Title:  "The Art of Computer Programming",
		Author: "Donald Knuth",
		Pages:  700,
	}
	book1.PrintInfo()

	book2 := book.NewBook("The Go Programming Language", "Alan A. A. Donovan", 380)
	book2.PrintInfo()

	book2.SetPrivateField("Im the private field")
	println(book2.GetPrivateField())

	textBook := book.NewTextBook("Mathematics", "John Doe", 200, 5, "Editorial")
	textBook.PrintInfo()

	// Polimorfismo
	book.Print(book2) // La función Print recibe un Printable, y Book implementa Printable
	book.Print(textBook)
}
