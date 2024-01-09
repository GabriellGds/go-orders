package handlers

import (
	"net/http"

	"github.com/GabriellGds/go-orders/pkg/logger"
	"github.com/GabriellGds/go-orders/pkg/response"
)

// @Summary List items
// @Description Retrieves all items
// @Tags Items
// @Accept json
// @Produce json
// @Success 200 {object} models.Item "User information retrieved successfully"
// @Failure 400 {object} errors.ErrorResponse "Error: Invalid id"
// @Failure 404 {object} errors.ErrorResponse "User not found"
// @Router /items/ [get]
// @Security KeyAuth
func (h *handler) ListItems(w http.ResponseWriter, r *http.Request) {
	logger := logger.NewLogger("all items")
	logger.Info("start all items")
	ctx := r.Context()

	items, err := h.service.ListItems(ctx)
	if err != nil {
		response.SendJSON(w, err.Code, err)
		return
	}

	logger.Info("items found successfully")
	response.SendJSON(w, http.StatusOK, items)
}
