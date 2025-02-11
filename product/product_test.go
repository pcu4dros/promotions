package product

import (
	"testing"
)

func TestDiscounts(t *testing.T) {
	product := &Product{
		Sku:      "000002",
		Name:     "Dummy product",
		Category: "sandals",
		Price: Price{
			Original:           89000,
			Final:              89000,
			DiscountPercentage: "",
			Currency:           "EUR",
		},
	}

	t.Run("returns a Product without discount", func(t *testing.T) {
		discount := product.CalculateDiscount()
		product.Price.ApplyDiscount(discount)

		assertPrice(t, product.Price.Original, 89000)
		assertPrice(t, product.Price.Final, 89000)
		assertDiscountValue(t, product.Price.DiscountPercentage, "")
	})

	t.Run("returns a Product with 30% discount", func(t *testing.T) {
		product.Category = "boots"
		discount := product.CalculateDiscount()
		product.Price.ApplyDiscount(discount)

		assertPrice(t, product.Price.Original, 89000)
		assertPrice(t, product.Price.Final, 62300)
		assertDiscountValue(t, product.Price.DiscountPercentage, "30%")
	})

	t.Run("returns a Product with 15% discount", func(t *testing.T) {
		product.Sku = "000003"
		product.Category = "sandals"
		discount := product.CalculateDiscount()
		product.Price.ApplyDiscount(discount)

		assertPrice(t, product.Price.Original, 89000)
		assertPrice(t, product.Price.Final, 75650)
		assertDiscountValue(t, product.Price.DiscountPercentage, "15%")
	})

	// When multiple discounts collide, the bigger discount must be applied.
	t.Run("returns a product with 30% discount, collision (bigger is applied)", func(t *testing.T) {
		product.Sku = "000003"
		product.Category = "boots"
		discount := product.CalculateDiscount()
		product.Price.ApplyDiscount(discount)

		assertPrice(t, product.Price.Original, 89000)
		assertPrice(t, product.Price.Final, 62300)
		assertDiscountValue(t, product.Price.DiscountPercentage, "30%")
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
