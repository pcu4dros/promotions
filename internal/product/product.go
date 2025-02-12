package product

type Price struct {
	Original           int
	Final              int
	DiscountPercentage string
}

type Product struct {
	Sku      string
	Name     string
	Category string
	Price    Price
}
