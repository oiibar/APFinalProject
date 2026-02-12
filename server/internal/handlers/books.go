package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"final/internal/models"
	"final/internal/store"
)

type BooksHandler struct {
	Store store.Store
}

func (h *BooksHandler) Books(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, http.StatusOK, h.Store.ListBooks())
	case http.MethodPost:
		var b models.Book
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			writeError(w, http.StatusBadRequest, "invalid json")
			return
		}
		b.Title = strings.TrimSpace(b.Title)
		b.Author = strings.TrimSpace(b.Author)
		if b.Title == "" || b.Author == "" || b.Price <= 0 {
			writeError(w, http.StatusBadRequest, "title, author and price (>0) are required")
			return
		}
		created := h.Store.CreateBook(b)
		writeJSON(w, http.StatusCreated, created)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *BooksHandler) BookByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}
	switch r.Method {
	case http.MethodGet:
		b, ok := h.Store.GetBook(id)
		if !ok {
			writeError(w, http.StatusNotFound, "book not found")
			return
		}
		writeJSON(w, http.StatusOK, b)
	case http.MethodPut:
		var b models.Book
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			writeError(w, http.StatusBadRequest, "invalid json")
			return
		}
		b.Title = strings.TrimSpace(b.Title)
		b.Author = strings.TrimSpace(b.Author)
		if b.Title == "" || b.Author == "" || b.Price <= 0 {
			writeError(w, http.StatusBadRequest, "title, author and price (>0) are required")
			return
		}
		b.ID = id
		updated, err := h.Store.UpdateBook(b)
		if err != nil {
			writeError(w, http.StatusNotFound, "book not found")
			return
		}
		writeJSON(w, http.StatusOK, updated)
	case http.MethodDelete:
		err := h.Store.DeleteBook(id)
		if err != nil {
			writeError(w, http.StatusNotFound, "book not found")
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
