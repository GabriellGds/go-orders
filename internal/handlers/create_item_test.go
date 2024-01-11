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
	"github.com/GabriellGds/go-orders/pkg/errors"
	"go.uber.org/mock/gomock"
)

func Test_CreateItem(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mocks.NewMockService(ctrl)
	handler := NewHandler(service)

	t.Run("service returns success", func(t *testing.T) {
		rr := httptest.NewRecorder()

		ItemRequest := models.ItemRequest{
			Name:  "hamburguer",
			Price: 15.50,
		}

		item := models.NewItem(ItemRequest.Name, ItemRequest.Price)
		service.EXPECT().CreateItemService(gomock.Any(), gomock.Any()).Return(
			item, nil)

		i, _ := json.Marshal(ItemRequest)
		reader := io.NopCloser(strings.NewReader(string(i)))

		req := httptest.NewRequest("POST", "/items/", reader)
		handler.CreateItem(rr, req)

		if http.StatusCreated != rr.Code {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusCreated, rr.Code)
		}

		var resp models.Item
		if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
			t.FailNow()
		}

		if item.Name != resp.Name {
			t.Errorf("returned wrong name; expected %s, but got %s", item.Name, resp.Name)
		}
		if item.Price != resp.Price {
			t.Errorf("returned wrong price; expected %f, but got %f", item.Price, resp.Price)
		}
	})

	t.Run("validation error", func(t *testing.T){
		rr := httptest.NewRecorder()

		itemRequest := models.ItemRequest{
			Name: "",
			Price: 4,
		}

		item, _ := json.Marshal(itemRequest)
		reader := io.NopCloser(strings.NewReader(string(item)))

		req := httptest.NewRequest("POST", "/items/", reader)
		handler.CreateItem(rr, req)

		if http.StatusBadRequest != rr.Code {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("services returns error", func(t *testing.T){
		rr := httptest.NewRecorder()

		itemRequest := models.ItemRequest{
			Name: "seu joao",
			Price: 15.78,
		}

		service.EXPECT().CreateItemService(gomock.Any(), gomock.Any()).Return(
			nil, &errors.ErrorResponse{Message: "error test", Code: http.StatusInternalServerError},
		)

		item, _ := json.Marshal(itemRequest)
		reader := io.NopCloser(strings.NewReader(string(item)))

		req := httptest.NewRequest("POST", "/items/", reader)
		handler.CreateItem(rr, req)

		if rr.Code != http.StatusInternalServerError {
			t.Errorf("CreateItem returned wrong status code; expected %d, but got %d",
			http.StatusInternalServerError, rr.Code)
		}
	})

}
