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

	Books  map[int]models.Book
	Orders map[int]models.Order
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		nextBookID:  1,
		nextOrderID: 1,
		Books:       make(map[int]models.Book),
		Orders:      make(map[int]models.Order),
	}
}

// BOOKS
func (s *MemoryStore) CreateBook(b models.Book) models.Book {
	s.mu.Lock()
	defer s.mu.Unlock()

	b.ID = s.nextBookID
	s.nextBookID++
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

// ORDERS
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

func (s *MemoryStore) GetOrder(id int) (models.Order, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	o, ok := s.Orders[id]
	return o, ok
}

func (s *MemoryStore) UpdateOrderStatus(id int, status string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	o := s.Orders[id]
	o.Status = status
	s.Orders[id] = o
}
