package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/mocks"
	"github.com/GabriellGds/go-orders/pkg/errors"
	"github.com/go-chi/chi/v5"

	"go.uber.org/mock/gomock"
)

func Test_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mocks.NewMockService(ctrl)
	handler := NewHandler(service)

	t.Run("delete user success", func(t *testing.T) {
		rr := httptest.NewRecorder()

		userID := 1
		user := models.User{
			ID: userID,
		}
		token, _ := user.GenerateToken()

		service.EXPECT().DeleteUserService(gomock.Any(), userID).Return(nil)

		req, _ := http.NewRequest("DELETE", "", nil)
		req.Header.Add("Authorization", "Bearer "+token)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("userID", strconv.Itoa(userID))
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

		handler.DeleteUser(rr, req)

		if http.StatusNoContent != rr.Code {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusNoContent, rr.Code)
		}
	})

	t.Run("invalid userID", func(t *testing.T){
		rr := httptest.NewRecorder()

		userID := "lalsdal"
		user := models.User{
			ID: 1,
		}
		token, _ := user.GenerateToken()
		
		req, _ := http.NewRequest("DELETE", "", nil)
		req.Header.Add("Authorization", "Bearer " + token)
		
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("userID", userID)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

		handler.DeleteUser(rr, req)

		if http.StatusBadRequest != rr.Code {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusNoContent, rr.Code)
		}
	})

	t.Run("service returns error", func(t *testing.T){
		rr := httptest.NewRecorder()

		userID := 1
		user := models.User{
			ID: userID,
		}
		token, _ := user.GenerateToken()
		service.EXPECT().DeleteUserService(gomock.Any(), userID).Return(&errors.ErrorResponse{
			Message: "error test",
			Code: http.StatusInternalServerError,
		})

		req, _ := http.NewRequest("DELETE", "", nil)
		req.Header.Add("Authorization", "Bearer " + token)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("userID", strconv.Itoa(userID))
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

		handler.DeleteUser(rr, req)

		if http.StatusInternalServerError != rr.Code {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusInternalServerError, rr.Code)
		}
	})
}
