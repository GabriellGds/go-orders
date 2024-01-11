package handlers

import (
	"net/http"
	"strconv"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/pkg/logger"
	"github.com/GabriellGds/go-orders/pkg/response"
	"github.com/go-chi/chi/v5"
)

// @Summary Find user 
// @Description Retrieves item details based on the item ID
// @Tags Items
// @Accept json
// @Produce json
// @Param itemID path string true "ID of the item to be retrieved"
// @Success 200 {object} models.Item "Item information retrieved successfully"
// @Failure 400 {object} errors.ErrorResponse "Error: Invalid id"
// @Failure 404 {object} errors.ErrorResponse "Item not found"
// @Router /items/{itemID} [get]
// @Security KeyAuth
func (h *handler) FindItem(w http.ResponseWriter, r *http.Request) {
	logger := logger.NewLogger("find item")
	logger.Info("start find item")
	ctx := r.Context()

	itemID := chi.URLParam(r, "itemID")

	ID, err := strconv.Atoi(itemID)
	if err != nil {
		logger.Error("error trying to convert id", err)
		response.SendJSON(w, http.StatusBadRequest, models.ItemErrorParam("id", "queryParameter"))
		return
	}

	item, err := h.service.FindItemService(ctx, ID)
	if err != nil {
		response.SendJSON(w, http.StatusNotFound, err)
		return
	}

	logger.Info("item found successfully")
	response.SendJSON(w, http.StatusOK, item)
}
