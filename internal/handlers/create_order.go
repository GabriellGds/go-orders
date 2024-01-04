package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/pkg/errors"
	"github.com/GabriellGds/go-orders/pkg/logger"
	"github.com/GabriellGds/go-orders/pkg/response"
)

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

	if err := orderRequest.Validate(); err != nil {
		logger.Error("error to validate", err)
		response.SendJSON(w, http.StatusBadRequest, err)
		return
	}

	for _, item := range orderRequest.Items {
		if err := item.Validate(); err != nil {
			logger.Error("error to validate", err)
			response.SendJSON(w, http.StatusBadRequest, err)
			return
		}
	}
	order := models.NewOrder(tokenID, orderRequest.Items)

	orderResult, er := h.service.CreateOrderService(ctx, order)
	if er != nil {
		response.SendJSON(w, er.Code, err)
		return
	}
	orderReponse := models.OrderCreatedResponse{ID: orderResult.ID}
	
	logger.Info("order creating successfuly")
	response.SendJSON(w, http.StatusCreated, orderReponse)
}
