package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/pkg/errors"
	"github.com/GabriellGds/go-orders/pkg/logger"
	"github.com/GabriellGds/go-orders/pkg/response"
)


// @Summary Create item
// @Description Create a new item 
// @Tags Items
// @Accept json
// @Produce json
// @Param request body models.ItemRequest true "request body"
// @Success 201 {object} models.Item
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /items/ [post]
// @Security KeyAuth
func (h *handler) CreateItem(w http.ResponseWriter, r *http.Request) {
	logger := logger.NewLogger("create item")
	logger.Info("start create item")
	ctx := r.Context()
	var i models.ItemRequest
	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		logger.Error("error unmarshal item", err)
		response.SendJSON(w, http.StatusBadRequest, errors.ErrorResponse{
			Message: "invalid type",
		})
		return
	}

	if err := i.Validate(); err != nil {
		logger.Error("error to validate item", err)
		response.SendJSON(w, http.StatusBadRequest, err)
		return
	}

	item := models.NewItem(i.Name, i.Price)
	itemResult, err := h.service.CreateItemService(ctx ,item)
	if err != nil {
		logger.Error("error to create item", err)
		response.SendJSON(w, err.Code, err)
		return
	}

	logger.Info("item created successfully")

	response.SendJSON(w, http.StatusCreated, itemResult)
}
