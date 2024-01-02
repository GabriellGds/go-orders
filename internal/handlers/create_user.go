package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/internal/repository"
	"github.com/GabriellGds/go-orders/pkg/logger"
	"github.com/GabriellGds/go-orders/pkg/response"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	logger := logger.NewLogger("create user handler")
	logger.Info("start create user handler")

	var userRequest models.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		logger.Error("error to unmarshal user", err)
		response.SendJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Message: "invalid type",
		})
		return
	}

	err = userRequest.Validate()
	if err != nil {
		logger.Error("error to validate user", err)
		response.SendJSON(w, http.StatusBadRequest, err)
		return
	}

	u := models.NewUser(userRequest.Name, userRequest.Password, userRequest.Email)
	u.HashPassword(u.Password)

	repo := repository.NewUserRepository(h.DB)
	registeredUser, _ := repo.FindEmail(u.Email)
	if registeredUser != nil {
		logger.Info("Email is already registered in another account")
		response.SendJSON(w, http.StatusBadRequest, models.UserErrorParam("email", "Email is already registered in another account"))
		return
	}

	user, err := repo.CreateUser(*u)
	if err != nil {
		logger.Error("error to create user on database", err)
		response.SendJSON(w, http.StatusInternalServerError, models.ErrorResponse{
			Message: "error to create user",
		})
		return
	}

	userResponse := models.NewUserResponse(user.ID,
		user.Name,
		user.Email,
		user.CreatedAt,
		user.CreatedAt,
		user.DeletedAt,
	)

	logger.Info("user created successfully")
	response.SendJSON(w, http.StatusCreated, userResponse)
}
