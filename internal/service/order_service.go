package service

import (
	"context"
	"net/http"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/pkg/errors"
	"github.com/GabriellGds/go-orders/pkg/logger"
)

func (s *service) CreateOrderService(ctx context.Context, order *models.Order) (*models.Order, *errors.ErrorResponse) {
	logger := logger.NewLogger("create order service")
	logger.Info("start createOrder service")
	id, err := s.repository.CreateOrderRepository(ctx, order)
	if err != nil {
		logger.Error("error to creating order on database", err)
		return nil, &errors.ErrorResponse{
			Message: "error to creating user",
			Code:    http.StatusInternalServerError,
		}
	}
	
	order.ID = id
	return order, nil
}

func (s *service) DeleteOrderService(ctx context.Context, userID, id int) *errors.ErrorResponse {
	logger := logger.NewLogger("deleteOrder service")
	logger.Info("start deleteOrder service")
	order, err := s.repository.FindOrderRepository(ctx, id)
	if err != nil {
		return &errors.ErrorResponse{
			Message: "order not found",
			Code:    http.StatusNotFound,
		}
	}
	
	if order.UserID != userID {
		logger.Error("forbiden, you not have permission")
		return &errors.ErrorResponse{
			Message: "access denied. You do not have permission to perform this action",
			Code:    http.StatusUnauthorized,
		}
	}

	err = s.repository.DeleteOrderRepository(ctx, id)
	if err != nil {
		return &errors.ErrorResponse{
			Message: "error to deleting order",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

func (s *service) FindOrderService(ctx context.Context, id int) (models.Order, error) {
	logger := logger.NewLogger("findOrder service")
	logger.Info("start findOrder service")

	order, err := s.repository.FindOrderRepository(ctx, id)
	if err != nil {
		logger.Error("error to finding order ", err)
		return models.Order{}, &errors.ErrorResponse{
			Message: "order not found",
		}
	}

	return order, nil
}
