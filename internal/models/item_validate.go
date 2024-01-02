package models


func ItemErrorParam(field, message string) error {
	if field == "price" {
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

func (ir *ItemRequest) Validate() error {
	if ir.Name == "" {
		return ItemErrorParam("name", " name (type: string) is a required")
	}

	if ir.Price < 5.00 {
		return ItemErrorParam("price", "price (type: float) has to be greater than 5")
	}

	return nil
}


