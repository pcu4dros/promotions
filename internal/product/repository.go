package product

import "context"

type Repository interface {
	List(ctx context.Context) ([]Product, error)
}
