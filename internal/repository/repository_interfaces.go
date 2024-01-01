package repository

import "github.com/GabriellGds/go-orders/internal/models"

type itemRepositoryInterface interface {
	CreateItem(models.Item) (models.Item, error)
	UpdateItem(id int, item models.ItemRequest) error
	DeleteItem(id int) error
	AllItems() ([]models.Item, error)
	FindItem(id int) (models.Item, error)
}

type UserRepositoryInterface interface {
	CreateUser(user models.User) (models.User, error)
	UpdateUser(id int, user models.UserUpdateRequest) error
	DeleteUser(id int) error
	User(id int) (models.User, error)
	FindEmail(string) (*models.User, error)
}
