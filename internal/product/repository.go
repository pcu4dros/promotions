package product

import "context"

type Repository interface {
	List(ctx context.Context) ([]Product, error)
	ListByCategory(ctx context.Context, category string) ([]Product, error)
	ListByPriceRange(ctx context.Context, min, max int) ([]Product, error)
	GetDiscountRules(ctx context.Context) ([]DiscountRule, error)
}
