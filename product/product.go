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

func (p *Product) CalculateDiscount() int {
	if p.Category == "boots" && p.Sku == "000003" {
		return 30
	}
	if p.Category == "boots" {
		return 30
	}
	if p.Sku == "000003" {
		return 15
	}
	return 0
}

func (p *Price) ApplyDiscount(percentage int) {
	p.Final = p.Original - (p.Original*percentage)/100
	switch percentage {
	case 30:
		p.DiscountPercentage = "30%"
	case 15:
		p.DiscountPercentage = "15%"
	}
}
