package discount

import (
	"promotions/internal/product"
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
		p := product.NewProduct("000002", "Dummy product", "sandals", 89000)
		discount, percentage := discounter.GetDiscount(p)

		assertPrice(t, discount, 0)
		assertDiscountValue(t, percentage, "")
	})

	t.Run("returns a Product with 30% discount", func(t *testing.T) {
		p := product.NewProduct("000002", "Dummy product", "boots", 89000)

		discount, percentage := discounter.GetDiscount(p)
		final := discounter.ApplyDiscount(discount, p)

		assertPrice(t, p.Price(), 89000)
		assertPrice(t, final, 62300)
		assertDiscountValue(t, percentage, "30%")
	})

	t.Run("returns a Product with 15% discount", func(t *testing.T) {
		p := product.NewProduct("000003", "Dummy product", "sandals", 89000)

		discount, percentage := discounter.GetDiscount(p)
		final := discounter.ApplyDiscount(discount, p)

		assertPrice(t, p.Price(), 89000)
		assertPrice(t, final, 75650)
		assertDiscountValue(t, percentage, "15%")
	})

	// When multiple discounts collide, the bigger discount must be applied.
	t.Run("returns a product with 30% discount, collision (bigger is applied)", func(t *testing.T) {
		p := product.NewProduct("000003", "Dummy product", "boots", 89000)

		discount, percentage := discounter.GetDiscount(p)
		final := discounter.ApplyDiscount(discount, p)

		assertPrice(t, p.Price(), 89000)
		assertPrice(t, final, 62300)
		assertDiscountValue(t, percentage, "30%")
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
