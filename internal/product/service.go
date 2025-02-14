package product

import (
	"context"
	"fmt"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) List(ctx context.Context, filter Filter) ([]EnhancedProduct, error) {
	var drules []DiscountRule

	drules, err := s.repo.GetDiscountRules(ctx)
	discounter := Discounter{
		Rules: drules,
	}
	if err != nil {
		return nil, fmt.Errorf("obtaining the discount rules from the repo : %v", err)
	}

	fmt.Println("filters: ", filter)

	var products []Product
	products, err = s.repo.List(ctx, filter)

	fmt.Println("products: ", products)
	if err != nil {
		return nil, fmt.Errorf("getting the products: %w", err)
	}
	eproducts := getEproducts(products, discounter, filter)
	return eproducts, nil
}

func getEproducts(products []Product, d Discounter, f Filter) []EnhancedProduct {
	eproducts := make([]EnhancedProduct, 0, len(products))
	for _, p := range products {
		var pr *Price
		if f.Price > 0 {
			pr = NewPrice(p.price, p.price, "", "")
		} else {
			pr = d.ApplyDiscount(&p)
		}
		ep := NewEnhancedProduct(p.sku, p.name, p.category, *pr)
		eproducts = append(eproducts, *ep)
	}
	return eproducts
}
