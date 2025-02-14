package product

import (
	"context"
	"fmt"
)

func (s *SQLiteRepository) InitProducts(ctx context.Context) error {
	_, err := s.db.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS products (
			sku TEXT NOT NULL,
			name TEXT NOT NULL,
			category TEXT NOT NULL,
			price INTEGER NOT NULL
    );`,
	)
	if err != nil {
		return fmt.Errorf("products table creation failed: %v", err)
	}
	return nil
}

func (s *SQLiteRepository) SeedProducts(ctx context.Context) error {
	sampleProducts := []Product{
		*NewProduct("000001", "BV Lean leather ankle boots", "boots", 89000),
		*NewProduct("000002", "BV Lean leather ankle boots", "boots", 99000),
		*NewProduct("000003", "Ashlington leather ankle boots", "boots", 71000),
		*NewProduct("000004", "Naima Embellished leather suede sandals", "sandals", 79500),
		*NewProduct("000005", "Nathane leather ankle boots", "boots", 59000),
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("transaction failed: %v", err)
	}
	defer tx.Rollback()
	for _, p := range sampleProducts {
		_, err = tx.Exec(`
			INSERT OR IGNORE INTO products(sku, name, category, price)
			VALUES (?, ?, ?, ?);
		`, p.sku, p.name, p.category, p.price)
		if err != nil {
			return fmt.Errorf("insert failed for %s: %v", p.sku, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit failed: %v", err)
	}

	return nil
}

func (s *SQLiteRepository) InitDiscountRules(ctx context.Context) error {
	_, err := s.db.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS discountRules (
			field TEXT NOT NULL,
			value TEXT NOT NULL,
			discount INTEGER NOT NULL
    );`,
	)
	if err != nil {
		return fmt.Errorf("DiscountRule table creation failed: %v", err)
	}
	return nil
}

func (s *SQLiteRepository) SeedDRules(ctx context.Context) error {
	drules := []DiscountRule{
		*NewDiscountRule("category", "boots", 30),
		*NewDiscountRule("sku", "000003", 15),
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("transaction failed: %v", err)
	}
	defer tx.Rollback()

	for _, dr := range drules {
		_, err = tx.Exec(`
			INSERT OR IGNORE INTO discountRules(field, value, discount)
			VALUES (?, ?, ?);
		`, dr.field, dr.value, dr.discount)
		if err != nil {
			return fmt.Errorf("insert failed for rule %v: %v", dr, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit failed: %v", err)
	}
	return nil
}
