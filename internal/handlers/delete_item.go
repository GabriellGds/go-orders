package handlers

import (
	"net/http"
	"strconv"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/internal/repository"
	"github.com/GabriellGds/go-orders/pkg/logger"
	"github.com/GabriellGds/go-orders/pkg/response"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	logger := logger.NewLogger("delete item")
	logger.Info("start delete item")

	itemID := chi.URLParam(r, "itemID")

	ID, err := strconv.Atoi(itemID)
	if err != nil {
		logger.Error("error trying to convert id", err)
		response.SendJSON(w, http.StatusBadRequest, models.ItemErrorParam("id", "missing or invalid parameter").Error())
		return
	}

	tokenID, err := models.GetUserIDFromToken(r)
	if err != nil {
		logger.Error("invalid token", err)
		response.SendJSON(w, http.StatusUnauthorized, err)
		return
	}

	if ID != tokenID {
		logger.Info("forbiden")
		response.SendJSON(w, http.StatusForbidden, models.ErrorResponse{
			Message: "Access denied. You do not have permission to perform this action",
		})
		return
	}

	repo := repository.NewItemRepository(h.DB)
	_, err = repo.FindItem(ID)
	if err != nil {
		logger.Error("error to find item on database", err)
		response.SendJSON(w, http.StatusNotFound, models.ErrorResponse{
			Message: "item not found",
		})
		return
	}

	err = repo.DeleteItem(ID)
	if err != nil {
		logger.Error("error to delete item on database", err)
		response.SendJSON(w, http.StatusInternalServerError, models.ErrorResponse{
			Message: "error to delete item on database",
		})
		return
	}

	logger.Info("Successfully deleted item")
	response.SendJSON(w, http.StatusNoContent, nil)
}
