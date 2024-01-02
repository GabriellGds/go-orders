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

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	logger := logger.NewLogger("find user")
	logger.Info("start find user")

	param := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(param)
	if err != nil {
		logger.Error("error trying to convert parameter", err)
		response.SendJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Message: "user not found",
		})
		return
	}

	var u models.UserUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		logger.Error("error to unmarshal user", err)
		response.SendJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Message: "invalid type",
		})
		return
	}

	err = u.Validate()
	if err != nil {
		logger.Error("error to validate", err)
		response.SendJSON(w, http.StatusBadRequest, err)
		return
	}

	repo := repository.NewUserRepository(h.DB)
	_, err = repo.User(userID)
	if err != nil {
		logger.Error("error trying to find user on database", err)
		response.SendJSON(w, http.StatusNotFound, models.ErrorResponse{
			Message: "user not found",
		})
		return
	}

	err = repo.UpdateUser(userID, u)
	if err != nil {
		logger.Error("error trying to update user on database", err)
		response.SendJSON(w, http.StatusInternalServerError, models.ErrorResponse{
			Message: "error updating user on database",
		})
		return
	}

	logger.Info("user updated successfully")
	response.SendJSON(w, http.StatusNoContent, nil)
}
