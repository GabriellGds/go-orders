package models

import (
	"net/http"

	"github.com/GabriellGds/go-orders/pkg/errors"
)


func ItemErrorParam(field, message string) *errors.ErrorResponse {
	return &errors.ErrorResponse{
		Field:   field,
		Message: message,
		Code: http.StatusBadRequest,
	}
}

func (ir *ItemRequest) Validate() error {
	if ir.Name == "" {
		return ItemErrorParam("name", " name (type: string) is required")
	}

	if ir.Price < 5.00 {
		return ItemErrorParam("price", "price (type: float) has to be greater than 5")
	}

	return nil
}


