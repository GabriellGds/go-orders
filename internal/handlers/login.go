package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/internal/repository"
	"github.com/GabriellGds/go-orders/pkg/logger"
	"github.com/GabriellGds/go-orders/pkg/response"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	logger := logger.NewLogger("login")
	logger.Info("start login")

	var u models.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		logger.Error("error to unmarshal user", err)
		response.SendJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Message: "invalid type",
		})
		return
	}

	err := u.Validate()
	if err != nil {
		logger.Error("error to validate user", err)
		response.SendJSON(w, http.StatusBadRequest, err)
		return
	}

	repo := repository.NewUserRepository(h.DB)
	user, err := repo.FindEmail(u.Email)
	if err != nil {
		logger.Error("error to finding email on database", err)
		response.SendJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Message: "invalid email",
		})
		return
	}

	err = user.ComparePasswordAndHash(user.Password, u.Password)
	if err != nil {
		logger.Error("invalid password", err)
		response.SendJSON(w, http.StatusUnauthorized, models.ErrorResponse{
			Message: "invalid password",
		})
		return
	}
	userResponse := models.NewUserResponse(user.ID, user.Name, user.Email, user.CreatedAt, user.UpdatedAt, user.DeletedAt)

	token, err := user.GenerateToken()
	if err != nil {
		logger.Error("error to generate token", err)
		response.SendJSON(w, http.StatusInternalServerError, models.ErrorResponse{
			Message: fmt.Sprintf("error to generate token, %s", err.Error()),
		})
		return
	}

	w.Header().Set("Authorization", token)
	response.SendJSON(w, http.StatusOK, userResponse)
}
