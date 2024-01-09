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

// @Summary Delete order
// @Description Deletes an order based on the ID provided
// @Tags Orders
// @Accept json
// @Produce json
// @Param orderID path string true "ID of the order to be deleted"
// @Success 204
// @Failure 400 {object} errors.ErrorResponse
// @Failure 401 {object} errors.ErrorResponse
// @Failure 403 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /orders/{orderID} [delete]
// @Security KeyAuth
func (h *handler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	logger := logger.NewLogger("delete order")
	logger.Info("start delete order")
	ctx := r.Context()
	tokenID, err := models.GetUserIDFromToken(r)
	if err != nil {
		logger.Error("invalid token")
		response.SendJSON(w, http.StatusUnauthorized, errors.ErrorResponse{
			Message: "invalid token",
		})
		return
	}
	
	id := chi.URLParam(r, "orderID")
	orderID, err := strconv.Atoi(id)
	if err != nil {
		logger.Error("error trying to convert id", err)
		response.SendJSON(w, http.StatusBadRequest, models.OrderErrorParam("id", "missing or invalid parameter"))
		return
	}

	if err := h.service.DeleteOrderService(ctx, tokenID, orderID); err != nil {
		response.SendJSON(w, err.Code, err)
		return
	}

	logger.Info("order deleting successfuly")
	response.SendJSON(w, http.StatusNoContent, nil)
}
