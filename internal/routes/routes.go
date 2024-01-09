package routes

import (
	"github.com/GabriellGds/go-orders/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	swagger"github.com/swaggo/http-swagger"
	_ "github.com/GabriellGds/go-orders/docs"
)

func InitRoutes(mux *chi.Mux, handler handlers.HandlerInterface) {

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	mux.Use(handlers.CorsMiddleware)
	mux.Post("/login", handler.Login)
	mux.Post("/user", handler.CreateUser)

	mux.Route("/users", func(r chi.Router) {
		r.Use(handlers.Authentication)
		r.Get("/", handler.ListUsers)
		r.Get("/{userID}", handler.FindUser)
		r.Put("/{userID}", handler.UpdateUser)
		r.Delete("/{userID}", handler.DeleteUser)
	})

	mux.Route("/items", func(r chi.Router) {
		r.Use(handlers.Authentication)
		r.Post("/", handler.CreateItem)
		r.Get("/", handler.ListItems)
		r.Get("/{itemID}", handler.FindItem)
		r.Put("/{itemID}", handler.UpdateItem)
		r.Delete("/{itemID}", handler.DeleteItem)
	})

	mux.Route("/orders", func(r chi.Router) {
		r.Use(handlers.Authentication)
		r.Post("/", handler.CreateOrder)
		r.Delete("/{orderID}", handler.DeleteOrder)
		r.Get("/{orderID}", handler.FindOrder)
	})

	mux.Get("/docs/*", swagger.Handler(swagger.URL("http://localhost:5000/docs/doc.json")))

}
