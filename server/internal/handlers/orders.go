package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"final/internal/middleware"
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
		// Admin can see all orders, regular users only their own.
		role := strings.ToLower(middleware.GetUserRole(r))
		if role == "admin" {
			writeJSON(w, http.StatusOK, h.Store.ListOrders())
			return
		}
		uid, ok := middleware.GetUserID(r)
		if !ok {
			writeError(w, http.StatusUnauthorized, "unauthorized")
			return
		}
		writeJSON(w, http.StatusOK, h.Store.ListOrdersByUser(uid))
	case http.MethodPost:
		var o models.Order
		if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
			writeError(w, http.StatusBadRequest, "invalid json")
			return
		}
		uid, ok := middleware.GetUserID(r)
		if !ok {
			writeError(w, http.StatusUnauthorized, "unauthorized")
			return
		}
		if len(o.Items) == 0 {
			writeError(w, http.StatusBadRequest, "order must contain items")
			return
		}
		for _, it := range o.Items {
			if it.BookID <= 0 || it.Quantity <= 0 {
				writeError(w, http.StatusBadRequest, "invalid order items")
				return
			}
		}
		o.UserID = uid
		created := h.Store.CreateOrder(o)
		h.OrderQueue <- created.ID
		writeJSON(w, http.StatusCreated, created)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *OrdersHandler) OrderByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}
	switch r.Method {
	case http.MethodGet:
		o, ok := h.Store.GetOrder(id)
		if !ok {
			writeError(w, http.StatusNotFound, "order not found")
			return
		}
		role := strings.ToLower(middleware.GetUserRole(r))
		if role != "admin" {
			uid, ok := middleware.GetUserID(r)
			if !ok || o.UserID != uid {
				writeError(w, http.StatusForbidden, "forbidden")
				return
			}
		}
		writeJSON(w, http.StatusOK, o)
	case http.MethodPut:
		var o models.Order
		if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
			writeError(w, http.StatusBadRequest, "invalid json")
			return
		}
		o.ID = id
		updated, err := h.Store.UpdateOrder(o)
		if err != nil {
			writeError(w, http.StatusNotFound, "order not found")
			return
		}
		writeJSON(w, http.StatusOK, updated)
	case http.MethodDelete:
		err := h.Store.DeleteOrder(id)
		if err != nil {
			writeError(w, http.StatusNotFound, "order not found")
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
