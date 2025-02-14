package product

import (
	"context"
	"errors"
	"testing"
)

type mockRepo struct {
	products      []Product
	drules        []DiscountRule
	errorOnList   error
	errorOnDrules error
}

func (m *mockRepo) GetDiscountRules(ctx context.Context) ([]DiscountRule, error) {
	if m.errorOnDrules != nil {
		return nil, m.errorOnDrules
	}
	return m.drules, nil
}

func (m *mockRepo) List(ctx context.Context) ([]Product, error) {
	if m.errorOnList != nil {
		return nil, m.errorOnList
	}
	return m.products, nil
}

func (m *mockRepo) ListByCategory(ctx context.Context, category string) ([]Product, error) {
	if m.errorOnList != nil {
		return nil, m.errorOnList
	}
	var filtered []Product
	for _, p := range m.products {
		if p.category == category {
			filtered = append(filtered, p)
		}
	}
	return filtered, nil
}

func (m *mockRepo) ListByPriceRange(ctx context.Context, min, max int) ([]Product, error) {
	if m.errorOnList != nil {
		return nil, m.errorOnList
	}
	var filtered []Product
	for _, p := range m.products {
		if p.price >= min && p.price <= max {
			filtered = append(filtered, p)
		}
	}
	return filtered, nil
}

func TestService_List(t *testing.T) {
	ctx := context.Background()
	repo := &mockRepo{
		products: []Product{
			{sku: "123", name: "Product A", category: "boots", price: 100},
			{sku: "456", name: "Product B", category: "sandals", price: 50},
		},
		drules: []DiscountRule{},
	}

	service := NewService(repo)

	t.Run("list all products", func(t *testing.T) {
		products, err := service.List(ctx, Filter{})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if len(products) != 2 {
			t.Fatalf("expected 2 products, got %d", len(products))
		}
	})

	t.Run("list by category", func(t *testing.T) {
		products, err := service.List(ctx, Filter{Category: "boots"})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if len(products) != 1 || products[0].sku != "123" {
			t.Fatalf("expected product with SKU 123, got %+v", products)
		}
	})

	t.Run("list by price range", func(t *testing.T) {
		products, err := service.List(ctx, Filter{Price: 60})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if len(products) != 1 || products[0].sku != "456" {
			t.Fatalf("expected product with SKU 456, got %+v", products)
		}
	})

	t.Run("error getting discount rules", func(t *testing.T) {
		repo.errorOnDrules = errors.New("failed to get discount rules")
		defer func() { repo.errorOnDrules = nil }()
		_, err := service.List(ctx, Filter{})
		if err == nil || err.Error() != "obtaining the discount rules from the repo : failed to get discount rules" {
			t.Fatalf("expected discount rules error, got %v", err)
		}
	})

	t.Run("error listing products", func(t *testing.T) {
		repo.errorOnList = errors.New("failed to list products")
		defer func() { repo.errorOnList = nil }()
		_, err := service.List(ctx, Filter{})
		if err == nil || err.Error() != "getting the products: failed to list products" {
			t.Fatalf("expected product list error, got %v", err)
		}
	})
}
