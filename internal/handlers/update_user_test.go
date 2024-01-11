package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/mocks"
	"github.com/GabriellGds/go-orders/pkg/errors"
	"github.com/go-chi/chi/v5"
	"go.uber.org/mock/gomock"
)

func Test_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mocks.NewMockService(ctrl)
	handler := NewHandler(service)

	t.Run("success", func(t *testing.T) {
		rr := httptest.NewRecorder()

		userRequest := models.UserUpdateRequest{
			Name: "nmadsfa",
		}
		userID := 2
		user := models.User{
			ID:   userID,
			Name: userRequest.Name,
		}
		token, _ := user.GenerateToken()

		userJson, _ := json.Marshal(userRequest)
		reader := io.NopCloser(strings.NewReader(string(userJson)))
		service.EXPECT().UpdateUserService(gomock.Any(), userID, gomock.Any()).Return(nil)

		req, _ := http.NewRequest("PUT", "", reader)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("userID", strconv.Itoa(userID))
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
		req.Header.Add("Authorization", "Bearer "+token)
		handler.UpdateUser(rr, req)

		if rr.Code != http.StatusNoContent {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusNoContent, rr.Code)
		}
	})

	t.Run("service returns an error", func(t *testing.T) {
		rr := httptest.NewRecorder()

		userRequest := models.UserUpdateRequest{
			Name: "nmadsfa",
		}
		userID := 2
		user := models.User{
			ID:   userID,
			Name: userRequest.Name,
		}
		token, _ := user.GenerateToken()

		userJson, _ := json.Marshal(userRequest)
		reader := io.NopCloser(strings.NewReader(string(userJson)))
		service.EXPECT().UpdateUserService(gomock.Any(), userID, gomock.Any()).Return(&errors.ErrorResponse{
			Message: "error",
			Code: http.StatusInternalServerError,
		})

		req, _ := http.NewRequest("PUT", "", reader)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("userID", strconv.Itoa(userID))
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
		req.Header.Add("Authorization", "Bearer "+token)
		handler.UpdateUser(rr, req)

		if rr.Code != http.StatusInternalServerError {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusInternalServerError, rr.Code)
		}
	})

	t.Run("validation error", func(t *testing.T) {
		rr := httptest.NewRecorder()

		userRequest := models.UserUpdateRequest{
			Name: "",
		}
		userID := 2
		user := models.User{
			ID:   userID,
			Name: userRequest.Name,
		}
		token, _ := user.GenerateToken()

		userJson, _ := json.Marshal(userRequest)
		reader := io.NopCloser(strings.NewReader(string(userJson)))

		req, _ := http.NewRequest("PUT", "", reader)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("userID", strconv.Itoa(userID))
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
		req.Header.Add("Authorization", "Bearer "+token)
		handler.UpdateUser(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("invalid param", func(t *testing.T) {
		rr := httptest.NewRecorder()

		userRequest := models.UserUpdateRequest{
			Name: "dfdfd",
		}
		userID := "fgsfgg"
		user := models.User{
			ID:   1,
			Name: userRequest.Name,
		}
		token, _ := user.GenerateToken()

		userJson, _ := json.Marshal(userRequest)
		reader := io.NopCloser(strings.NewReader(string(userJson)))

		req, _ := http.NewRequest("PUT", "", reader)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("userID", userID)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
		req.Header.Add("Authorization", "Bearer "+token)
		handler.UpdateUser(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("acesss denied", func(t *testing.T) {
		rr := httptest.NewRecorder()

		userRequest := models.UserUpdateRequest{
			Name: "nmadsfa",
		}
		userID := 2
		user := models.User{
			ID:   3,
			Name: userRequest.Name,
		}
		token, _ := user.GenerateToken()

		userJson, _ := json.Marshal(userRequest)
		reader := io.NopCloser(strings.NewReader(string(userJson)))

		req, _ := http.NewRequest("PUT", "", reader)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("userID", strconv.Itoa(userID))
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
		req.Header.Add("Authorization", "Bearer "+token)
		handler.UpdateUser(rr, req)

		if rr.Code != http.StatusForbidden {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusForbidden, rr.Code)
		}
	})
}
