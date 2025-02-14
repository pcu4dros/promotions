package product

const DefaultCurrency = "EUR"

// DiscountRule represents a condition for applying a discount.
type DiscountRule struct {
	field    string // e.g., "category" or "sku"
	value    string // e.g., "boots" or "000003"
	discount int    // e.g., 30 for 30%
}

type Filter struct {
	Category string
	Price    int
}

type Product struct {
	sku      string
	name     string
	category string
	price    int
}

type Price struct {
	original int
	final    int
	discount *string
	currency string
}

type EnhancedProduct struct {
	sku      string
	name     string
	category string
	price    Price
}

func NewDiscountRule(field string, value string, discount int) *DiscountRule {
	return &DiscountRule{
		field:    field,
		value:    value,
		discount: discount,
	}
}

func NewProduct(sku string, name string, category string, price int) *Product {
	return &Product{
		sku:      sku,
		name:     name,
		category: category,
		price:    price,
	}
}

func NewEnhancedProduct(sku string, name string, category string, price Price) *EnhancedProduct {
	return &EnhancedProduct{
		sku:      sku,
		name:     name,
		category: category,
		price:    price,
	}
}

func NewPrice(original int, final int, discount string, curr string) *Price {
	if curr == "" {
		curr = DefaultCurrency
	}
	return &Price{
		original: original,
		final:    final,
		discount: &discount,
		currency: curr,
	}
}
