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

func Test_UpdateItems(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mocks.NewMockService(ctrl)
	handler := NewHandler(service)

	t.Run("success", func(t *testing.T) {
		rr := httptest.NewRecorder()

		itemRequest := models.ItemRequest{
			Name:  "testando",
			Price: 258.5,
		}
		itemJson, _ := json.Marshal(itemRequest)
		reader := io.NopCloser(strings.NewReader(string(itemJson)))
		itemID := 5

		service.EXPECT().UpdateItemSvice(gomock.Any(), itemID, gomock.Any()).Return(nil)
		req, _ := http.NewRequest("PUT", "", reader)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("itemID", strconv.Itoa(itemID))
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
		handler.UpdateItem(rr, req)

		if rr.Code != http.StatusNoContent {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusNoContent, rr.Code)
		}
	})

	t.Run("service returns an error", func(t *testing.T) {
		rr := httptest.NewRecorder()

		itemRequest := models.ItemRequest{
			Name:  "testando",
			Price: 258.5,
		}
		itemJson, _ := json.Marshal(itemRequest)
		reader := io.NopCloser(strings.NewReader(string(itemJson)))
		itemID := 5

		service.EXPECT().UpdateItemSvice(gomock.Any(), itemID, gomock.Any()).Return(&errors.ErrorResponse{
			Message: "error",
			Code: http.StatusInternalServerError,
		})
		req, _ := http.NewRequest("PUT", "", reader)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("itemID", strconv.Itoa(itemID))
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
		handler.UpdateItem(rr, req)

		if rr.Code != http.StatusInternalServerError {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusInternalServerError, rr.Code)
		}
	})
	
	t.Run("invalid request", func(t *testing.T){
		rr := httptest.NewRecorder()

		itemRequest := models.ItemRequest{
			Name:  "",
			Price: 4,
		}
		itemJson, _ := json.Marshal(itemRequest)
		reader := io.NopCloser(strings.NewReader(string(itemJson)))
		
		req, _ := http.NewRequest("PUT", "", reader)
		handler.UpdateItem(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusBadRequest, rr.Code)
		}

	})
}
