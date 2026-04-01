package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"stream-orders/internal/order"

	"github.com/google/uuid"
)

// OrderHandler gerencia as requisições relacionadas a pedidos
type OrderHandler struct {
	// TODO: Adicionar dependências (Kafka producer, DB, etc)
}

// NewOrderHandler cria uma nova instância do OrderHandler
func NewOrderHandler() *OrderHandler {
	return &OrderHandler{}
}

// CreateOrder lida com a criação de novos pedidos
func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req order.CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validação básica
	if req.Customer == "" || req.Product == "" || req.Quantity <= 0 || req.Price <= 0 {
		http.Error(w, "Invalid order data", http.StatusBadRequest)
		return
	}

	// Criar pedido
	newOrder := order.Order{
		ID:        uuid.New().String(),
		Customer:  req.Customer,
		Product:   req.Product,
		Quantity:  req.Quantity,
		Price:     req.Price,
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	// TODO: Publicar no Kafka
	// TODO: Salvar no banco de dados

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newOrder)
}

// ListOrders lista todos os pedidos
func (h *OrderHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// TODO: Buscar do banco de dados
	orders := []order.Order{}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

// HealthCheck retorna o status da API
func (h *OrderHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "healthy",
		"time":   time.Now().Format(time.RFC3339),
	})
}
