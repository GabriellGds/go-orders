package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/pkg/errors"
	"github.com/GabriellGds/go-orders/pkg/logger"
	"github.com/GabriellGds/go-orders/pkg/response"
)


// @Summary User Login
// @Description Allows a user to log in and receive an authentication token.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body models.UserLogin true "User login credentials"
// @Success 200 {object} models.UserResponse "Login successful"
// @Header 200 {string} Authorization "Authentication token"
// @Failure 403 {object} errors.ErrorResponse "Error: Invalid login credentials"
// @Router /login [post]
func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	logger := logger.NewLogger("login")
	logger.Info("start login")
	ctx := r.Context()

	var u models.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		logger.Error("error to unmarshal user", err)
		response.SendJSON(w, http.StatusBadRequest, errors.ErrorResponse{
			Message: "invalid type",
		})
		return
	}

	if err := u.Validate(); err != nil {
		logger.Error("error to validate user", err)
		response.SendJSON(w, http.StatusBadRequest, err)
		return
	}

	user := models.ConvertLoginToUser(u.Email, u.Password)
	user, token, err := h.service.Login(ctx, user)
	if err != nil {
		response.SendJSON(w, err.Code, err)
		return
	}
	w.Header().Set("Authorization", token)
	userResponse := models.NewUserResponse(user.ID, user.Name, user.Email, user.CreatedAt, user.UpdatedAt)

	response.SendJSON(w, http.StatusOK, userResponse)
}
