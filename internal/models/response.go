package models

type ErrorResponse struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

func (e *ErrorResponse) Error() string {
	return e.Field
}