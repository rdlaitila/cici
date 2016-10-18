package api

// ErrorResult is a generic error type returned
// when an error occurs in the api
type ErrorResult struct {
	StatusCode int
	Message    string
}
