package handler

// ErrorResponse represents the error what the handler send to the client in case of Error
type ErrorResponse struct {
	StatusCode int64
	Message    string
}

// SuccesResponse represent the response what the handler send to the clinet in case of Success
type SuccesResponse struct {
	StatusCode int64
	Token      string
}
