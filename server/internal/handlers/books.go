package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"final/internal/models"
	"final/internal/store"
)

type BooksHandler struct {
	Store *store.MemoryStore
}

func (h *BooksHandler) Books(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(h.Store.ListBooks())
		return
	}

	if r.Method == http.MethodPost {
		var b models.Book
		json.NewDecoder(r.Body).Decode(&b)
		created := h.Store.CreateBook(b)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(created)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (h *BooksHandler) BookByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	b, ok := h.Store.GetBook(id)
	if !ok {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(b)
}
