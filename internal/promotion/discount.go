package promotion

import (
	"encoding/json"
	"fmt"
	"io"
	"promotions/internal/product"
)

type Discounter struct {
	category map[string]int
	sku      map[string]int
}

func NewDiscounter(r io.Reader) (*Discounter, error) {
	discounter := &Discounter{
		category: make(map[string]int),
		sku:      make(map[string]int),
	}

	if r != nil {
		var config struct {
			Categories map[string]int `json:"category_discounts"`
			SKUs       map[string]int `json:"sku_discounts"`
		}

		if err := json.NewDecoder(r).Decode(&config); err != nil {
			return nil, err
		}

		discounter.category = config.Categories
		discounter.sku = config.SKUs
	}

	return discounter, nil
}

func (d *Discounter) getDiscount(p *product.Product) int {
	var percentage int
	if value, exists := d.category[p.Category]; exists && value > percentage {
		percentage = value
	}
	if value, exists := d.sku[p.Sku]; exists && value > percentage {
		percentage = value
	}
	return percentage
}

// Apply a discount percentage to an existing Prodcut
func (d *Discounter) ApplyDiscount(p *product.Product) *product.Product {
	discount := d.getDiscount(p)
	p.Price.Final = p.Price.Original - (p.Price.Original*discount)/100
	if discount > 0 {
		p.Price.DiscountPercentage = fmt.Sprintf("%d%%", discount)
	}
	return p
}
