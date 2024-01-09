package repository

import (
	"context"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/jmoiron/sqlx"
)

type RepositoryInterface interface {
	CreateItemRepository(context.Context, *models.Item) (*models.Item, error)
	UpdateItemRepository(ctx context.Context, id int, item *models.Item) error
	DeleteItemRepository(ctx context.Context, id int) error
	ItemsRepository(ctx context.Context) ([]models.Item, error)
	FindItemRepository(ctx context.Context, id int) (models.Item, error)

	CreateUserRepository(ctx context.Context, user models.User) (int, error)
	UpdateUserRepository(ctx context.Context, id int, user *models.User) error
	DeleteUserRepository(ctx context.Context, id int) error
	FindUserRepository(ctx context.Context, id int) (models.User, error)
	FindUserByEmailRepository(ctx context.Context, email string) (*models.User, error)
	ListUserRepository(ctx context.Context) ([]models.User, error)

	CreateOrderRepository(ctx context.Context, order *models.Order) (int, error)
	DeleteOrderRepository(ctx context.Context, orderID int) error
	FindOrderRepository(ctx context.Context, orderID int) (models.Order, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) RepositoryInterface {
	return &repository{db: db}
}
