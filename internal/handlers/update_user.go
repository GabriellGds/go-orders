package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/pkg/errors"
	"github.com/GabriellGds/go-orders/pkg/logger"
	"github.com/GabriellGds/go-orders/pkg/response"
	"github.com/go-chi/chi/v5"
)

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	logger := logger.NewLogger("find user")
	logger.Info("start find user")
	ctx := r.Context()

	tokenID, err := models.GetUserIDFromToken(r)
	if err != nil {
		logger.Error("invalid token")
		response.SendJSON(w, http.StatusUnauthorized, errors.ErrorResponse{
			Message: "invalid token",
		})
		return
	}

	param := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(param)
	if err != nil {
		logger.Error("error trying to convert parameter", err)
		response.SendJSON(w, http.StatusBadRequest, errors.ErrorResponse{
			Message: "user not found",
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

	var u models.UserUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		logger.Error("error to unmarshal user", err)
		response.SendJSON(w, http.StatusBadRequest, errors.ErrorResponse{
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

	user := models.ConvertUpdateUserToUser(u.Name)
	if err := h.service.UpdateUserService(ctx, userID, user); err != nil {
		response.SendJSON(w, err.Code, err)
		return
	}

	logger.Info("user updated successfully")
	response.SendJSON(w, http.StatusNoContent, nil)
}
