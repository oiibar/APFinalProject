package routes

import (
	"net/http"

	"final/internal/handlers"
	"final/internal/store"
)

func Routes(mux *http.ServeMux, st *store.MemoryStore, q chan int) {
	books := &handlers.BooksHandler{Store: st}
	orders := &handlers.OrdersHandler{Store: st, OrderQueue: q}

	mux.HandleFunc("/api/books", books.Books)
	mux.HandleFunc("/api/book", books.BookByID)

	mux.HandleFunc("/api/orders", orders.Create)
	mux.HandleFunc("/api/order", orders.GetByID)
}
