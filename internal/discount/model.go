package discount

// DiscountRule represents a condition for applying a discount.
type DiscountRule struct {
	field    string // e.g., "category" or "sku"
	value    string // e.g., "boots" or "000003"
	discount int    // e.g., 30 for 30%
}

func NewDiscountRule(field string, value string, discount int) *DiscountRule {
	return &DiscountRule{
		field:    field,
		value:    value,
		discount: discount,
	}
}

func (d *DiscountRule) Field() string {
	return d.field
}

func (d *DiscountRule) Value() string {
	return d.value
}

func (d *DiscountRule) Discount() int {
	return d.discount
}
