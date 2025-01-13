package exception

import (
	"github.com/hipeday/rosen/internal/messages"
	"net/http"
)

type ForbiddenError struct {
	Message messages.ErrorMessage
}

func (u *ForbiddenError) Error() string {
	return u.Message.String()
}

func (u *ForbiddenError) Status() int {
	return http.StatusForbidden
}

func NewForbiddenError(message messages.ErrorMessage) ForbiddenError {
	return ForbiddenError{
		Message: message,
	}
}
