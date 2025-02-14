package product

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

type PriceResponse struct {
	Original int     `json:"original"`
	Final    int     `json:"final"`
	Discount *string `json:"discount_percentage"`
	Currency string  `json:"currency"`
}

type ProductResponse struct {
	Sku      string        `json:"sku"`
	Name     string        `json:"name"`
	Category string        `json:"category"`
	Price    PriceResponse `json:"price"`
}

func ConvertToProductResponse(products []EnhancedProduct) map[string][]ProductResponse {
	var productResponses []ProductResponse
	for _, p := range products {
		var discount *string
		if p.price.discount != nil && *p.price.discount != "" { // Prevent empty string
			discount = p.price.discount
		} else {
			discount = nil
		}
		productResponse := ProductResponse{
			Sku:      p.sku,
			Name:     p.name,
			Category: p.category,
			Price: PriceResponse{
				Original: p.price.original,
				Final:    p.price.final,
				Discount: discount,
				Currency: p.price.currency,
			},
		}
		productResponses = append(productResponses, productResponse)
	}

	return map[string][]ProductResponse{"products": productResponses}
}

func HandleProduct(ctx context.Context, log *slog.Logger, pservice Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		category := r.URL.Query().Get("category")
		priceLessThan := r.URL.Query().Get("priceLessThan")

		var price int
		if priceLessThan != "" {
			priceLessThan, err := strconv.Atoi(priceLessThan)
			price = priceLessThan
			if err != nil {
				log.Error("invalid priceLessThan filter", "with: ", err)
			}
		}

		if price < 0 {
			http.Error(w, "Invalid priceLessThan filter", http.StatusBadRequest)
			return
		}

		log.Info("price ", "with: ", strconv.Itoa(price))

		log.Info("category ", "with: ", category)
		filter := Filter{
			Category: category,
			Price:    price,
		}

		products, err := pservice.List(ctx, filter)
		if err != nil {
			log.Error("obtaining the products", "with", err)
			http.Error(w, "Error fetching products", http.StatusInternalServerError)
			return
		}

		response := ConvertToProductResponse(products)

		marshaled, err := json.MarshalIndent(response, "", "   ")
		if err != nil {
			log.Error("marshaling reponse", "with", err)
			http.Error(w, "obtaining the product response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(marshaled)
	})
}
