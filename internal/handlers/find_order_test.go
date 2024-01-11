package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/mocks"
	"github.com/go-chi/chi/v5"
	"go.uber.org/mock/gomock"
)

func Test_FindOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mocks.NewMockService(ctrl)
	handler := NewHandler(service)

	t.Run("success", func(t *testing.T){
		rr := httptest.NewRecorder()

		orderID := 7
		order := models.Order{
			ID: 7,
			Items: []models.OrderItems{{ItemID: 1,OrderID: 7,Quantity: 10,Price: 25.0}},
		}

		service.EXPECT().FindOrderService(gomock.Any(), orderID).Return(order, nil)
		req, _ := http.NewRequest("GET", "", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("orderID", strconv.Itoa(orderID))
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

		handler.FindOrder(rr, req)
		if rr.Code != http.StatusOK {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusOK, rr.Code)
		}
		var orderResp models.Order
		if err := json.NewDecoder(rr.Body).Decode(&orderResp); err != nil {
			t.FailNow()
		}

		if orderResp.ID != order.ID {
			t.Errorf("returned wrong id; expected %d, but got %d", order.ID, orderResp.ID)
		}
		if !reflect.DeepEqual(orderResp.Items, order.Items) {
			t.Errorf("returned wrong items; expected %v, but got %v", order.Items, orderResp.Items)
		}
	})

	t.Run("invalid orderID", func(t *testing.T){
		rr := httptest.NewRecorder()

		orderID := "invalido"
		
		req, _ := http.NewRequest("GET", "", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("orderID", orderID)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

		handler.FindOrder(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("service returns an", func(t *testing.T){
		rr := httptest.NewRecorder()

		orderID := 7
		order := models.Order{}
		testError := errors.New("error")

		service.EXPECT().FindOrderService(gomock.Any(), orderID).Return(order, testError)
		req, _ := http.NewRequest("GET", "", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("orderID", strconv.Itoa(orderID))
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

		handler.FindOrder(rr, req)
		if rr.Code != http.StatusNotFound {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusNotFound, rr.Code)
		}
	})

}
