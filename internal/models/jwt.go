package models

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	errorResponse "github.com/GabriellGds/go-orders/pkg/errors"
	"github.com/golang-jwt/jwt/v5"
)

var (
	JWT_SECRET = "JWT_SECRET"
)

func (u *User) GenerateToken() (string, error) {
	secret := os.Getenv(JWT_SECRET)

	claims := jwt.MapClaims{
		"userID": u.ID,
		"exp":    time.Now().Add(time.Hour * 6).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(w http.ResponseWriter, r *http.Request) (User, error) {
	secret := os.Getenv(JWT_SECRET)

	tokenValue := RemoveBearerPrefix(r.Header.Get("Authorization"))
	token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, errors.New("invalid token")
	})
	if err != nil {
		return User{}, &errorResponse.ErrorResponse{Message: "invalid token"}
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return User{}, &errorResponse.ErrorResponse{Message: "Invalid token"}
	}

	return User{
		ID: int(claims["userID"].(float64)),
	}, nil
}

func RemoveBearerPrefix(token string) string {
	return strings.TrimPrefix(token, "Bearer ")
}

func GetUserIDFromToken(r *http.Request) (int, error) {
	secret := os.Getenv(JWT_SECRET)
	tokenValue := RemoveBearerPrefix(r.Header.Get("Authorization"))

	token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, &errorResponse.ErrorResponse{Message: "invalid token or claims"}

	}

	userID, ok := claims["userID"].(float64)
	if !ok {
		return 0, &errorResponse.ErrorResponse{Message: "invalid token or claims"}
	}

	return int(userID), nil
}
