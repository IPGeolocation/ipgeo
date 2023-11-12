package error

import (
    "fmt"
)

type CustomError struct {
    StatusCode int    `json:"statusCode"`
    Message    string `json:"message"`
}

func (e *CustomError) Error() string {
    return fmt.Sprintf("Status: %d, Message: %s", e.StatusCode, e.Message)
}

func HandleErrorResponse(statusCode int, apiMessage string) *CustomError {
    return &CustomError{
        StatusCode: statusCode,
        Message:    apiMessage,
    }
}

