package product

import "context"

type Repository interface {
	List(ctx context.Context, filter Filter) ([]Product, error)
	GetDiscountRules(ctx context.Context) ([]DiscountRule, error)
}
