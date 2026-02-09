package handlers

import (
	"encoding/json"
	"net/http"

	"final/internal/middleware"
	"final/internal/store"
)

type UsersHandler struct {
	Store store.Store
}

func (h *UsersHandler) Me(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value(middleware.ContextUserID)
	if uid == nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized"})
		return
	}
	userID, ok := uid.(int)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized"})
		return
	}
	u, found := h.Store.GetUserByID(userID)
	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(u)
}

func (h *UsersHandler) Users(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	json.NewEncoder(w).Encode(h.Store.ListUsers())
}
