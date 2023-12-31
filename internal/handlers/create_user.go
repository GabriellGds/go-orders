package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/pkg/errors"
	"github.com/GabriellGds/go-orders/pkg/logger"
	"github.com/GabriellGds/go-orders/pkg/response"
)


// CreateUser Creates a new user
// @Summary Create user
// @Description Create a new user 
// @Tags Users
// @Accept json
// @Produce json
// @Param request body models.UserRequest true "request body"
// @Success 201 {object} models.UserResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /user [post]
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	logger := logger.NewLogger("create user handler")
	logger.Info("start create user handler")
	ctx := r.Context()

	var userRequest models.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		logger.Error("error to unmarshal user", err)
		response.SendJSON(w, http.StatusBadRequest, errors.ErrorResponse{
			Message: "invalid type",
		})
		return
	}

	if err := userRequest.Validate(); err != nil {
		logger.Error("error to validate user", err)
		response.SendJSON(w, err.Code, err)
		return
	}

	u := models.NewUser(userRequest.Name, userRequest.Password, userRequest.Email)
	user, err := h.service.CreateUserService(ctx, u)
	if err != nil {
		response.SendJSON(w, err.Code, err)
		return
	}

	userResponse := models.NewUserResponse(user.ID,
		user.Name,
		user.Email,
		user.CreatedAt,
		user.CreatedAt,
	)

	logger.Info("user created successfully")
	response.SendJSON(w, http.StatusCreated, userResponse)
}
