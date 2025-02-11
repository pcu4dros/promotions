package product

type Price struct {
	Original           int    `json:"original"`
	Final              int    `json:"final"`
	DiscountPercentage string `json:"discount_percentage"`
	Currency           string `json:"currency"`
}

type Product struct {
	Sku      string `json:"sku"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    Price  `json:"price"`
}

func (p *Product) ApplyDiscount(percentage int) *Product {
	// TODO: Apply discounts logic goes here
	return &Product{}
}
