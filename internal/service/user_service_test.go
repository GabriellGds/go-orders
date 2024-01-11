package service

import (
	"context"
	"net/http"

	"testing"

	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/GabriellGds/go-orders/mocks"
	"github.com/GabriellGds/go-orders/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_CreateUserService(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repository := mocks.NewMockRepositoryInterface(ctl)
	service := NewService(repository)

	t.Run("success", func(t *testing.T) {
		user := &models.User{
			ID:       1,
			Name:     "test",
			Email:    "email@eail.com",
			Password: "22456789954",
		}
		ctx := context.Background()
		repository.EXPECT().FindUserByEmailRepository(ctx, user.Email).Return(nil, nil)
		repository.EXPECT().CreateUserRepository(ctx, gomock.Any()).Return(1, nil)

		u, err := service.CreateUserService(ctx, user)

		assert.Nil(t, err)
		assert.EqualValues(t, user.Name, u.Name)
		assert.EqualValues(t, user.ID, u.ID)
		assert.EqualValues(t, user.Password, u.Password)
	})

	t.Run("email already exists", func(t *testing.T) {
		user := &models.User{
			Name:     "test",
			Email:    "email@eail.com",
			Password: "22456789954",
		}
		ctx := context.Background()
		repository.EXPECT().FindUserByEmailRepository(ctx, user.Email).Return(user, nil)

		user, err := service.CreateUserService(ctx, user)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Email is already registered in another account")
	})

	t.Run("error create user", func(t *testing.T) {
		user := &models.User{
			ID:       1,
			Name:     "test",
			Email:    "email@eail.com",
			Password: "22456789954",
		}
		ctx := context.Background()
		repository.EXPECT().FindUserByEmailRepository(ctx, user.Email).Return(nil, nil)
		repository.EXPECT().CreateUserRepository(ctx, gomock.Any()).Return(0, &errors.ErrorResponse{
			Message: "error to creating user on database",
			Code:    http.StatusInternalServerError,
		})

		user, err := service.CreateUserService(ctx, user)

		assert.Empty(t, user)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error to creating user on database")
	})
}

func Test_FindUserService(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repository := mocks.NewMockRepositoryInterface(ctl)
	service := NewService(repository)

	t.Run("failed", func(t *testing.T) {
		id := 4
		ctx := context.Background()
		user := models.User{}
		repository.EXPECT().FindUserRepository(ctx, id).Return(user, &errors.ErrorResponse{
			Message: "user not found",
		})
		u, err := service.FindUserService(ctx, id)

		errorResponse := err.(*errors.ErrorResponse)

		assert.Empty(t, u)
		assert.NotNil(t, err)
		assert.EqualValues(t, errorResponse.Message, "user not found")
	})

	t.Run("success", func(t *testing.T) {
		id := 4
		user := models.User{
			Name:     "test",
			Email:    "email@eail.com",
			Password: "22456789954",
		}
		ctx := context.Background()

		repository.EXPECT().FindUserRepository(ctx, id).Return(user, nil)
		u, err := service.FindUserService(ctx, id)

		assert.Nil(t, err)
		assert.EqualValues(t, user.Name, u.Name)
		assert.EqualValues(t, user.Email, u.Email)
		assert.EqualValues(t, user.Password, u.Password)
	})
}

func Test_UpdateUserService(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repository := mocks.NewMockRepositoryInterface(ctl)
	service := NewService(repository)

	t.Run("success", func(t *testing.T) {
		user := models.User{
			Name: "jao",
		}
		id := 1
		ctx := context.Background()
		repository.EXPECT().FindUserRepository(ctx, id).Return(models.User{}, nil)
		repository.EXPECT().UpdateUserRepository(ctx, id, gomock.Any()).Return(nil)
		err := service.UpdateUserService(ctx, id, &user)

		assert.Nil(t, err)
	})

	t.Run("failed", func(t *testing.T) {
		user := models.User{
			Name: "jao",
		}
		id := 1
		ctx := context.Background()
		repository.EXPECT().FindUserRepository(ctx, id).Return(models.User{}, nil)
		repository.EXPECT().UpdateUserRepository(ctx, id, gomock.Any()).Return(&errors.ErrorResponse{
			Message: "error to update user",
		})
		err := service.UpdateUserService(ctx, id, &user)

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error to update user")
	})
}

func Test_DeleteUserService(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repository := mocks.NewMockRepositoryInterface(ctl)
	service := NewService(repository)

	t.Run("success", func(t *testing.T) {
		id := 1
		ctx := context.Background()

		repository.EXPECT().FindUserRepository(ctx, id).Return(models.User{}, nil)
		repository.EXPECT().DeleteUserRepository(ctx, id).Return(nil)
		
		err := service.DeleteUserService(ctx, id)
		assert.Nil(t, err)
	})

	t.Run("failed", func(t *testing.T) {
		id := 1
		ctx := context.Background()

		repository.EXPECT().FindUserRepository(ctx, id).Return(models.User{}, nil)
		repository.EXPECT().DeleteUserRepository(ctx, id).Return(&errors.ErrorResponse{
			Message: "error to deleting user",
		})
		
		err := service.DeleteUserService(ctx, id)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error to deleting user")
	})
}

func Test_ListUsersService(t *testing.T){
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repository := mocks.NewMockRepositoryInterface(ctl)
	service := NewService(repository)

	t.Run("success", func(t *testing.T){
		ctx := context.Background()
		user := []models.User{
			{ID: 1,Name: "gabriel", Email: "gabriel123@gmail.com"},
			{ID: 2,Name: "maria", Email: "maria123@gmail.com"},
			{ID: 3,Name: "carlos", Email: "carlos123@gmail.com"},
		}

		repository.EXPECT().ListUserRepository(ctx).Return(user, nil)
		users, err := service.ListUsers(ctx)

		assert.Nil(t, err)
		for i := 0; i < len(users); i++ {
			assert.Equal(t, users[i].Name, user[i].Name)
			assert.Equal(t, users[i].ID, user[i].ID)
			assert.Equal(t, users[i].Email, user[i].Email)
		}
	})

	t.Run("failed", func(t *testing.T){
		ctx := context.Background()

		repository.EXPECT().ListUserRepository(ctx).Return(nil, &errors.ErrorResponse{
			Message: "error to listing users",
		})
		users, err := service.ListUsers(ctx)

		assert.Nil(t, users)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error to listing users")
		
	})
}

func Test_LoginService(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repository := mocks.NewMockRepositoryInterface(ctl)
	service := NewService(repository)

	t.Run("error find user", func(t *testing.T){
		u := &models.User{
			Email: "cristiano@gmail.com",
			Password: "12345678",
		}
		ctx := context.Background()

		repository.EXPECT().FindUserByEmailRepository(ctx, u.Email).Return(nil, &errors.ErrorResponse{
			Message: "email not found",
		})
		user, token, err := service.Login(ctx, u)
	
		assert.Nil(t, user)
		assert.Empty(t, token)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "email not found")
	})

	t.Run("success", func(t *testing.T){
		u := &models.User{
			Email: "cristiano@gmail.com",
			Password: "12345678",
		}
		uHash := &models.User{
			Email: "cristiano@gmail.com",
			Password: "12345678",
		}
		_ = uHash.HashPassword(uHash.Password)
		ctx := context.Background()

		repository.EXPECT().FindUserByEmailRepository(ctx, u.Email).Return(uHash, nil)
		user, token, err := service.Login(ctx, u)
	
		assert.Nil(t, err)	
		assert.EqualValues(t, user.Name, u.Name)
		assert.NotEmpty(t, token)
	})
}