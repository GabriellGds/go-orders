package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/GabriellGds/go-orders/pkg/errors"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/mocks"
	"go.uber.org/mock/gomock"
)

func Test_CreateUser(t *testing.T) {
	// tests := []struct {
	// 	name        string
	// 	requestBody string
	// 	statusCode  int
	// }{
	// 	{"valid user", `{"name":"gabriel","email":"gabriel@test.com","password":"12345678"}`, http.StatusCreated},
	// 	{"empty user", `{"email":"gabriel@test.com","password":"12345678"}`, http.StatusBadRequest},
	// 	{"empty password", `{"name":"gabriel","email":"gabriel@test.com"}`, http.StatusBadRequest},
	// 	{"valid user", `{"name":"gabriel","email":"gabriel@test.com","password":"12345678"}`, http.StatusCreated},
	// 	{"empty email", `{"name":"gabriel","password":"12345678"}`, http.StatusBadRequest},
	// 	{"invalid email", `{"name":"gabriel","email":"gabrieltest.com","password":"12345678"}`, http.StatusBadRequest},
	// 	{"empty json", `{}`, http.StatusBadRequest},
	// 	{"not JSON", `{name:gabrie,email:gabriel@test.com,password:12345678}`, http.StatusBadRequest},
	// }
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mocks.NewMockService(ctrl)
	handler := NewHandler(service)

	t.Run("service returns success", func(t *testing.T) {
		rr := httptest.NewRecorder()

		userRequest := models.UserRequest{
			Name:     "gabriel",
			Email:    "gabriel@test.com",
			Password: "12345678",
		}

		user := models.NewUser(userRequest.Name, userRequest.Password, userRequest.Email)

		service.EXPECT().CreateUserService(gomock.Any(), gomock.Any()).Return(
			user, nil)

		u, _ := json.Marshal(userRequest)
		reader := io.NopCloser(strings.NewReader(string(u)))

		req := httptest.NewRequest("POST", "/user", reader)
		handler.CreateUser(rr, req)

		if http.StatusCreated != rr.Code {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusCreated, rr.Code)
		}

		var resp models.User
		if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
			t.FailNow()
		}

		if user.Name != resp.Name {
			t.Errorf("returned wrong name; expected %s, but got %s", user.Name, resp.Name)
		}
		if user.Email != resp.Email {
			t.Errorf("returned wrong email; expected %s, but got %s", user.Email, resp.Email)
		}

	})

	t.Run("validation error", func(t *testing.T) {
		rr := httptest.NewRecorder()
		userRequest := models.UserRequest{
			Name:     "gabriel",
			Email:    "gabrcom",
			Password: "12378",
		}

		u, _ := json.Marshal(userRequest)
		reader := io.NopCloser(strings.NewReader(string(u)))

		req := httptest.NewRequest("POST", "/user", reader)
		handler.CreateUser(rr, req)

		if http.StatusBadRequest != rr.Code {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusInternalServerError, rr.Code)
		}
	})

	t.Run("service returns error", func(t *testing.T) {
		rr := httptest.NewRecorder()

		userRequest := models.UserRequest{
			Name:     "gabriel",
			Email:    "gabriel@test.com",
			Password: "12345678",
		}

		service.EXPECT().CreateUserService(gomock.Any(), gomock.Any()).Return(
			nil, &errors.ErrorResponse{Message: "error test", Code: http.StatusInternalServerError})

		u, _ := json.Marshal(userRequest)
		reader := io.NopCloser(strings.NewReader(string(u)))

		req := httptest.NewRequest("POST", "/user", reader)
		handler.CreateUser(rr, req)

		if http.StatusInternalServerError != rr.Code {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusBadRequest, rr.Code)
		}

	})

}
