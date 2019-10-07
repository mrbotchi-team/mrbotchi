package errors

import "net/http"

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
