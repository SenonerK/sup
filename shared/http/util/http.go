package util

import "github.com/gin-gonic/gin"

// AppError is a general type for error handling
type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

// ErrInavlidForm if error
func ErrInavlidForm() *AppError {
	return &AppError{
		Code:    406,
		Message: "Invalid form",
	}
}

// ErrBadRequest br error
func ErrBadRequest() *AppError {
	return &AppError{
		Code:    400,
		Message: "Bad Request",
	}
}

// ErrUnauthorized u error
func ErrUnauthorized() *AppError {
	return &AppError{
		Code:    401,
		Message: "Unauthorized",
	}
}

// Ok sends a 200 ok message
func Ok(c *gin.Context, d interface{}) {
	c.JSON(200, gin.H{"message": "Success", "data": d})
}
