package model

// ErrorMsg : struct to hold error message to be returned on error
type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
