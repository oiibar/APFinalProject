package models

import "time"

type OrderItem struct {
	BookID   int `json:"bookId"`
	Quantity int `json:"quantity"`
}

type Order struct {
	ID        int         `json:"id"`
	UserID    int         `json:"userId"`
	Items     []OrderItem `json:"items"`
	Status    string      `json:"status"`
	CreatedAt time.Time   `json:"createdAt"`
}
