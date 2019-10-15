package errors

import (
	"fmt"
	"net/http"
)

func EndpointNotFoundError() APIError {
	return APIError{
		ErrorID: "ENDPOINT_NOTFOUND",
		HTTPError: HTTPError{
			StatusCode: http.StatusNotFound,
			Message:    "That endpoint not found.",
		},
	}
}

func InternalServerError() APIError {
	return APIError{
		ErrorID: "INTERNAL_SERVER_ERROR",
		HTTPError: HTTPError{
			StatusCode: http.StatusInternalServerError,
			Message:    "An internal error has occurred. If the error persists, please report an issue on GitHub.",
		},
	}
}

func MethodNotAllowed() APIError {
	return APIError{
		ErrorID: "METHOD_NOT_ALLOWED",
		HTTPError: HTTPError{
			StatusCode: http.StatusMethodNotAllowed,
			Message:    "This endpoint doesn't allow this method.",
		},
	}
}

func ImSchwimmwagen() APIError {
	return APIError{
		ErrorID: "IM_SCHWIMMWAGEN",
		HTTPError: HTTPError{
			StatusCode: http.StatusTeapot,
			Message:    "I broke your teapot. So sorry...",
		},
	}
}

func InvalidRequest() APIError {
	return APIError{
		ErrorID: "INVALID_REQUEST",
		HTTPError: HTTPError{
			StatusCode: http.StatusBadRequest,
			Message:    "This request is invalid.",
		},
	}
}

func InvalidCT(contentType string) APIError {
	return APIError{
		ErrorID: "INVALID_CONTENTTYPE",
		HTTPError: HTTPError{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("Can't accept this request. Because Content-Type isn't %s.", contentType),
		},
	}
}

func AuthFailed() APIError {
	return APIError{
		ErrorID: "AUTH_FAILED",
		HTTPError: HTTPError{
			StatusCode: http.StatusForbidden,
			Message:    "username or password is incorrect.",
		},
	}
}
