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

	// Books: GET is public, mutations are admin-only.
	mux.Handle("/api/books", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			books.Books(w, r)
			return
		}
		middleware.RequireAdmin(http.HandlerFunc(books.Books)).ServeHTTP(w, r)
	}))
	mux.Handle("/api/book", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			books.BookByID(w, r)
			return
		}
		middleware.RequireAdmin(http.HandlerFunc(books.BookByID)).ServeHTTP(w, r)
	}))

	// Orders: user must be authenticated. Only admin can update/delete by id.
	mux.Handle("/api/orders", middleware.RequireAuth(http.HandlerFunc(orders.Orders)))
	mux.Handle("/api/order", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut || r.Method == http.MethodDelete {
			middleware.RequireAdmin(http.HandlerFunc(orders.OrderByID)).ServeHTTP(w, r)
			return
		}
		middleware.RequireAuth(http.HandlerFunc(orders.OrderByID)).ServeHTTP(w, r)
	}))

	mux.HandleFunc("/api/auth/signup", auth.Signup)
	mux.HandleFunc("/api/auth/login", auth.Login)

	mux.Handle("/api/users/me", middleware.RequireAuth(http.HandlerFunc(users.Me)))
	mux.Handle("/api/users", middleware.RequireAdmin(http.HandlerFunc(users.Users)))

	// Global middleware chain: CORS -> (optional) Auth parser -> mux
	return middleware.CORS(middleware.AuthMiddleware(mux))
}
