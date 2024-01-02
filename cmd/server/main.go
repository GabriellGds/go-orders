package main

import (
	"log"
	"net/http"
	"os"

	"github.com/GabriellGds/go-orders/internal/routes"
	database "github.com/GabriellGds/go-orders/pkg/database/postgres"
	"github.com/GabriellGds/go-orders/pkg/logger"
	"github.com/go-chi/chi/v5"
)

func main() {

	logger := logger.NewLogger("main")
	logger.Info("start service")

	db, err := database.Connect()
	if err != nil {
		logger.Error("error", "database error", err)
		os.Exit(-1)
	}

	mux := chi.NewRouter()
	routes.InitRoutes(mux, db)

	if err := http.ListenAndServe(":5000", mux); err != nil {
		log.Fatal(err)
	}
}
