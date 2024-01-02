package models

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func UserErrorParam(field, message string) error {
	if field == "password" {
		return &ErrorResponse{
			Field:   field,
			Message: message,
		}
	}

	return &ErrorResponse{
		Field:   field,
		Message: message,
	}
}

func (ur *UserLogin) Validate() error {
	if ur.Email == "" {
		return UserErrorParam("email", "email (type string) is required")
	}
	if ur.Email == "" {
		return UserErrorParam("email", "email (type string) is required")
	}
	if len(ur.Password) < 8 {
		return UserErrorParam("password", "the password must have at least 8 characters")

	}
	return nil
}

func (ur *UserRequest) Validate() error {
	if ur.Name == "" && ur.Email == "" && ur.Password == "" {
		return fmt.Errorf("request body is empty")
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
		return UserErrorParam("password", "the password must have at least 8 characters")

	}
	return nil
}

func (ur *UserUpdateRequest) Validate() error {
	if ur.Name == "" {
		return UserErrorParam("name", "name (type string) is required")
	}
	return nil
}
