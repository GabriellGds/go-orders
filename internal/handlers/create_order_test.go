package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/mocks"
	"go.uber.org/mock/gomock"
)

func Test_CreateOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mocks.NewMockService(ctrl)
	handler := NewHandler(service)

	t.Run("create order success", func(t *testing.T) {
		rr := httptest.NewRecorder()

		orderItems := []models.OrderItem{
			{
				ItemID:   1,
				Quantity: 2,
			},
		}

		orderRequest := models.OrderRequest{
			Items: orderItems,
		}

		order := models.NewOrder(1, models.OrderRequestToOrderItems(orderRequest))

		service.EXPECT().CreateOrderService(gomock.Any(), gomock.Any()).Return(order, nil)

		o, _ := json.Marshal(orderRequest)
		reader := io.NopCloser(strings.NewReader(string(o)))
		user := models.User{
			ID: 1,
		}
		token, _ := user.GenerateToken()

		req, _ := http.NewRequest("POST", "", reader)
		req.Header.Add("Authorization", "Bearer " + token)
		handler.CreateOrder(rr, req)

		if http.StatusCreated != rr.Code {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusCreated, rr.Code)
		}

		var resp models.OrderCreatedResponse
		if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
			t.FailNow()
		}

		if resp.ID != order.ID {
			t.Errorf("returned wrong order ID; expected %d, but got %d", order.ID, resp.ID)
		}
	})

}
