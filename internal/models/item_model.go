package models

import "time"

type ItemRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Item struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Price     float64    `json:"price"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

func NewItem(name string, price float64) *Item {
	return &Item{
		Name: name,
		Price: price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}
}
