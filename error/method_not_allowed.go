package error

import "net/http"

type MethodNotAllowed struct {
	Error
}

func NewMethodNotAllowed() *MethodNotAllowed {
	return &MethodNotAllowed{
		Error{
			StatusCode: http.StatusMethodNotAllowed,
			ErrorID:    "METHOD_NOT_ALLOWED",
			Message:    "This endpoint doesn't allow this method.",
		},
	}
}
