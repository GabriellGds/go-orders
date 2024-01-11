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

func Test_FindUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mocks.NewMockService(ctrl)
	handler := NewHandler(service)

	t.Run("success", func(t *testing.T){
		rr := httptest.NewRecorder()

		userID := 1
		user := models.User{
			ID: userID,
			Name: "senhor test",
			Email: "test@gmail.com",
			Password: "mr banana",
		}

		service.EXPECT().FindUserService(gomock.Any(), userID).Return(user, nil)
		
		req, _ := http.NewRequest("GET", "", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("userID", strconv.Itoa(userID))
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

		handler.FindUser(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusOK, rr.Code)
		}

		var u models.User
		if err := json.NewDecoder(rr.Body).Decode(&u); err != nil {
			t.FailNow()
		}

		if u.Name != user.Name {
			t.Errorf("returned wrong name; expected %s, but got %s", user.Name, u.Name)
		}
		if u.Email != user.Email {
			t.Errorf("returned wrong email; expected %s, but got %s", user.Email, u.Email)
		}
				
	})

	t.Run("invalid userID", func(t *testing.T){
		rr := httptest.NewRecorder()

		userID := "fdafdasfsa"
		
		req, _ := http.NewRequest("GET", "", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("userID",userID)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

		handler.FindUser(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusBadRequest, rr.Code)
		}	
	})

	t.Run("service returns an error", func(t *testing.T){
		rr := httptest.NewRecorder()

		userID := 1
		user := models.User{}
		errorTest := errors.New("error")

		service.EXPECT().FindUserService(gomock.Any(), userID).Return(user, errorTest)
		
		req, _ := http.NewRequest("GET", "", nil)
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("userID", strconv.Itoa(userID))
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

		handler.FindUser(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("returned wrong status code; expected %d, but got %d", http.StatusNotFound, rr.Code)
		}
	})

}
