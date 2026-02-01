package models

type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Price  float64 `json:"price"`
	Author string  `json:"author"`
}
