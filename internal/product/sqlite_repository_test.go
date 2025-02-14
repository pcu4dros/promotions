package product

import (
	"context"
	"fmt"
	"promotions/internal/sqlite"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB(t *testing.T) (*SQLiteRepository, context.Context) {
	ctx := context.Background()
	db := sqlite.Connect(":memory:")

	r := NewSQLiteRepository(db)
	err := r.InitProducts(ctx)
	if err != nil {
		t.Fatalf("creating products table = %v", err)
	}
	err = r.InitDiscountRules(ctx)
	if err != nil {
		t.Fatalf("creating discountRules table = %v", err)
	}
	return &r, ctx
}

func TestList(t *testing.T) {
	r, ctx := setupTestDB(t)

	t.Run("list all products", func(t *testing.T) {
		err := r.SeedProducts(ctx)
		if err != nil {
			t.Fatalf("adding sample products = %v", err)
		}
		products, err := r.List(ctx, Filter{})
		if err != nil {
			t.Fatalf("List() error = %v", err)
		}

		if len(products) != 5 {
			t.Errorf("expected 5 products, got %d", len(products))
		}
	})
}

func TestListByPriceRange(t *testing.T) {
	repo, ctx := setupTestDB(t)
	testCases := []struct {
		name      string
		min       int
		max       int
		wantCount int
	}{
		{"range matches two products", 50000, 72000, 2},
		{"exact price match", 59000, 59000, 1},
		{"no matches", 10000, 20000, 0},
		{"inverted range", 60000, 50000, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			products, err := repo.List(ctx, Filter{Price: tc.max, Category: ""})
			if err != nil {
				t.Fatalf("ListByPriceRange() error = %v", err)
			}

			if len(products) != tc.wantCount {
				t.Errorf("expected %d products, got %d", tc.wantCount, len(products))
			}

			for _, p := range products {
				if p.price < tc.min || p.price > tc.max {
					t.Errorf("product %s price %d out of range [%d-%d]",
						p.sku, p.price, tc.min, tc.max)
				}
			}
		})
	}
}

func TestListByCategory(t *testing.T) {
	repo, ctx := setupTestDB(t)

	testCases := []struct {
		name      string
		category  string
		wantCount int
	}{
		{"existing category", "boots", 4},
		{"case-sensitive match", "skirts", 0},
		{"non-existent category", "sandals", 1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			products, err := repo.List(ctx, Filter{Price: 0, Category: tc.category})
			if err != nil {
				t.Fatalf("ListByCategory() error = %v", err)
			}

			if len(products) != tc.wantCount {
				t.Errorf("expected %d products, got %d", tc.wantCount, len(products))
			}

			for _, p := range products {
				if p.category != tc.category {
					t.Errorf("product %s has category %s, expected %s",
						p.sku, p.category, tc.category)
				}
			}
		})
	}
}

func TestGetDiscountRules(t *testing.T) {
	r, ctx := setupTestDB(t)

	t.Run("get all rules", func(t *testing.T) {
		err := r.SeedDRules(ctx)
		if err != nil {
			fmt.Println("adding sample discount rules", "with", err)
		}
		rules, err := r.GetDiscountRules(ctx)
		if err != nil {
			t.Fatalf("GetDiscountRules() error = %v", err)
		}

		if len(rules) != 2 {
			t.Fatalf("expected 2 rules, got %d", len(rules))
		}

		expected := []DiscountRule{
			{field: "category", value: "boots", discount: 30},
			{field: "sku", value: "000003", discount: 15},
		}

		for i, rule := range rules {
			if rule != expected[i] {
				t.Errorf("rule %d mismatch:\ngot %+v\nwant %+v", i, rule, expected[i])
			}
		}
	})
}
