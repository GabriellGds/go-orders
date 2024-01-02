package routes

import (
	"github.com/GabriellGds/go-orders/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
)

func InitRoutes(mux *chi.Mux, db *sqlx.DB) {
	handler := handlers.Handler{DB: db}

	mux.Use(middleware.Recoverer)
	mux.Post("/login", handler.Login)
	mux.Post("/user", handler.CreateUser)


	mux.Route("/users", func(r chi.Router) {
		r.Use(handlers.Authentication)
		r.Get("/{userID}", handler.FindUser)
		r.Put("/{userID}", handler.UpdateUser)
		r.Delete("/{userID}", handler.DeleteUser)
	})

	mux.Route("/orders", func(r chi.Router) {
		r.Use(handlers.Authentication)
		r.Post("/", handler.CreateItem)
		r.Get("/", handler.AllItems)
		r.Get("/{itemID}", handler.FindItem)
		r.Put("/{itemID}", handler.UpdateItem)
		r.Delete("/{itemID}", handler.DeleteItem)
	})

}
