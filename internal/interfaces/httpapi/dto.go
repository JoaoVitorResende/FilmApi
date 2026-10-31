package httpapi

type errorResponse struct {
	Error string `json:"error"`
}

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}
