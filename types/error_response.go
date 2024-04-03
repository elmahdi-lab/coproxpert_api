package types

// ErrorResponse TODO: this is not good, we should use a better error handling
type ErrorResponse struct {
	// Message is the error message
	Code    string
	Message string
}
