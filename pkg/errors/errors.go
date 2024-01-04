package errors

type ErrorResponse struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
	Code    int    `json:"-"`
}

func (e *ErrorResponse) Error() string {
	return e.Field
}
