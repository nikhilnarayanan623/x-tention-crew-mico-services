package response

import "strings"

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   any    `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}

// Make success response
func SuccessResponse(message string, data any) Response {
	return Response{
		Success: true,
		Message: message,
		Data:    data,
	}
}

// Make error response
func ErrorResponse(message string, err error) Response {
	// separate error string by newline
	errArr := strings.Split(err.Error(), "\n")
	return Response{
		Success: false,
		Message: message,
		Error:   errArr,
	}
}
