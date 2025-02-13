package product

import (
	"context"
	"database/sql"
	"errors"
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

func (s *SQLiteRepository) List(ctx context.Context) ([]Product, error) {
	rows, err := s.db.QueryContext(ctx,
		`SELECT sku, name, category, price FROM products`)
	if err != nil {
		return nil, fmt.Errorf("obtaining the products from the DB: %w", err)
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var sku, name, category string
		var price int

		if err := rows.Scan(&sku, &name, &category, &price); err != nil {
			return nil, fmt.Errorf("scanning a product row: %w", err)
		}
		p := NewProduct(sku, name, category, price)
		products = append(products, *p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating the rows: %w", err)
	}

	return products, nil
}

func (s *SQLiteRepository) InitProducts(ctx context.Context) error {
	r, err := s.db.ExecContext(ctx,
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
	if i, err := r.RowsAffected(); err != nil || i > 0 {
		return errors.New("rows are affected")
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
		`, p.Sku(), p.Name(), p.Category(), p.Price())
		if err != nil {
			return fmt.Errorf("insert failed for %s: %v", p.Sku(), err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit failed: %v", err)
	}

	return nil
}
