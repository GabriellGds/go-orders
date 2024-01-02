package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/internal/repository"
	"github.com/GabriellGds/go-orders/pkg/logger"
	"github.com/GabriellGds/go-orders/pkg/response"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	DB *sqlx.DB
}

func (h *Handler) CreateItem(w http.ResponseWriter, r *http.Request) {
	logger := logger.NewLogger("create item")
	logger.Info("start create item")

	var i models.ItemRequest
	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		logger.Error("error unmarshal item", err)
		response.SendJSON(w, http.StatusBadRequest, models.ErrorResponse{
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
	repo := repository.NewItemRepository(h.DB)
	itemResponse, err := repo.CreateItem(*item)
	if err != nil {
		logger.Error("error create item on database", err)
		response.SendJSON(w, http.StatusInternalServerError, models.ErrorResponse{
			Message: "error creating user on database ",
		})
		return
	}

	logger.Info("item created successfully")

	response.SendJSON(w, http.StatusCreated, itemResponse)
}
