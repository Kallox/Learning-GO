package book

type Book struct { // Similar a describir los atributos de una clase
	Title        string
	Author       string
	Pages        int
	privateField string // No se puede acceder a este campo desde otro paquete (Encapsulamiento)
}

type Printable interface { // Similar a definir una interfaz de una clase (Polimorfismo)
	PrintInfo()
}

func Print(p Printable) { // Similar a definir un método de una interfaz de una clase (Polimorfismo)
	p.PrintInfo()
}

func (b *Book) PrintInfo() { // Similar a definir un método de una clase
	println(b.Title)
	println(b.Author)
	println(b.Pages)
}

func NewBook(title, author string, pages int) *Book { // Similar a definir un constructor de una clase
	return &Book{
		Title:  title,
		Author: author,
		Pages:  pages,
	}
}

func (b *Book) GetPrivateField() string { // Similar a definir un getter de una clase (encapsulamiento)
	return b.privateField
}

func (b *Book) SetPrivateField(privateField string) { // Similar a definir un setter de una clase (encapsulamiento)
	b.privateField = privateField
}

// En Go no existe la herencia, pero se puede simular con composición
type TextBook struct {
	Book      // Composición para reutilizar los atributos de Book
	Grade     int
	Editorial string
}

func NewTextBook(title, author string, pages, grade int, editorial string) *TextBook {
	return &TextBook{
		Book: Book{
			Title:  title,
			Author: author,
			Pages:  pages,
		},
		Grade:     grade,
		Editorial: editorial,
	}
}

func (b *TextBook) PrintInfo() { // Similar a definir un método de una clase
	println(b.Title)
	println(b.Author)
	println(b.Pages)
	println(b.Grade)
	println(b.Editorial)
}
