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

func Test_DeleteOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mocks.NewMockService(ctrl)
	handler := NewHandler(service)

	t.Run("success", func(t *testing.T) {
		rr := httptest.NewRecorder()

		orderID := 1
		user := models.User{
			ID: 1,
		}
		token, _ := user.GenerateToken()

		service.EXPECT().DeleteOrderService(gomock.Any(), gomock.Any(), orderID).Return(nil)

		req, _ := http.NewRequest("DELETE", "", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("orderID", strconv.Itoa(orderID))
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
		req.Header.Add("Authorization", "Bearer "+token)

		handler.DeleteOrder(rr, req)

		if rr.Code != http.StatusNoContent {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusNoContent, rr.Code)
		}
	})

	t.Run("invalid orderID", func(t *testing.T) {
		rr := httptest.NewRecorder()

		orderID := "pedidoInvalido"
		user := models.User{
			ID: 1,
		}
		token, _ := user.GenerateToken()

		req, _ := http.NewRequest("DELETE", "", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("orderID", orderID)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
		req.Header.Add("Authorization", "Bearer "+token)

		handler.DeleteOrder(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("services return an error", func(t *testing.T){
		rr := httptest.NewRecorder()

		user := models.User{
			ID: 1,
		}
		token, _ := user.GenerateToken()
		orderID := 1

		service.EXPECT().DeleteOrderService(gomock.Any(), gomock.Any(), orderID).Return(&errors.ErrorResponse{
			Message: "error",
			Code: http.StatusInternalServerError,
		})

		req := httptest.NewRequest("DELETE", "/orders/{orderID}", nil)
		req.Header.Add("Authorization", "Bearer " + token)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("orderID", strconv.Itoa(orderID))
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

		handler.DeleteOrder(rr, req)

		if rr.Code != http.StatusInternalServerError {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusInternalServerError, rr.Code)
		}
	})

	t.Run("invalid token", func(t *testing.T){
		rr := httptest.NewRecorder()

		token := "liopi9po29-8if290f"
		orderID := 1

		req := httptest.NewRequest("DELETE", "/orders/{orderID}", nil)
		req.Header.Add("Authorization", "Bearer " + token)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("orderID", strconv.Itoa(orderID))
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

		handler.DeleteOrder(rr, req)

		if rr.Code != http.StatusUnauthorized {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusUnauthorized, rr.Code)
		}
	})
}
