package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"final/internal/models"
	"final/internal/store"
)

type BooksHandler struct {
	Store store.Store
}

func (h *BooksHandler) Books(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(h.Store.ListBooks())
	case http.MethodPost:
		var b models.Book
		body, _ := io.ReadAll(r.Body)
		json.Unmarshal(body, &b)
		created := h.Store.CreateBook(b)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(created)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *BooksHandler) BookByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	switch r.Method {
	case http.MethodGet:
		b, ok := h.Store.GetBook(id)
		if !ok {
			http.NotFound(w, r)
			return
		}
		json.NewEncoder(w).Encode(b)
	case http.MethodPut:
		var b models.Book
		json.NewDecoder(r.Body).Decode(&b)
		b.ID = id
		updated, err := h.Store.UpdateBook(b)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(updated)
	case http.MethodDelete:
		err := h.Store.DeleteBook(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
