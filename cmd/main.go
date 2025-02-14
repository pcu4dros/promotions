package main

import (
	"context"
	"log/slog"
	"os"
	"promotions/internal/http"
	"promotions/internal/product"
	"promotions/internal/sqlite"
)

func main() {
	ctx := context.Background()
	opts := &slog.HandlerOptions{Level: slog.LevelInfo}
	logger := slog.New(slog.NewTextHandler(os.Stderr, opts))

	db := sqlite.Connect(":memory:")
	pr := product.NewSQLiteRepository(db)

	// This is a helper function to initialize the DB with the tables and some rows
	err := initDB(ctx, logger, pr)
	if err != nil {
		logger.Error("initializing the DB", "error", err)
		panic(1)
	}
	ps := product.NewService(&pr)

	// Initialize a simple http.ServeMux server
	srv := http.NewServer(logger, ctx, ps)
	http.Run(srv, logger)
}

func initDB(ctx context.Context, log *slog.Logger, pr product.SQLiteRepository) error {
	log.Info("::::::: creating the products table ::::::")
	err := pr.InitProducts(ctx)
	if err != nil {
		log.Error("creating products table", "error", err)
		panic(0)
	}
	log.Info("::::::: creating the discountRules table ::::::")

	err = pr.InitDiscountRules(ctx)
	if err != nil {
		log.Error("creating discountRules table", "error", err)
		panic(0)
	}
	log.Info("::::::: adding sample products to table ::::::")

	err = pr.SeedProducts(ctx)
	if err != nil {
		log.Error("adding sample products", "error", err)
		panic(0)
	}

	log.Info("::::::: adding sample discount rules to the table ::::::")
	err = pr.SeedDRules(ctx)
	if err != nil {
		log.Error("adding sample discount rules", "error", err)
		panic(0)
	}
	return nil
}
