package controller

// ErrorResponse is the response if something goes bad
type ErrorResponse struct {
	StatusCode int
	Message    string
}

// Response is the response if everything worked fine
type Response struct {
	Name     string
	HomePage string
	IsAdmin  bool
}
