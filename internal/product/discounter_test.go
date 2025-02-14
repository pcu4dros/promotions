package product

import (
	"testing"
)

func TestDiscounter(t *testing.T) {
	var discountRules []DiscountRule

	categoryRule := NewDiscountRule("category", "boots", 30)
	skuRule := NewDiscountRule("sku", "000003", 15)

	discountRules = append(discountRules, *categoryRule, *skuRule)

	discounter := Discounter{
		Rules: discountRules,
	}

	t.Run("returns a Product without discount", func(t *testing.T) {
		p := NewProduct("000002", "Dummy product", "sandals", 89000)
		price := discounter.ApplyDiscount(p)

		assertPrice(t, price.final, p.price)
		assertDiscountValue(t, *price.discount, "")
	})

	t.Run("returns a Product with 30% discount", func(t *testing.T) {
		pdiscount := NewProduct("000002", "Dummy product", "boots", 89000)
		price := discounter.ApplyDiscount(pdiscount)

		assertPrice(t, price.original, 89000)
		assertPrice(t, price.final, 62300)
		assertDiscountValue(t, *price.discount, "30%")
	})

	t.Run("returns a Product with 15% discount", func(t *testing.T) {
		p := NewProduct("000003", "Dummy product", "sandals", 89000)
		price := discounter.ApplyDiscount(p)

		assertPrice(t, price.original, 89000)
		assertPrice(t, price.final, 75650)
		assertDiscountValue(t, *price.discount, "15%")
	})

	// When multiple discounts collide, the bigger discount must be applied.
	t.Run("returns a product with 30% discount, collision (bigger is applied)", func(t *testing.T) {
		p := NewProduct("000003", "Dummy product", "boots", 89000)
		price := discounter.ApplyDiscount(p)

		assertPrice(t, price.original, 89000)
		assertPrice(t, price.final, 62300)
		assertDiscountValue(t, *price.discount, "30%")
	})
}

func assertPrice(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get the correct price, got %d want %d", got, want)
	}
}

func assertDiscountValue(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("did not get the correct discount_percentage value, got %s want %s", got, want)
	}
}
