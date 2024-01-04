package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/pkg/errors"
	"github.com/GabriellGds/go-orders/pkg/logger"
	"github.com/GabriellGds/go-orders/pkg/response"
	"github.com/go-chi/chi/v5"
)

func (h *handler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	logger := logger.NewLogger("update item")
	logger.Info("start update item")
	ctx := r.Context()

	userID := chi.URLParam(r, "itemID")

	ID, err := strconv.Atoi(userID)
	if err != nil {
		logger.Error("error trying to convert id", err)
		response.SendJSON(w, http.StatusBadRequest, models.ItemErrorParam("id", "queryParameter"))
		return
	}

	var i models.ItemRequest
	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		logger.Error("error to unmarshal item", err)
		response.SendJSON(w, http.StatusBadRequest, errors.ErrorResponse{
			Message: "invalid type",
		})
		return
	}

	err = i.Validate()
	if err != nil {
		logger.Error("error to validate item", err)
		response.SendJSON(w, http.StatusBadRequest, err)
		return
	}

	item := models.NewItem(i.Name, i.Price)
	error := h.service.UpdateItemSvice(ctx, ID, item)
	if err != nil {
		response.SendJSON(w, error.Code, err)
		return
	}

	logger.Info("Item Updated Successfully")
	response.SendJSON(w, http.StatusNoContent, nil)
}
