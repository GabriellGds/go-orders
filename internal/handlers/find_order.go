package handlers

import (
	"net/http"
	"strconv"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/pkg/logger"
	"github.com/GabriellGds/go-orders/pkg/response"
	"github.com/go-chi/chi/v5"
)

func (h *handler) FindOrder(w http.ResponseWriter, r *http.Request) {
	logger := logger.NewLogger("find order")
	logger.Info("start find order")
	ctx := r.Context()

	itemID := chi.URLParam(r, "orderID")
	ID, err := strconv.Atoi(itemID)
	if err != nil {
		logger.Error("error trying to convert id", err)
		response.SendJSON(w, http.StatusBadRequest, models.OrderErrorParam("id", "missing or invalid parameter"))
		return
	}

	order, err := h.service.FindOrderService(ctx, ID)
	if err != nil {
		response.SendJSON(w, http.StatusNotFound, err)
		return
	}

	logger.Info("order found successfuly")
	response.SendJSON(w, http.StatusOK, order)
}
