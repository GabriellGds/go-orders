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
// @Summary Delete user
// @Description Deletes a user based on the ID provided
// @Tags Users
// @Accept json
// @Produce json
// @Param userID path string true "ID of the user to be deleted"
// @Success 204
// @Failure 400 {object} errors.ErrorResponse
// @Failure 403 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /users/{userID} [delete]
// @Security KeyAuth
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	logger := logger.NewLogger("delete user")
	logger.Info("start delete user")
	ctx := r.Context()
	
	tokenID, err := models.GetUserIDFromToken(r)
	if err != nil {
		logger.Error("invalid token")
		response.SendJSON(w, http.StatusUnauthorized, errors.ErrorResponse{
			Message: "invalid token",
		})
		return
	}
	logger.Info(tokenID)

	param := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(param)
	if err != nil {
		logger.Error("error trying to convert parameter", err)
		response.SendJSON(w, http.StatusBadRequest, errors.ErrorResponse{
			Message: "invalid id",
		})
		return
	}

	if userID != tokenID {
		logger.Error("forbiden, user not have permission")
		response.SendJSON(w, http.StatusForbidden, errors.ErrorResponse{
			Message: "access denied. You do not have permission to perform this action" ,
		})
		return
	}

	if err := h.service.DeleteUserService(ctx, userID); err != nil {
		response.SendJSON(w, err.Code, err)
		return
	}

	logger.Info("user successfully deleted")
	response.SendJSON(w, http.StatusNoContent, nil)
}
