package service

import (
	"context"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/internal/repository"
	"github.com/GabriellGds/go-orders/pkg/errors"
)

type service struct {
	repository repository.RepositoryInterface
}

type Service interface {
	CreateUserService(context.Context, *models.User) (*models.User, *errors.ErrorResponse)
	DeleteUserService(context.Context, int) *errors.ErrorResponse
	UpdateUserService(context.Context, int, *models.User) *errors.ErrorResponse
	FindUserService(context.Context, int)(models.User, error)
	Login(context.Context, *models.User) (*models.User, string, *errors.ErrorResponse)
	ListUsers(ctx context.Context) ([]models.User, *errors.ErrorResponse)

	CreateOrderService(context.Context, *models.Order) (*models.Order, *errors.ErrorResponse)
	DeleteOrderService(ctx context.Context, userID, id int) *errors.ErrorResponse
	FindOrderService(context.Context, int) (models.Order, error)

	CreateItemService(context.Context, *models.Item) (*models.Item, *errors.ErrorResponse)
	DeleteItemService(context.Context, int) *errors.ErrorResponse
	UpdateItemSvice(context.Context, int, *models.Item) *errors.ErrorResponse
	FindItemService(context.Context, int)(models.Item, error)
	ListItems(ctx context.Context) ([]models.Item, *errors.ErrorResponse)
}

func NewService(repository repository.RepositoryInterface) Service {
	return &service{repository: repository}
}




