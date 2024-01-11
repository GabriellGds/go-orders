package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/GabriellGds/go-orders/mocks"
	"github.com/GabriellGds/go-orders/pkg/errors"
	"github.com/go-chi/chi/v5"
	"go.uber.org/mock/gomock"
)

func Test_DeleteItem(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mocks.NewMockService(ctrl)
	handler := NewHandler(service)

	t.Run("success", func(t *testing.T){
		rr := httptest.NewRecorder()
		itemID := 1

		service.EXPECT().DeleteItemService(gomock.Any(), itemID).Return(nil)

		req, _ := http.NewRequest("DELETE", "", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("itemID", strconv.Itoa(itemID))
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

		handler.DeleteItem(rr, req)
		
		if rr.Code != http.StatusNoContent {
			t.Errorf("returned wrong status code; expected %d ,but got %d", http.StatusNoContent, rr.Code)
		}
	})

	t.Run("invalid itemID", func(t *testing.T){
		rr := httptest.NewRecorder()

		itemID := "invalido"

		req, _ := http.NewRequest("DELETE", "", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("itemID", itemID)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

		handler.DeleteItem(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("services returns an error", func(t *testing.T){
		rr := httptest.NewRecorder()
		itemID := 1

		service.EXPECT().DeleteItemService(gomock.Any(), itemID).Return(&errors.ErrorResponse{
			Message: "error",
			Code: http.StatusInternalServerError,
		})

		req, _ := http.NewRequest("DELETE", "", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("itemID", strconv.Itoa(itemID))
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

		handler.DeleteItem(rr, req)

		if rr.Code != http.StatusInternalServerError {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusInternalServerError, rr.Code)
		}
	})
}
