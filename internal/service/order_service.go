package service

import (
	"net/http"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/pkg/errors"
	"github.com/GabriellGds/go-orders/pkg/logger"
)

func (s *service) CreateOrderService(order *models.Order) (*models.Order, *errors.ErrorResponse) {
	logger := logger.NewLogger("create order service")
	logger.Info("start createOrder service")
	id, err := s.repository.CreateOrderRepository(order)
	if err != nil {
		logger.Error("error to creating order on database", err)
		return nil, &errors.ErrorResponse{
			Message: "error to creating user",
			Code: http.StatusInternalServerError,
		}
	}
	logger.Info("ola", order.Items)
	logger.Info(id)
	order.ID = id
	logger.Info(order.ID)
	return order, nil
}

func (s *service) DeleteOrderService(userID, id int) *errors.ErrorResponse {
	logger := logger.NewLogger("deleteOrder service")
	logger.Info("start deleteOrder service")
	order, err := s.repository.FindOrderRepository(id)
	if err != nil {
		return &errors.ErrorResponse{
			Message: "order not found",
			Code: http.StatusNotFound,
		}
	}
	if order.UserID != userID {
		logger.Error("forbiden, you not have permission")
		return &errors.ErrorResponse {
			Message: "access denied. You do not have permission to perform this action",
			Code: http.StatusUnauthorized,
		}
	}

	err = s.repository.DeleteOrderRepository(id)
	if err != nil {
		return &errors.ErrorResponse{
			Message: "error to deleting order",
			Code: http.StatusInternalServerError,
		}
	}

	return nil
}

func (s *service) FindOrderService(id int) (models.Order, error) {
	logger := logger.NewLogger("findOrder service")
	logger.Info("start findOrder service")
	
	order, err := s.repository.FindOrderRepository(id)
	if err != nil {
		logger.Error("error to finding order ", err)
		return models.Order{}, &errors.ErrorResponse{
			Message: "order not found",
		}
	}

	return order, nil
}
