package helper

import "strings"

// Response for static shape json return
type Response struct {
	Message string `json:"message"`
	Error   any    `json:"errors"`
	Data    any    `json:"data"`
}

// Empty object used for data that doesn't want to be null on json
type EmptyObj struct{}

// BuildResponse to inject data value to dynamic success resopnse
func BuildResponse(message string, data any) Response {
	res := Response{
		Message: message,
		Error:   nil,
		Data:    data,
	}

	return res
}

// BuildErrorResopnse to inject data value to dynamic fail response
func BuildErrorResopnse(message, err string, data any) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Message: message,
		Error:   splittedError,
		Data:    data,
	}

	return res
}
