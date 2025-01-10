package exception

import (
	"fmt"
	"github.com/hipeday/rosen/internal/messages"
	"net/http"
)

type NotFoundError struct {
	Values  []string
	Message string
}

func getFormatMessage(values ...string) []any {
	if values == nil {
		return nil
	}
	var format []any
	for _, value := range values {
		format = append(format, value)
	}
	return format
}

func (n *NotFoundError) Error() string {
	return fmt.Sprintf(n.Message+": %v", getFormatMessage(n.Values...))
}

func (n *NotFoundError) Status() int {
	return http.StatusNotFound
}

func NewNotFoundError(values ...string) NotFoundError {
	return NotFoundError{
		Values:  values,
		Message: messages.DataDoesNotExist.String(),
	}
}
