package handlers

import (
	"net/http"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/pkg/logger"
	"github.com/GabriellGds/go-orders/pkg/response"
)

func (h handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	logger := logger.NewLogger("list user")
	logger.Info("start handler listUser")
	ctx := r.Context()

	users, err := h.service.ListUsers(ctx)
	if err != nil {
		response.SendJSON(w, err.Code, err)
		return
	}
	
	var userRespose []models.UserResponse
	for _, u := range users {
		user := models.NewUserResponse(u.ID, u.Name, u.Email, u.CreatedAt, u.UpdatedAt)
		userRespose = append(userRespose, *user)
	}


	logger.Info("found users successfuly")
	response.SendJSON(w, http.StatusOK, userRespose)
}