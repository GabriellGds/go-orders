package handlers

import (
	"net/http"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/internal/repository"
	"github.com/GabriellGds/go-orders/pkg/logger"
	"github.com/GabriellGds/go-orders/pkg/response"
)

func (h *Handler) AllItems(w http.ResponseWriter, r *http.Request) {
	logger := logger.NewLogger("all items")
	logger.Info("start all items")

	repo := repository.NewItemRepository(h.DB)
	items, err := repo.AllItems()
	if err != nil {
		logger.Error("error to find items on database")
		response.SendJSON(w, http.StatusInternalServerError, models.ErrorResponse{
			Message: "error finding items on database",
		})
		return
	}

	logger.Info("items found successfully")
	response.SendJSON(w, http.StatusOK, items)
}
