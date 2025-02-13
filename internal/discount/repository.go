package discount

import "context"

type Repository interface {
	List(ctx context.Context) ([]DiscountRule, error)
}
