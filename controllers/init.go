package controllers

// BadRequestError custom HTTP 400 Bad Request Error
type BadRequestError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewBadRequestError return custom 400 error
func NewBadRequestError(code int, message string) *BadRequestError {
	return &BadRequestError{code, message}
}
