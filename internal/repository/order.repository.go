package repository

import (
	"context"

	"github.com/GabriellGds/go-orders/internal/models"
)

func (r *repository) CreateOrderRepository(ctx context.Context, order *models.Order) (int, error) {
	stmt, err := r.db.PreparexContext(ctx, "INSERT INTO orders (user_id) VALUES ($1) RETURNING id")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var orderID int
	err = stmt.QueryRowContext(ctx, order.UserID).Scan(&orderID)
	if err != nil {
		return 0, err
	}

	statement, err := r.db.PreparexContext(ctx, "INSERT INTO order_items (order_id, item_id, quantity) VALUES ($1, $2, $3)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	for _, item := range order.Items {
		_, err = statement.ExecContext(ctx, orderID, item.ItemID, item.Quantity)
		if err != nil {
			return 0, err
		}
	}

	return orderID, nil
}

func (r *repository) DeleteOrderRepository(ctx context.Context, orderID int) error {
	stmt, err := r.db.PreparexContext(ctx, "UPDATE orders SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.ExecContext(ctx, orderID); err != nil {
		return err
	}

	return nil
}

func (r *repository) FindOrderRepository(ctx context.Context, orderID int) (models.Order, error) {
	stmt, err := r.db.PreparexContext(ctx, "SELECT * FROM orders WHERE id = $1 AND deleted_at IS NULL")
	if err != nil {
		return models.Order{}, err
	}
	defer stmt.Close()

	var order models.Order
	if err := stmt.GetContext(ctx, &order, orderID); err != nil {
		return models.Order{}, err
	}

	statement, err := r.db.PreparexContext(ctx, `SELECT oi.*, i.name AS item_name, i.price AS item_price 
	FROM order_items oi 
	INNER JOIN items i ON oi.item_id = i.id 
	WHERE oi.order_id = $1`)
	if err != nil {
		return models.Order{}, err
	}
	defer statement.Close()
	
	var items []models.OrderItems
	err = statement.SelectContext(ctx, &items, orderID)
	if err != nil {
		return models.Order{}, err
	}

	for i := range items {
		order.Items = append(order.Items, models.OrderItems{
			ID:       items[i].ID,
			ItemID:   items[i].ItemID,
			OrderID:  items[i].OrderID,
			Quantity: items[i].Quantity,
			Name:     items[i].Name,
			Price:    items[i].Price,
		})
	}

	return order, nil
}
