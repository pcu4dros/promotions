package main

import (
	"context"
	"fmt"
	"promotions/internal/discount"
	"promotions/internal/product"
	"promotions/internal/sqlite"
)

func main() {
	ctx := context.Background()

	db := sqlite.Connect(":memory:")

	pr := product.NewSQLiteRepository(db)
	dr := discount.NewSQLiteRepository(db)

	err := initDB(ctx, pr, dr)
	if err != nil {
		fmt.Println("initializing the DB", "in", err)
		panic(1)
	}

	ps := product.NewService(&pr)
	ds := discount.NewService(&dr)

	products, err := ps.List(ctx)
	if err != nil {
		fmt.Println("obtaining products from DB", "in", err)
		panic(1)
	}
	fmt.Println(products)
	drules, err := ds.List(ctx)
	if err != nil {
		fmt.Println("obtaining discount rules from DB", "in", err)
		panic(1)
	}
	fmt.Println(drules)
}

func initDB(ctx context.Context, pr product.SQLiteRepository, dr discount.SQLiteRepository) error {
	err := pr.InitProducts(ctx)
	if err != nil {
		fmt.Println("creating products table", "with", err)
	}
	err = dr.InitDiscountRules(ctx)
	if err != nil {
		fmt.Println("creating discountRules table", "with", err)
	}
	err = pr.SeedProducts(ctx)
	if err != nil {
		fmt.Println("adding sample products", "with", err)
	}
	err = dr.SeedDRules(ctx)
	if err != nil {
		fmt.Println("adding sample discount rules", "with", err)
	}
	return nil
}
