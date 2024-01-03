package models

func OrderErrorParam(field, message string) error {
	return &ErrorResponse{
		Field:   field,
		Message: message,
	}
}

func(oi *OrderItems) Validate() error {
	if oi.ItemID <= 0 {
		return OrderErrorParam("item_id", "item_id (type: int)invalid id")
	}
	if oi.Quantity <= 0 {
		return OrderErrorParam("quantity", "quantity (type: int)The quantity must be greater than 0")
	}

	return nil
}

func (or *OrderRequest) Validate() error {
	if len(or.Items) == 0 {
        return OrderErrorParam("items", "items cannot be empty")
    }
	
	return nil
}