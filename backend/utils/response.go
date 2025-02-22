package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ValidationErrors defines a standard structure for validation errors
type ValidationErrors struct {
	Payload string `json:"payload"`
	Message string `json:"message"`
}

// BaseResponse defines a standard structure for API responses
type BaseResponse struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Data    interface{}        `json:"data,omitempty"`
	Errors  []ValidationErrors `json:"errors,omitempty"`
}

// SuccessResponse sends a successful response
func SuccessResponse(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusOK, BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse sends an error response
func ErrorResponse(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, BaseResponse{
		Success: false,
		Message: message,
	})
}

// ValidationErrorResponse sends a validation error response
func ValidationErrorResponse(c echo.Context, message string, errors []ValidationErrors) error {
	return c.JSON(http.StatusBadRequest, BaseResponse{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}
