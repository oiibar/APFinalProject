package routes

import (
	"net/http"

	"final/internal/handlers"
	"final/internal/middleware"
	"final/internal/store"
)

func Routes(mux *http.ServeMux, st store.Store, q chan int) http.Handler {
	books := &handlers.BooksHandler{Store: st}
	orders := &handlers.OrdersHandler{Store: st, OrderQueue: q}
	auth := &handlers.AuthHandler{Store: st}
	users := &handlers.UsersHandler{Store: st}

	mux.HandleFunc("/api/books", books.Books)
	mux.HandleFunc("/api/book", books.BookByID)

	mux.HandleFunc("/api/orders", orders.Orders)
	mux.HandleFunc("/api/order", orders.OrderByID)

	mux.HandleFunc("/api/auth/signup", auth.Signup)
	mux.HandleFunc("/api/auth/login", auth.Login)

	mux.HandleFunc("/api/users/me", users.Me)
	mux.HandleFunc("/api/users", users.Users)

	return middleware.AuthMiddleware(mux)
}
