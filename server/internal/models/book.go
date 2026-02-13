package models

import "time"

type Book struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Price     float64   `json:"price"`
	Author    string    `json:"author"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"createdAt"`
}
