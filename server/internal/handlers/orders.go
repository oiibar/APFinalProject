package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"final/internal/models"
	"final/internal/store"
)

type OrdersHandler struct {
	Store      *store.MemoryStore
	OrderQueue chan int
}

func (h *OrdersHandler) Create(w http.ResponseWriter, r *http.Request) {
	var o models.Order
	json.NewDecoder(r.Body).Decode(&o)
	created := h.Store.CreateOrder(o)

	h.OrderQueue <- created.ID

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (h *OrdersHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	o, ok := h.Store.GetOrder(id)
	if !ok {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(o)
}
