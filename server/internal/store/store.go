package store

import "final/internal/models"

type Store interface {
	CreateUser(models.User) (models.User, error)
	GetUserByEmail(string) (models.User, bool)
	GetUserByID(int) (models.User, bool)
	ListUsers() []models.User
	UpdateUser(models.User) (models.User, error)
	DeleteUser(int) error

	CreateBook(models.Book) models.Book
	ListBooks() []models.Book
	GetBook(int) (models.Book, bool)
	UpdateBook(models.Book) (models.Book, error)
	DeleteBook(int) error

	CreateOrder(models.Order) models.Order
	ListOrders() []models.Order
	ListOrdersByUser(userID int) []models.Order
	GetOrder(int) (models.Order, bool)
	UpdateOrder(models.Order) (models.Order, error)
	DeleteOrder(int) error
	UpdateOrderStatus(id int, status string) error
}
