package product

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) SQLiteRepository {
	return SQLiteRepository{
		db: db,
	}
}

func (s *SQLiteRepository) queryProducts(ctx context.Context, query string, args ...any) ([]Product, error) {
	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query products: %w", err)
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.sku, &p.name, &p.category, &p.price); err != nil {
			return nil, fmt.Errorf("failed to scan product row: %w", err)
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return products, nil
}

func (s *SQLiteRepository) List(ctx context.Context) ([]Product, error) {
	return s.queryProducts(ctx, `SELECT sku, name, category, price FROM products`, nil)
}

func (s *SQLiteRepository) ListByPriceRange(ctx context.Context, min, max int) ([]Product, error) {
	return s.queryProducts(ctx,
		`SELECT sku, name, category, price FROM products WHERE price BETWEEN ? AND ?`,
		min, max,
	)
}

func (s *SQLiteRepository) ListByCategory(ctx context.Context, category string) ([]Product, error) {
	return s.queryProducts(ctx,
		`SELECT sku, name, category, price FROM products WHERE category = ?`,
		category,
	)
}

func (s *SQLiteRepository) GetDiscountRules(ctx context.Context) ([]DiscountRule, error) {
	rows, err := s.db.QueryContext(ctx,
		`SELECT field, value, discount FROM discountRules`)
	if err != nil {
		return nil, fmt.Errorf("failed to query discount rules: %w", err)
	}
	var drules []DiscountRule

	for rows.Next() {
		var dr DiscountRule
		if err := rows.Scan(&dr.field, &dr.value, &dr.discount); err != nil {
			return nil, fmt.Errorf("failed to scan discount rule row: %w", err)
		}
		drules = append(drules, dr)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return drules, nil
}
