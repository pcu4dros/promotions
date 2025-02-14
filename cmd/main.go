package main

import (
	"context"
	"fmt"
	"promotions/internal/product"
	"promotions/internal/sqlite"
)

func main() {
	ctx := context.Background()

	db := sqlite.Connect(":memory:")
	pr := product.NewSQLiteRepository(db)

	err := initDB(ctx, pr)
	if err != nil {
		fmt.Println("initializing the DB", "in", err)
		panic(1)
	}

	ps := product.NewService(&pr)

	products, err := ps.List(ctx)
	if err != nil {
		fmt.Println("obtaining products from DB", "in", err)
		panic(1)
	}
	fmt.Println(products)
	drules, err := ps.GetDiscountRules(ctx)
	if err != nil {
		fmt.Println("obtaining discount rules from DB", "in", err)
		panic(1)
	}
	fmt.Println(drules)
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
