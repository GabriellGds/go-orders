package handlers

import (
	"net/http"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/pkg/response"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)

			return
		}

		next.ServeHTTP(w, r)
	})
}

func Authentication (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		_, err := models.VerifyToken(w, r)
		if err != nil {
			response.SendJSON(w, http.StatusUnauthorized, err)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}

	

	
