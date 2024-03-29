package main

import (
	"log"
	"net/http"
	"os"

	"github.com/GabriellGds/go-orders/internal/handlers"
	"github.com/GabriellGds/go-orders/internal/repository"
	"github.com/GabriellGds/go-orders/internal/routes"
	"github.com/GabriellGds/go-orders/internal/service"
	database "github.com/GabriellGds/go-orders/pkg/database/postgres"
	"github.com/GabriellGds/go-orders/pkg/logger"
	"github.com/go-chi/chi/v5"
)


// @title Go Orders API
// @version 1.0
// @description Orders API with authentication
// @contact.name Gabriel Gomes
// @host localhost:5000
// @BasePath /
// @license MIT
// @securityDefinitions.apiKey KeyAuth
// @in header
// @name Authorization
func main() {

	logger := logger.NewLogger("main")
	logger.Info("start service")

	db, err := database.Connect()
	if err != nil {
		logger.Error("error", "database error", err)
		os.Exit(-1)
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handlers.NewHandler(service)

	mux := chi.NewRouter()
	routes.InitRoutes(mux, handler)

	if err := http.ListenAndServe(":5000", mux); err != nil {
		log.Fatal(err)
	}
}
