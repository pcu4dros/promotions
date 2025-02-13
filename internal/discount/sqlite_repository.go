package discount

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

func (s *SQLiteRepository) InitDiscountRules(ctx context.Context) error {
	rd, err := s.db.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS discountRules (
			field TEXT NOT NULL,
			value TEXT NOT NULL,
			discount INTEGER NOT NULL
    );`,
	)
	if err != nil {
		return fmt.Errorf("DiscountRule table creation failed: %v", err)
	}
	if i, err := rd.RowsAffected(); err != nil || i > 0 {
		return errors.New("rows are affected")
	}
	return nil
}

func (s *SQLiteRepository) List(ctx context.Context) ([]DiscountRule, error) {
	rows, err := s.db.QueryContext(ctx,
		`SELECT field, value, discount FROM discountRules`)
	if err != nil {
		return nil, fmt.Errorf("failed to query discount rules: %w", err)
	}
	var drules []DiscountRule

	for rows.Next() {
		var field, value string
		var disc int

		if err := rows.Scan(&field, &value, &disc); err != nil {
			return nil, fmt.Errorf("failed to scan discount rule row: %w", err)
		}
		dr := NewDiscountRule(field, value, disc)
		drules = append(drules, *dr)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return drules, nil
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
		`, dr.Field(), dr.Value(), dr.Discount())
		if err != nil {
			return fmt.Errorf("insert failed for rule %v: %v", dr, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit failed: %v", err)
	}
	return nil
}
