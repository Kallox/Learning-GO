package structs

type Product struct {
	ID    uint     `json:"id"` // Nombre --- Tipo --- Etiqueta JSON
	Name  string   `json:"name"`
	Type  Type     `json:"type"` // Struct anidado
	Price float32  `json:"price"`
	Tags  []string `json:"tags"` // Slice de strings
}

type Type struct { // Struct anidado
	ID   uint   `json:"id"`
	Code string `json:"code"`
	Desc string `json:"desc"`
}

func (p Product) TotalPrice() float64 { // Struct method de Product
	return float64(p.Price)
}

func (p *Product) SetName(name string) { // Para modificar se debe indicar como puntero
	p.Name = name
}
