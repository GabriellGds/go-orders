package models

import (
	
	"net/http"

	"github.com/GabriellGds/go-orders/pkg/errors"
	"github.com/badoux/checkmail"
)

func UserErrorParam(field, message string) *errors.ErrorResponse{
	return &errors.ErrorResponse{
		Field:   field,
		Message: message,
		Code: http.StatusBadRequest,
	}
}

func (ur *UserLogin) Validate() error {
	if ur.Email == "" {
		return UserErrorParam("email", "email (type: string) is required")
	}
	if ur.Email == "" {
		return UserErrorParam("email", "email (type: string) is required")
	}
	if len(ur.Password) < 8 {
		return UserErrorParam("password", "password(type: string) the password must have at least 8 characters")
	}
	return nil
}

func (ur *UserRequest) Validate() *errors.ErrorResponse {
	if ur.Name == "" && ur.Email == "" && ur.Password == "" {
		return &errors.ErrorResponse{
		Message: "request body is empty",
		Code: http.StatusBadRequest,
	}

	}
	if ur.Name == "" {
		return UserErrorParam("name", "name (type: string) is required")
	}
	if ur.Email == "" {
		return UserErrorParam("email", "email (type string) is required")
	}
	if err := checkmail.ValidateFormat(ur.Email); err != nil {
		return UserErrorParam("email", "The entered email is invalid.")
	}

	if len(ur.Password) < 8 {
		return UserErrorParam("password", "password (type: string) the password must have at least 8 characters")

	}
	return nil
}

func (ur *UserUpdateRequest) Validate() *errors.ErrorResponse {
	if ur.Name == "" {
		return UserErrorParam("name", "name (type string) is required")
	}
	return nil
}
