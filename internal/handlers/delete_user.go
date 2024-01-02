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

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	logger := logger.NewLogger("delete user")
	logger.Info("start delete user")

	param := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(param)
	if err != nil {
		logger.Error("error trying to convert parameter", err)
		response.SendJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Message: "invalid id",
		})
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

	repo = repository.NewUserRepository(h.DB)
	err = repo.DeleteUser(userID)
	if err != nil {
		logger.Error("error trying to delete user on database", err)
		response.SendJSON(w, http.StatusInternalServerError, models.ErrorResponse{
			Message: "error to deleting user on database",
		})
		return
	}

	logger.Info("user successfully deleted")
	response.SendJSON(w, http.StatusNoContent, nil)
}
