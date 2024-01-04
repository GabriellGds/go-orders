package service

import (
	"context"
	"net/http"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/pkg/errors"
	"github.com/GabriellGds/go-orders/pkg/logger"
	
)

func (s *service) CreateItemService(ctx context.Context, item *models.Item) (*models.Item, *errors.ErrorResponse) {
	item, err := s.repository.CreateItemRepository(ctx, item)
	if err != nil {
		return nil, &errors.ErrorResponse{
			Message: "error to creating user",
			Code:    http.StatusInternalServerError,
		}
	}

	return item, nil
}

func (s *service) DeleteItemService(ctx context.Context, id int) *errors.ErrorResponse {
	logger := logger.NewLogger("deleteItem service")
	logger.Info("start deleteItem service")
	_, err := s.repository.FindItemRepository(ctx, id)
	if err != nil {
		return &errors.ErrorResponse{
			Message: "item not found",
			Code:    http.StatusNotFound,
		}
	}

	err = s.repository.DeleteItemRepository(ctx, id)
	if err != nil {
		return &errors.ErrorResponse{
			Message: "error to deleting item on database",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

func (s *service) UpdateItemSvice(ctx context.Context, id int, item *models.Item) *errors.ErrorResponse {
	logger := logger.NewLogger("updateItem service")
	logger.Info("start updataItem service")
	_, err := s.repository.FindItemRepository(ctx, id)
	if err != nil {
		logger.Error("error to finding item", err)
		return &errors.ErrorResponse{
			Message: "item not found",
			Code:    http.StatusNotFound,
		}
	}

	if err := s.repository.UpdateItemRepository(ctx, id, item); err != nil {
		logger.Error("error to updating item on database ", err)
		return &errors.ErrorResponse{
			Message: "error to updating item",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

func (s *service) FindItemService(ctx context.Context, id int) (models.Item, error) {
	logger := logger.NewLogger("find item service")
	logger.Info("start find item service")
	item, err := s.repository.FindItemRepository(ctx, id)
	if err != nil {
		logger.Error("error to finding item", err)
		return models.Item{}, &errors.ErrorResponse{
			Message: "item not found",
			Code:    http.StatusNotFound,
		}
	}

	return item, nil
}

func (s *service) ListItems(ctx context.Context) ([]models.Item, *errors.ErrorResponse) {
	logger := logger.NewLogger("list items service")
	logger.Info("start list items service")

	items, err := s.repository.ItemsRepository(ctx)
	if err != nil {
		logger.Error("error to listing items", err)
		return []models.Item{}, &errors.ErrorResponse{
			Message: "error to finding items on database",
			Code:    http.StatusInternalServerError,
		}
	}

	return items, nil
}
