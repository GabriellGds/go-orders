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

func (h *Handler) FindUser(w http.ResponseWriter, r *http.Request) {
	logger := logger.NewLogger("findUser")
	logger.Info("start findUser")

	param := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(param)
	if err != nil {
		logger.Error("error trying to convert parameter", err)
		response.SendJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Message: "invalid user",
		})
		return
	}

	repo := repository.NewUserRepository(h.DB)
	user, err := repo.User(userID)
	if err != nil {
		logger.Error("error trying to find user on database", err)
		response.SendJSON(w, http.StatusNotFound, models.ErrorResponse{
			Message: "user not found",
		})
		return
	}

	data := models.NewUserResponse(user.ID, user.Name, user.Email, user.CreatedAt, user.UpdatedAt, user.DeletedAt)

	logger.Info("user found successfully")
	response.SendJSON(w, http.StatusOK, data)
}
