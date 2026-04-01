package order

import "time"

// Order representa um pedido no sistema
type Order struct {
	ID        string    `json:"id"`
	Customer  string    `json:"customer"`
	Product   string    `json:"product"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateOrderRequest representa a requisição para criar um pedido
type CreateOrderRequest struct {
	Customer string  `json:"customer"`
	Product  string  `json:"product"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}
