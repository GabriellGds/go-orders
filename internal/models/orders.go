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
	Price    float64 `json:"item_price" db:"item_price"`
}

type OrderCreatedResponse struct {
	ID int `json:"id"`
}

type OrderRequest struct {
	Items []OrderItem `json:"items"`
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

func OrderRequestToOrderItems(request OrderRequest) []OrderItems {
	var orderItems []OrderItems
	for _, order := range request.Items {
		orders := OrderItems{
			Quantity: order.Quantity,
			ItemID: order.ItemID,
		}
		orderItems = append(orderItems, orders)
	}

	return orderItems
}