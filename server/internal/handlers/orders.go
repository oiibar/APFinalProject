package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"final/internal/models"
	"final/internal/store"
)

type OrdersHandler struct {
	Store      store.Store
	OrderQueue chan int
}

func (h *OrdersHandler) Orders(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(h.Store.ListOrders())
	case http.MethodPost:
		var o models.Order
		json.NewDecoder(r.Body).Decode(&o)
		created := h.Store.CreateOrder(o)
		h.OrderQueue <- created.ID
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(created)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *OrdersHandler) OrderByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	switch r.Method {
	case http.MethodGet:
		o, ok := h.Store.GetOrder(id)
		if !ok {
			http.NotFound(w, r)
			return
		}
		json.NewEncoder(w).Encode(o)
	case http.MethodPut:
		var o models.Order
		json.NewDecoder(r.Body).Decode(&o)
		o.ID = id
		updated, err := h.Store.UpdateOrder(o)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(updated)
	case http.MethodDelete:
		err := h.Store.DeleteOrder(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
