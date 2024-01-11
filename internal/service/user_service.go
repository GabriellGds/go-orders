package service

import (
	"context"
	"net/http"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/pkg/errors"
	"github.com/GabriellGds/go-orders/pkg/logger"
)

func (s *service) CreateUserService(ctx context.Context, user *models.User) (*models.User, *errors.ErrorResponse) {
	u, _ := s.repository.FindUserByEmailRepository(ctx, user.Email)
	if u != nil {
		return nil, &errors.ErrorResponse{
			Field:   "email",
			Message: "Email is already registered in another account",
			Code:    http.StatusBadRequest,
		}
	}

	user.HashPassword(user.Password)
	id, err := s.repository.CreateUserRepository(ctx, *user)
	if err != nil {
		return nil, &errors.ErrorResponse{
			Message: "error to creating user on database",
			Code:    http.StatusInternalServerError,
		}
	}
	user.ID = id

	return user, nil
}

func (s *service) DeleteUserService(ctx context.Context, id int) *errors.ErrorResponse {
	logger := logger.NewLogger("deleteUser service")
	logger.Info("delete user service")
	if _, err := s.repository.FindUserRepository(ctx, id); err != nil {
		logger.Error("error to find user", err)
		return &errors.ErrorResponse{
			Message: "user not found",
			Code:    http.StatusNotFound,
		}
	}

	if err := s.repository.DeleteUserRepository(ctx, id); err != nil {
		logger.Error("error to deleting user on database", err)
		return &errors.ErrorResponse{
			Message: "error to deleting user",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

func (s *service) UpdateUserService(ctx context.Context, id int, user *models.User) *errors.ErrorResponse {
	logger := logger.NewLogger("updateUserService")
	logger.Info("start updateUserService")
	_, err := s.repository.FindUserRepository(ctx, id)
	if err != nil {
		logger.Error("error to finding user ", err)
		return &errors.ErrorResponse{
			Message: "user not found",
			Code:    http.StatusNotFound,
		}
	}
	if err := s.repository.UpdateUserRepository(ctx, id, user); err != nil {
		logger.Error("error to updating user on database", err)
		return &errors.ErrorResponse{
			Message: "error to update user",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

func (s *service) FindUserService(ctx context.Context, id int) (models.User, error) {
	user, err := s.repository.FindUserRepository(ctx, id)
	if err != nil {
		return models.User{}, &errors.ErrorResponse{
			Message: "user not found",
		}
	}

	return user, nil
}

func (s *service) Login(ctx context.Context, user *models.User) (*models.User, string, *errors.ErrorResponse) {
	logger := logger.NewLogger("login service")
	logger.Error("start login service")
	u, err := s.repository.FindUserByEmailRepository(ctx, user.Email)
	if err != nil {
		logger.Error("error to finding email on database", err)
		return nil, "", &errors.ErrorResponse{
			Message: "email not found",
			Code:    http.StatusNotFound,
		}
	}

	err = user.ComparePasswordAndHash(u.Password, user.Password)
	if err != nil {
		logger.Error("invalid password", err)
		return nil, "", &errors.ErrorResponse{
			Message: "invalid password",
			Code:    http.StatusUnauthorized,
		}
	}

	token, err := u.GenerateToken()
	if err != nil {
		logger.Error("error to generate token", err)
		return nil, "", &errors.ErrorResponse{
			Message: "error to generate token",
			Code: http.StatusInternalServerError,
		}
	}

	return u, token, nil
}

func (s *service) ListUsers(ctx context.Context) ([]models.User, *errors.ErrorResponse) {
	logger:= logger.NewLogger("listUser service")
	logger.Info("start listUser service")

	users, err := s.repository.ListUserRepository(ctx)
	if err != nil {
		logger.Error("error to listing users", err)
		return nil, &errors.ErrorResponse{
			Message: "error to listing users",
			Code: http.StatusInternalServerError,
		}
	}

	return users, nil
}