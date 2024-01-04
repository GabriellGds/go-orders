package handlers

import (
	"net/http"
	"strconv"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/pkg/errors"
	"github.com/GabriellGds/go-orders/pkg/logger"
	"github.com/GabriellGds/go-orders/pkg/response"
	"github.com/go-chi/chi/v5"
)

func (h *handler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	logger := logger.NewLogger("delete order")
	logger.Info("start delete order")

	tokenID, err := models.GetUserIDFromToken(r)
	if err != nil {
		logger.Error("invalid token")
		response.SendJSON(w, http.StatusUnauthorized, errors.ErrorResponse{
			Message: "invalid token",
		})
		return
	}
	logger.Info(tokenID)
	id := chi.URLParam(r, "orderID")
	orderID, err := strconv.Atoi(id)
	if err != nil {
		logger.Error("error trying to convert id", err)
		response.SendJSON(w, http.StatusBadRequest, models.OrderErrorParam("id", "missing or invalid parameter"))
		return
	}

	if err := h.service.DeleteOrderService(tokenID, orderID); err != nil {
		response.SendJSON(w, err.Code, err)
		return
	}

	logger.Info("order deleting successfuly")
	response.SendJSON(w, http.StatusNoContent, nil)
}
