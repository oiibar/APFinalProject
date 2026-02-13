package store

import (
	"sync"
	"time"

	"final/internal/models"
)

type MemoryStore struct {
	mu sync.RWMutex

	nextBookID  int
	nextOrderID int
	nextUserID  int

	Books  map[int]models.Book
	Orders map[int]models.Order
	Users  map[int]models.User
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		nextBookID:  1,
		nextOrderID: 1,
		nextUserID:  1,
		Books:       make(map[int]models.Book),
		Orders:      make(map[int]models.Order),
		Users:       make(map[int]models.User),
	}
}

func (s *MemoryStore) CreateBook(b models.Book) models.Book {
	s.mu.Lock()
	defer s.mu.Unlock()

	b.ID = s.nextBookID
	s.nextBookID++
	b.CreatedAt = time.Now()
	s.Books[b.ID] = b
	return b
}
func (s *MemoryStore) ListBooks() []models.Book {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := []models.Book{}
	for _, b := range s.Books {
		out = append(out, b)
	}
	return out
}
func (s *MemoryStore) GetBook(id int) (models.Book, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	b, ok := s.Books[id]
	return b, ok
}
func (s *MemoryStore) UpdateBook(b models.Book) (models.Book, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.Books[b.ID]
	if !ok {
		return models.Book{}, ErrNotFound
	}
	b.CreatedAt = s.Books[b.ID].CreatedAt
	s.Books[b.ID] = b
	return b, nil
}
func (s *MemoryStore) DeleteBook(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.Books[id]
	if !ok {
		return ErrNotFound
	}
	delete(s.Books, id)
	return nil
}

func (s *MemoryStore) CreateOrder(o models.Order) models.Order {
	s.mu.Lock()
	defer s.mu.Unlock()

	o.ID = s.nextOrderID
	s.nextOrderID++
	o.Status = "processing"
	o.CreatedAt = time.Now()
	s.Orders[o.ID] = o
	return o
}
func (s *MemoryStore) ListOrders() []models.Order {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := []models.Order{}
	for _, o := range s.Orders {
		out = append(out, o)
	}
	return out
}

func (s *MemoryStore) ListOrdersByUser(userID int) []models.Order {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := []models.Order{}
	for _, o := range s.Orders {
		if o.UserID == userID {
			out = append(out, o)
		}
	}
	return out
}
func (s *MemoryStore) GetOrder(id int) (models.Order, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	o, ok := s.Orders[id]
	return o, ok
}
func (s *MemoryStore) UpdateOrder(o models.Order) (models.Order, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.Orders[o.ID]
	if !ok {
		return models.Order{}, ErrNotFound
	}
	o.CreatedAt = s.Orders[o.ID].CreatedAt
	s.Orders[o.ID] = o
	return o, nil
}
func (s *MemoryStore) DeleteOrder(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.Orders[id]
	if !ok {
		return ErrNotFound
	}
	delete(s.Orders, id)
	return nil
}
func (s *MemoryStore) UpdateOrderStatus(id int, status string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	o, ok := s.Orders[id]
	if !ok {
		return ErrNotFound
	}
	o.Status = status
	s.Orders[id] = o
	return nil
}

func (s *MemoryStore) CreateUser(u models.User) (models.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	u.ID = s.nextUserID
	s.nextUserID++
	u.CreatedAt = time.Now()
	s.Users[u.ID] = u
	return u, nil
}
func (s *MemoryStore) ListUsers() []models.User {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := []models.User{}
	for _, u := range s.Users {
		out = append(out, u)
	}
	return out
}
func (s *MemoryStore) GetUserByEmail(email string) (models.User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, u := range s.Users {
		if u.Email == email {
			return u, true
		}
	}
	return models.User{}, false
}
func (s *MemoryStore) GetUserByID(id int) (models.User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	u, ok := s.Users[id]
	return u, ok
}
func (s *MemoryStore) UpdateUser(u models.User) (models.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.Users[u.ID]
	if !ok {
		return models.User{}, ErrNotFound
	}
	u.CreatedAt = s.Users[u.ID].CreatedAt
	s.Users[u.ID] = u
	return u, nil
}
func (s *MemoryStore) DeleteUser(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.Users[id]
	if !ok {
		return ErrNotFound
	}
	delete(s.Users, id)
	return nil
}

var ErrNotFound = &NotFoundError{}

type NotFoundError struct{}

func (e *NotFoundError) Error() string { return "not found" }
