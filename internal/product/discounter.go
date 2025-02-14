package product

import (
	"fmt"
)

type Discounter struct {
	Rules []DiscountRule
}

func (d *Discounter) getDiscount(p *Product) (int, string) {
	var discounts []int
	var discount int

	for _, rule := range d.Rules {
		switch rule.field {
		case "category":
			if p.category == rule.value {
				discounts = append(discounts, rule.discount)
			}
		case "sku":
			if p.sku == rule.value {
				discounts = append(discounts, rule.discount)
			}
		}
	}
	if getMaxDiscount(discounts) > 0 {
		discount := getMaxDiscount(discounts)
		return getMaxDiscount(discounts), fmt.Sprintf("%d%%", discount)
	}
	return discount, ""
}

// Apply a discount percentage to an existing Price
func (d *Discounter) ApplyDiscount(p *Product) *Price {
	discount, percentage := d.getDiscount(p)
	final := p.price - (p.price*discount)/100
	price := NewPrice(p.price, final, percentage, "")
	return price
}

// getMaxDiscount returns the highest discount.
func getMaxDiscount(discounts []int) int {
	var max int
	for _, d := range discounts {
		if d > max {
			max = d
		}
	}
	return max
}
