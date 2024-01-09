package handlers

import (
	"net/http"
	"strconv"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/pkg/logger"
	"github.com/GabriellGds/go-orders/pkg/response"
	"github.com/go-chi/chi/v5"
)

// @Summary Delete item
// @Description Deletes an item based on the ID provided
// @Tags Items
// @Accept json
// @Produce json
// @Param itemID path string true "ID of the item to be deleted"
// @Success 204
// @Failure 400 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /items/{itemID} [delete]
// @Security KeyAuth
func (h *handler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	logger := logger.NewLogger("delete item")
	logger.Info("start delete item")
	ctx := r.Context()
	itemID := chi.URLParam(r, "itemID")

	ID, err := strconv.Atoi(itemID)
	if err != nil {
		logger.Error("error trying to convert id", err)
		response.SendJSON(w, http.StatusBadRequest, models.ItemErrorParam("id", "missing or invalid parameter").Error())
		return
	}

	if err := h.service.DeleteItemService(ctx, ID); err != nil {
		response.SendJSON(w, err.Code, err)
		return
	}

	logger.Info("Successfully deleted item")
	response.SendJSON(w, http.StatusNoContent, nil)
}
