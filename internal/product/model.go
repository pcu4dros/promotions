package product

type Product struct {
	sku      string
	name     string
	category string
	price    int
}

func NewProduct(sku string, name string, category string, price int) *Product {
	return &Product{
		sku:      sku,
		name:     name,
		category: category,
		price:    price,
	}
}

func (p *Product) Sku() string {
	return p.sku
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) Category() string {
	return p.category
}

func (p *Product) Price() int {
	return p.price
}
