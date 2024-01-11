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

func Test_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mocks.NewMockService(ctrl)
	handler := NewHandler(service)

	t.Run("success", func(t *testing.T){
		rr := httptest.NewRecorder()

		login := models.UserLogin{
			Email: "sla@gmail.com",
			Password: "12345678",
		}
		token := "souUmToken"
		user := models.ConvertLoginToUser(login.Email, login.Password)
		service.EXPECT().Login(gomock.Any(), user).Return(user, token, nil)

		loginJson, _ := json.Marshal(login)
		reader := io.NopCloser(strings.NewReader(string(loginJson)))
		req, _ := http.NewRequest("POST", "", reader)
		
		handler.Login(rr, req)
		if rr.Code != http.StatusOK {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusOK, rr.Code)
		}		
		tokenResp := rr.Header().Values("Authorization")[0]
		if 	tokenResp != token {
			t.Errorf("returned wrong value; expected %s, but got %s",token, tokenResp)
		}
	})

	t.Run("service returns an error", func(t *testing.T){
		rr := httptest.NewRecorder()

		login := models.UserLogin{
			Email: "sla@gmail.com",
			Password: "12345678",
		}
		token := "souUmToken"
		
		user := models.ConvertLoginToUser(login.Email, login.Password)
		service.EXPECT().Login(gomock.Any(), user).Return(user, token, &errors.ErrorResponse{
			Message: "error",
			Code: http.StatusInternalServerError,
		})

		loginJson, _ := json.Marshal(login)
		reader := io.NopCloser(strings.NewReader(string(loginJson)))
		req, _ := http.NewRequest("POST", "", reader)
		
		handler.Login(rr, req)
		if rr.Code != http.StatusInternalServerError {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusInternalServerError, rr.Code)
		}		
		
	})

	t.Run("validation error", func(t *testing.T){
		rr := httptest.NewRecorder()

		login := models.UserLogin{
			Email: "slagmail.com",
			Password: "12678",
		}		

		loginJson, _ := json.Marshal(login)
		reader := io.NopCloser(strings.NewReader(string(loginJson)))
		req, _ := http.NewRequest("POST", "", reader)
		
		handler.Login(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusBadRequest, rr.Code)
		}		
	})
}
