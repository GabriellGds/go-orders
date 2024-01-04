package handlers

import (
	"net/http"
	"strconv"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/pkg/logger"
	"github.com/GabriellGds/go-orders/pkg/response"
	"github.com/go-chi/chi/v5"
)

func (h *handler) FindItem(w http.ResponseWriter, r *http.Request) {
	logger := logger.NewLogger("find item")
	logger.Info("start find item")

	itemID := chi.URLParam(r, "itemID")

	ID, err := strconv.Atoi(itemID)
	if err != nil {
		logger.Error("error trying to convert id", err)
		response.SendJSON(w, http.StatusBadRequest, models.ItemErrorParam("id", "queryParameter"))
		return
	}

	item, err := h.service.FindItemService(ID)
	if err != nil {
		response.SendJSON(w, http.StatusBadRequest, err)
		return
	}

	logger.Info("item found successfully")
	response.SendJSON(w, http.StatusOK, item)
}
