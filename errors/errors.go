package errors

import "net/http"

type (
	HTTPError struct {
		StatusCode int    `json:"status"`
		Message    string `json:"message"`
	}
	APIError struct {
		HTTPError
		ErrorID string `json:"errorID"`
	}
)

func (e HTTPError) Error() string {
	if "" == e.Message {
		return http.StatusText(e.StatusCode)
	}

	return e.Message
}
