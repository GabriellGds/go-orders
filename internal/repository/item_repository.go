package repository

import (
	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/jmoiron/sqlx"
)

type itemRepository struct {
	db *sqlx.DB
}

func NewItemRepository(db *sqlx.DB) itemRepositoryInterface {
	return &itemRepository{db: db}
}

func (ir itemRepository) CreateItem(item models.Item) (models.Item, error){
	stmt, err := ir.db.Preparex(`INSERT INTO items (name, price) VALUES($1, $2) RETURNING id`)
	if err != nil {
		return models.Item{}, err
	}
	defer stmt.Close()

	var lastInsertID int
	err = stmt.QueryRow(item.Name, item.Price).Scan(&lastInsertID)
	if err != nil {
		return models.Item{}, err
	}

	item.ID = lastInsertID

	return item, nil
}

func (ir itemRepository) UpdateItem(id int, item models.ItemRequest) error {
	stmt, err := ir.db.Preparex(`UPDATE items SET name = $1, price = $2 WHERE id = $3 AND deleted_at IS NULL`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(item.Name, item.Price, id); err != nil {
		return err
	}

	return nil
}

func (ir itemRepository) DeleteItem(id int) error {
	stmt, err := ir.db.Preparex(`UPDATE items SET deleted_at = NOW() where id = $1 and deleted_at IS NULL`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return err
	}

	return nil
}

func (ir itemRepository) FindItem(id int) (models.Item, error) {
	stmt, err := ir.db.Preparex(`SELECT * FROM items WHERE id = $1 AND deleted_at IS NULL`)
	if err != nil {
		return models.Item{}, err
	}
	defer stmt.Close()

	var item models.Item
	if err := stmt.Get(&item, id); err != nil {
		return models.Item{}, err
	}

	return item, nil
}

func (ir itemRepository) AllItems() ([]models.Item, error) {
	stmt, err := ir.db.Preparex(`SELECT * FROM items WHERE deleted_at IS NULL`)
	if err != nil {
		return []models.Item{}, nil
	}
	defer stmt.Close()

	var items []models.Item
	if err := stmt.Select(&items); err != nil {
		return []models.Item{}, err
	}

	return items, nil
}
