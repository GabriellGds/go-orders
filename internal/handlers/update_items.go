package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/internal/repository"
	"github.com/GabriellGds/go-orders/pkg/logger"
	"github.com/GabriellGds/go-orders/pkg/response"
	"github.com/go-chi/chi/v5"
)


func (h *Handler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	logger := logger.NewLogger("update item")
	logger.Info("start update item")

	userID := chi.URLParam(r, "itemID")

	ID, err := strconv.Atoi(userID)
	if err != nil {
		logger.Error("error trying to convert id", err)
		response.SendJSON(w, http.StatusBadRequest, models.ItemErrorParam("id", "queryParameter"))
		return
	}

	var item models.ItemRequest
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		logger.Error("error to unmarshal item", err)
		response.SendJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Message: "invalid type",
		})
		return
	}

	err = item.Validate()
	if err != nil {
		logger.Error("error trying to validate item", err)
		response.SendJSON(w, http.StatusBadRequest, err)
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

	err = repo.UpdateItem(ID, item)
	if err != nil {
		logger.Error("error to update item on database", err)
		response.SendJSON(w, http.StatusInternalServerError, models.ErrorResponse{
			Message: "error to updating item on database",
		})
		return
	}

	logger.Info("Item Updated Successfully")
	response.SendJSON(w, http.StatusNoContent, nil)
}
