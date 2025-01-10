package exception

import (
	"github.com/hipeday/rosen/internal/messages"
	"net/http"
)

type ValidationError struct {
	Message messages.ErrorMessage
}

func (v *ValidationError) Error() string {
	return v.Message.String()
}

func (v *ValidationError) Status() int {
	return http.StatusBadRequest
}

func NewValidationError(message messages.ErrorMessage) ValidationError {
	return ValidationError{
		Message: message,
	}
}
