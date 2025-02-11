package product

import "testing"

func TestApplyDiscounts(t *testing.T) {
	product := &Product{
		Sku:      "000003",
		Name:     "BV lean leather ankle boots",
		Category: "boots",
		Price: Price{
			Original:           89000,
			Final:              89000,
			DiscountPercentage: "",
			Currency:           "EUR",
		},
	}

	productWithDiscount := product
	productWithDiscount.Price.Final = 62300

	discount := 30

	want := productWithDiscount

	got := product.ApplyDiscount(discount)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
