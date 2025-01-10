package exception

import (
	"net/http"
)

type UnauthorizedError struct {
}

func (u *UnauthorizedError) Error() string {
	return "Unauthorized"
}

func (u *UnauthorizedError) Status() int {
	return http.StatusUnauthorized
}

func NewUnauthorizedError() UnauthorizedError {
	return UnauthorizedError{}
}
