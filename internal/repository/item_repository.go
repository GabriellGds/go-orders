package repository

import (
	"context"

	"github.com/GabriellGds/go-orders/internal/models"
)

func (r repository) CreateItemRepository(ctx context.Context, item *models.Item) (*models.Item, error) {
	stmt, err := r.db.PreparexContext(ctx, `INSERT INTO items (name, price) VALUES($1, $2) RETURNING id`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var lastInsertID int
	err = stmt.QueryRowContext(ctx, item.Name, item.Price).Scan(&lastInsertID)
	if err != nil {
		return nil, err
	}

	item.ID = lastInsertID

	return item, nil
}

func (r repository) UpdateItemRepository(ctx context.Context, id int, item *models.Item) error {
	stmt, err := r.db.PreparexContext(ctx, `UPDATE items SET name = $1, price = $2 WHERE id = $3 AND deleted_at IS NULL`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.ExecContext(ctx, item.Name, item.Price, id); err != nil {
		return err
	}

	return nil
}

func (r repository) DeleteItemRepository(ctx context.Context, id int) error {
	stmt, err := r.db.PreparexContext(ctx, `UPDATE items SET deleted_at = NOW() where id = $1 and deleted_at IS NULL`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.ExecContext(ctx, id); err != nil {
		return err
	}

	return nil
}

func (r repository) FindItemRepository(ctx context.Context, id int) (models.Item, error) {
	stmt, err := r.db.PreparexContext(ctx, `SELECT * FROM items WHERE id = $1 AND deleted_at IS NULL`)
	if err != nil {
		return models.Item{}, err
	}
	defer stmt.Close()

	var item models.Item
	if err := stmt.GetContext(ctx, &item, id); err != nil {
		return models.Item{}, err
	}

	return item, nil
}

func (r repository) ItemsRepository(ctx context.Context) ([]models.Item, error) {
	stmt, err := r.db.PreparexContext(ctx, `SELECT * FROM items WHERE deleted_at IS NULL`)
	if err != nil {
		return []models.Item{}, nil
	}
	defer stmt.Close()

	var items []models.Item
	if err := stmt.SelectContext(ctx, &items); err != nil {
		return []models.Item{}, err
	}

	return items, nil
}
