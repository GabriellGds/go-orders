package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/mocks"
	"github.com/GabriellGds/go-orders/pkg/errors"
	"go.uber.org/mock/gomock"
)

func Test_ListItems(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mocks.NewMockService(ctrl)
	handler := NewHandler(service)

	t.Run("success", func(t *testing.T) {
		rr := httptest.NewRecorder()

		items := []models.Item{
			{
				ID:    1,
				Name:  "donalds",
				Price: 55.8,
			},
			{
				ID:    2,
				Name:  "limao",
				Price: 27.8,
			},
		}

		service.EXPECT().ListItems(gomock.Any()).Return(items, nil)
		req, _ := http.NewRequest("GET", "", nil)

		handler.ListItems(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusOK, rr.Code)
		}
		var itemResp []models.Item
		if err := json.NewDecoder(rr.Body).Decode(&itemResp); err != nil {
			t.FailNow()
		}
		if len(itemResp) != len(items) {
			t.Errorf("returned wrong size; expected %d, but got %d", len(items), len(itemResp))
		}

		for i := 0; i < len(items) - 1; i++ {
			if itemResp[i].ID != items[i].ID {
				t.Errorf("returned wrong id; expected %d, but got %d", items[i].ID, itemResp[i].ID)
			}
			if itemResp[i].Name != items[i].Name {
				t.Errorf("returned wrong name; expected %s, but got %s", items[i].Name, itemResp[i].Name)
			}
			if itemResp[i].Price != items[i].Price {
				t.Errorf("returned wrong price; expected %f, but got %f", items[i].Price, itemResp[i].Price)
			}
		}
	})

	t.Run("service returns an error", func(t *testing.T){
		rr := httptest.NewRecorder()

		service.EXPECT().ListItems(gomock.Any()).Return(nil, &errors.ErrorResponse{
			Message: "error",
			Code: http.StatusInternalServerError,
		})
		req, _ := http.NewRequest("GET", "", nil)
		handler.ListItems(rr, req)

		if rr.Code != http.StatusInternalServerError {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusInternalServerError, rr.Code)
		}
	})
}
