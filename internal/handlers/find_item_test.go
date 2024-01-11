package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/mocks"
	"github.com/go-chi/chi/v5"
	"go.uber.org/mock/gomock"
)

func Test_FindItem(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mocks.NewMockService(ctrl)
	handler := NewHandler(service)

	t.Run("success", func(t *testing.T) {
		rr := httptest.NewRecorder()

		itemID := 5
		item := models.Item{
			ID:    itemID,
			Name:  "hot dog",
			Price: 500.75,
		}

		service.EXPECT().FindItemService(gomock.Any(), itemID).Return(item, nil)
		req, _ := http.NewRequest("GET", "", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("itemID", strconv.Itoa(itemID))
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

		handler.FindItem(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusOK, rr.Code)
		}
		var i models.Item
		if err := json.NewDecoder(rr.Body).Decode(&i); err != nil {
			t.FailNow()
		}

		if i.ID != item.ID {
			t.Errorf("returned wrong id; expected %d, but got %d", item.ID, i.ID)
		}
		if i.Name != item.Name {
			t.Errorf("returnerd wrong name; expected %s, but got %s", item.Name, i.Name)
		}
		if i.Price != item.Price {
			t.Errorf("returned wrong price; expected %f, but got %f", item.Price, i.Price)
		}

	})

	t.Run("invalid itemID", func(t *testing.T) {
		rr := httptest.NewRecorder()

		itemID := "nao sou valido"
	
		req, _ := http.NewRequest("GET", "", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("itemID", itemID)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

		handler.FindItem(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("service returns an error", func(t *testing.T) {
		rr := httptest.NewRecorder()

		itemID := 5
		item := models.Item{}
		testErr := errors.New("error")

		service.EXPECT().FindItemService(gomock.Any(), itemID).Return(item, testErr)
		req, _ := http.NewRequest("GET", "", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("itemID", strconv.Itoa(itemID))
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

		handler.FindItem(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusNotFound, rr.Code)
		}
	})
}
