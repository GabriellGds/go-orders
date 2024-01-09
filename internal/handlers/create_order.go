package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/pkg/errors"
	"github.com/GabriellGds/go-orders/pkg/logger"
	"github.com/GabriellGds/go-orders/pkg/response"
)

// @Summary Create order
// @Description Create a new order 
// @Tags Orders
// @Accept json
// @Produce json
// @Param request body models.OrderRequest true "request body"
// @Success 201 {object} models.OrderCreatedResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 401 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /orders/ [post]
// @Security KeyAuth
func (h *handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	logger := logger.NewLogger("createOrder")
	logger.Info("start create order")
	ctx := r.Context()

	tokenID, err := models.GetUserIDFromToken(r)
	if err != nil {
		logger.Error("invalid token", err)
		response.SendJSON(w, http.StatusUnauthorized, err)
		return
	}
	
	var orderRequest models.OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&orderRequest); err != nil {
		logger.Error("error to unmarshal ", err)
		response.SendJSON(w, http.StatusBadRequest, errors.ErrorResponse{
			Message: "invalid type",
		})
		return
	}

	for _, order := range orderRequest.Items {
		if err := order.Validate(); err != nil {
			logger.Error("error to validate", err)
			response.SendJSON(w, err.Code, err)
			return
		}
	}
	orderItems := models.OrderRequestToOrderItems(orderRequest)

	order := models.NewOrder(tokenID, orderItems)

	orderResult, er := h.service.CreateOrderService(ctx, order)
	if er != nil {
		response.SendJSON(w, er.Code, er)
		return
	}
	orderReponse := models.OrderCreatedResponse{ID: orderResult.ID}
	
	logger.Info("order creating successfuly")
	response.SendJSON(w, http.StatusCreated, orderReponse)
}
