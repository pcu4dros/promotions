package http

import (
	"context"
	"log/slog"
	"net/http"
	"promotions/internal/product"
)

func addRoutes(mux *http.ServeMux, ctx context.Context, logger *slog.Logger, pservice *product.Service) {
	mux.Handle("/products", product.HandleProduct(ctx, logger, *pservice))
}
