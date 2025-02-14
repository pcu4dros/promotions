package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"promotions/internal/http"
	"promotions/internal/product"
	"promotions/internal/sqlite"
)

func main() {
	ctx := context.Background()
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	db := sqlite.Connect(":memory:")
	pr := product.NewSQLiteRepository(db)

	// This is a helper function to initialize the DB with the tables and some rows
	err := initDB(ctx, pr)
	if err != nil {
		fmt.Println("initializing the DB", "in", err)
		panic(1)
	}
	ps := product.NewService(&pr)

	srv := http.NewServer(logger, ctx, ps)
	http.Run(srv, logger)
}

func initDB(ctx context.Context, pr product.SQLiteRepository) error {
	err := pr.InitProducts(ctx)
	if err != nil {
		fmt.Println("creating products table", "with", err)
	}
	err = pr.InitDiscountRules(ctx)
	if err != nil {
		fmt.Println("creating discountRules table", "with", err)
	}
	err = pr.SeedProducts(ctx)
	if err != nil {
		fmt.Println("adding sample products", "with", err)
	}
	err = pr.SeedDRules(ctx)
	if err != nil {
		fmt.Println("adding sample discount rules", "with", err)
	}
	return nil
}
