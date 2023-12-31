package handlers

import (
	"net/http"
	"strconv"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/pkg/errors"
	"github.com/GabriellGds/go-orders/pkg/logger"
	"github.com/GabriellGds/go-orders/pkg/response"
	"github.com/go-chi/chi/v5"
)

// @Summary Find user 
// @Description Retrieves user details based on the user ID
// @Tags Users
// @Accept json
// @Produce json
// @Param userID path string true "ID of the user to be retrieved"
// @Success 200 {object} models.UserResponse "User information retrieved successfully"
// @Failure 400 {object} errors.ErrorResponse "Error: Invalid id"
// @Failure 404 {object} errors.ErrorResponse "User not found"
// @Router /users/{userID} [get]
// @Security KeyAuth
func (h *handler) FindUser(w http.ResponseWriter, r *http.Request) {
	logger := logger.NewLogger("findUser")
	logger.Info("start findUser")
	ctx := r.Context()

	param := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(param)
	if err != nil {
		logger.Error("error trying to convert parameter", err)
		response.SendJSON(w, http.StatusBadRequest, errors.ErrorResponse{
			Message: "invalid user",
		})
		return
	}

	user, err := h.service.FindUserService(ctx, userID)
	if err != nil {
		response.SendJSON(w, http.StatusNotFound, err)
		return
	}

	data := models.NewUserResponse(user.ID, user.Name, user.Email, user.CreatedAt, user.UpdatedAt)

	logger.Info("user found successfully")
	response.SendJSON(w, http.StatusOK, data)
}
