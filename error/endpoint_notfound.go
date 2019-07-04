package error

import "net/http"

type EndpointNotFoundError struct {
	Error
}

func NewEndpointNotFoundError() *EndpointNotFoundError {
	return &EndpointNotFoundError{
		Error{
			StatusCode: http.StatusNotFound,
			ErrorID:    "ENDPOINT_NOTFOUND",
			Message:    "That endpoint not found.",
		},
	}
}
