package handlers

import (
	"net/http"
	"strconv"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/pkg/logger"
	"github.com/GabriellGds/go-orders/pkg/response"
	"github.com/go-chi/chi/v5"
)

func (h *handler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	logger := logger.NewLogger("delete item")
	logger.Info("start delete item")

	itemID := chi.URLParam(r, "itemID")

	ID, err := strconv.Atoi(itemID)
	if err != nil {
		logger.Error("error trying to convert id", err)
		response.SendJSON(w, http.StatusBadRequest, models.ItemErrorParam("id", "missing or invalid parameter").Error())
		return
	}

	if err := h.service.DeleteItemService(ID); err != nil {
		response.SendJSON(w, err.Code, err)
		return
	}

	logger.Info("Successfully deleted item")
	response.SendJSON(w, http.StatusNoContent, nil)
}
