package handlers

import (
	"net/http"

	"github.com/GabriellGds/go-orders/internal/service"
)

type HandlerInterface interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	FindUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	ListUsers(w http.ResponseWriter, r *http.Request)

	CreateItem(w http.ResponseWriter, r *http.Request)
	FindItem(w http.ResponseWriter, r *http.Request)
	DeleteItem(w http.ResponseWriter, r *http.Request)
	UpdateItem(w http.ResponseWriter, r *http.Request)
	ListItems(w http.ResponseWriter, r *http.Request)

	CreateOrder(w http.ResponseWriter, r *http.Request)
	FindOrder(w http.ResponseWriter, r *http.Request)
	DeleteOrder(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	service service.Service
}

func NewHandler(service service.Service) HandlerInterface {
	return &handler{service: service}
}
