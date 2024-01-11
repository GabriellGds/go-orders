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

func Test_ListUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mocks.NewMockService(ctrl)
	handler := NewHandler(service)

	t.Run("success", func(t *testing.T){
		rr := httptest.NewRecorder()

		users := []models.User{
			{ID: 1, Name: "jao", Email: "jao@gmail.com", Password: "marquinho"},
			{ID: 2, Name: "marocs", Email: "marcos@gmail.com", Password: "mardofoo"},
			{ID: 3, Name: "mbape", Email: "mbape@gmail.com", Password: "mbape1123"},
			{ID: 4, Name: "messi", Email: "messi@gmail.com", Password: "messe134"},
		}
		service.EXPECT().ListUsers(gomock.Any()).Return(users, nil)

		req, _ := http.NewRequest("GET", "", nil)
		handler.ListUsers(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusOK, rr.Code)
		}
		var usersResp []models.User
		if err := json.NewDecoder(rr.Body).Decode(&usersResp); err != nil {
			t.FailNow()
		}

		if len(users) != len(usersResp) {
			t.Errorf("returned wrong size; expected %d, but got %d", len(usersResp), len(users))
		}
		
		for i := 0; i < len(users) - 1; i++ {
			if users[i].ID != usersResp[i].ID {
				t.Errorf("returned wrong id; expected %d, but got %d", users[i].ID, usersResp[i].ID)
			}
			if users[i].Name != usersResp[i].Name {
				t.Errorf("returned wrong name; expected %s, but got %s", users[i].Name, usersResp[i].Name)
			}
			if users[i].Email != usersResp[i].Email {
				t.Errorf("returned wrong email; expected %s, but got %s", users[i].Email, usersResp[i].Email)
			}
		}
	})

	t.Run("services returns an error", func(t *testing.T){
		rr := httptest.NewRecorder()

		service.EXPECT().ListUsers(gomock.Any()).Return(nil, &errors.ErrorResponse{
			Message: "error",
			Code: http.StatusInternalServerError,
		})

		req, _ := http.NewRequest("GET", "", nil)
		handler.ListUsers(rr, req)

		if rr.Code != http.StatusInternalServerError {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusInternalServerError, rr.Code)
		}
	})
}
