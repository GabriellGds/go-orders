package models

import "time"

type Order struct {
	ID        int          `json:"id"`
	UserID    int          `json:"-" db:"user_id"`
	Items     []OrderItems `json:"items"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time   `json:"deleted_at" db:"deleted_at"`
}

type OrderItem struct {
	ItemID   int `json:"item_id"`
	Quantity int `json:"quantity"`
}

type OrderItems struct {
	ID       int    `json:"-"`
	ItemID   int    `json:"item_id" db:"item_id"`
	OrderID  int    `json:"order_id" db:"order_id"`
	Quantity int    `json:"quantity"`
	Name     string `json:"name" db:"item_name"`
	Price    string `json:"item_price" db:"item_price"`
}

type OrderRequest struct {
	Items []OrderItems `json:"items"`
}

type OrderCreatedResponse struct {
	ID int `json:"id"`
}

func NewOrder(userID int, items []OrderItems) *Order {
	return &Order{
		UserID:    userID,
		Items:     items,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}
}
