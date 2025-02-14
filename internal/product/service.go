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

func (s *Service) List(ctx context.Context) ([]Product, error) {
	products, err := s.repo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("obtaining the products from the store : %v", err)
	}
	return products, nil
}

func (s *Service) GetDiscountRules(ctx context.Context) ([]DiscountRule, error) {
	drules, err := s.repo.GetDiscountRules(ctx)
	if err != nil {
		return nil, fmt.Errorf("obtaining the discount rules from the store : %v", err)
	}
	return drules, nil
}
