package errors

import "fmt"

const (
	NOT_FOUND_ERROR = iota
)

type NotFoundError struct {
	Message string
	Code    int
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{
		Code:    NOT_FOUND_ERROR,
		Message: message,
	}
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%d - %s", e.Code, e.Message)
}
