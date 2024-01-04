package handlers

import (
	"net/http"

	"github.com/GabriellGds/go-orders/pkg/logger"
	"github.com/GabriellGds/go-orders/pkg/response"
)

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
