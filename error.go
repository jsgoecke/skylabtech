package skylabtech

type Error struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Errors  map[string][]string `json:"errors"`
}

type APIError struct {
	StatusCode int
	Err        Error
}

func (e *APIError) Error() string {
	return e.Err.Message
}
