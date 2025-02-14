package http

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"promotions/internal/product"
	"time"
)

func NewServer(
	logger *slog.Logger,
	ctx context.Context,
	pservice *product.Service,
) http.Handler {
	mux := http.NewServeMux()
	addRoutes(
		mux,
		ctx,
		logger,
		pservice,
	)
	var handler http.Handler = mux
	return handler
}

func headers(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func Run(mux http.Handler, log *slog.Logger) {
	log.Info("Starting the server ...")

	s := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	addr := "http://localhost:8080"
	log.Info("Server started", slog.String("address", addr))

	s.ListenAndServe()
}
