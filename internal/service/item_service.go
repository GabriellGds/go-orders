package service

import (
	"net/http"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/pkg/errors"
	"github.com/GabriellGds/go-orders/pkg/logger"
)

func (s *service) CreateItemService(item *models.Item) (*models.Item, *errors.ErrorResponse) {
	item, err := s.repository.CreateItemRepository(item)
	if err != nil {
		return nil, &errors.ErrorResponse{
			Message: "error to creating user",
			Code: http.StatusInternalServerError,
		}
	}

	return item, nil
}

func (s *service) DeleteItemService(id int) *errors.ErrorResponse {
	logger := logger.NewLogger("deleteItem service")
	logger.Info("start deleteItem service")
	_, err := s.repository.FindItemRepository(id)
	if err != nil {
		return &errors.ErrorResponse{
			Message: "item not found",
			Code: http.StatusNotFound,
		}
	}

	err = s.repository.DeleteItemRepository(id)
	if err != nil {
		return &errors.ErrorResponse{
			Message: "error to deleting item on database",
			Code: http.StatusInternalServerError,
		}
	}

	return nil
}

func (s *service) UpdateItemSvice(id int, item *models.Item) *errors.ErrorResponse{
	logger := logger.NewLogger("updateItem service")
	logger.Info("start updataItem service")
	_, err := s.repository.FindItemRepository(id)
	if err != nil {
		logger.Error("error to finding item", err)
		return &errors.ErrorResponse{
			Message: "item not found",
			Code: http.StatusNotFound,
		}
	}

	if err := s.repository.UpdateItemRepository(id, item); err != nil {
		logger.Error("error to updating item on database ", err)
		return &errors.ErrorResponse{
			Message: "error to updating item",
			Code: http.StatusInternalServerError,
		}
	}

	return nil
}

func (s *service) FindItemService(id int)(models.Item, error) {
	logger := logger.NewLogger("find item service")
	logger.Info("start find item service")
	item, err := s.repository.FindItemRepository(id)
	if err != nil {
		logger.Error("error to finding item", err)
		return models.Item{}, &errors.ErrorResponse{
			Message: "item not found",
			Code: http.StatusNotFound,
		}
	}

	return item, nil
}

func (s *service) ListItems() ([]models.Item, *errors.ErrorResponse) {
	logger := logger.NewLogger("list items service")
	logger.Info("start list items service")
	
	items, err := s.repository.ItemsRepository()
	if err != nil {
		logger.Error("error to listing items", err)
		return []models.Item{}, &errors.ErrorResponse{
			Message: "error to finding items on database",
			Code: http.StatusInternalServerError,
		}
	}

	return items, nil
}


