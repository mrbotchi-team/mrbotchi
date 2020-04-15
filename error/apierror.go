package error

// APIError はエラーを表現する構造体。
type APIError struct {
	error   `json:"-"`
	ID      string `json:"id"`
	Message string `json:"message"`
}
