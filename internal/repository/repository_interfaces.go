package repository

import (
	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/jmoiron/sqlx"
)

type RepositoryInterface interface {
	CreateItemRepository(*models.Item) (*models.Item, error)
	UpdateItemRepository(id int, item *models.Item) error
	DeleteItemRepository(id int) error
	ItemsRepository() ([]models.Item, error)
	FindItemRepository(id int) (models.Item, error)

	CreateUserRepository(user models.User) (int, error)
	UpdateUserRepository(id int, user *models.User) error
	DeleteUserRepository(id int) error
	FindUserRepository(id int) (models.User, error)
	FindUserByEmailRepository(string) (*models.User, error)

	CreateOrderRepository(order *models.Order) (int, error)
	DeleteOrderRepository(orderID int) error
	FindOrderRepository(orderID int) (models.Order, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) RepositoryInterface {
	return &repository{db: db}
}
