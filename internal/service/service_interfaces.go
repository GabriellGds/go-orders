package service

import (
	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/internal/repository"
	"github.com/GabriellGds/go-orders/pkg/errors"
)

type service struct {
	repository repository.RepositoryInterface
}

type Service interface {
	CreateUserService(*models.User) (*models.User, *errors.ErrorResponse)
	DeleteUserService(int) *errors.ErrorResponse
	UpdateUserService(int, *models.User) *errors.ErrorResponse
	FindUserService(int)(models.User, error)
	Login(*models.User) (*models.User, string, *errors.ErrorResponse)

	CreateOrderService(*models.Order) (*models.Order, *errors.ErrorResponse)
	DeleteOrderService(userID, id int) *errors.ErrorResponse
	FindOrderService(int) (models.Order, error)

	CreateItemService(*models.Item) (*models.Item, *errors.ErrorResponse)
	DeleteItemService(int) *errors.ErrorResponse
	UpdateItemSvice(int, *models.Item) *errors.ErrorResponse
	FindItemService(id int)(models.Item, error)
	ListItems() ([]models.Item, *errors.ErrorResponse)
}

func NewUserService(repository repository.RepositoryInterface) Service {
	return &service{repository: repository}
}




