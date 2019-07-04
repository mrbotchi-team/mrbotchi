package error

import "net/http"

type InternalServerError struct {
	Error
}

func NewInternalServerError() *InternalServerError {
	return &InternalServerError{
		Error{
			StatusCode: http.StatusInternalServerError,
			ErrorID:    "INTERNAL_SERVER_ERROR",
			Message:    "An internal error has occurred. If the error persists, please report an issue on GitHub.",
		},
	}
}
