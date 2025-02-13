package discount

import (
	"fmt"
	"promotions/internal/product"
)

type Discounter struct {
	Rules []DiscountRule
}

func (d *Discounter) GetDiscount(p *product.Product) (int, string) {
	var discounts []int
	var discount int

	for _, rule := range d.Rules {
		switch rule.field {
		case "category":
			if p.Category() == rule.value {
				discounts = append(discounts, rule.discount)
			}
		case "sku":
			if p.Sku() == rule.value {
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

// Apply a discount percentage to an existing Product
func (d *Discounter) ApplyDiscount(discount int, p *product.Product) int {
	discount = p.Price() - (p.Price()*discount)/100
	return discount
}

// getMaxDiscount returns the highest discount percentage.
func getMaxDiscount(discounts []int) int {
	var max int
	for _, d := range discounts {
		if d > max {
			max = d
		}
	}
	return max
}
