package api

// Result represents a generic api response object
type Result struct {
	StatusCode int
	Data       interface{}
}
