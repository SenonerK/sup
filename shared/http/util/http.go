package util

import (
	"github.com/gin-gonic/gin"
	"github.com/senonerk/sup/shared/aerr"
)

// ErrInavlidForm if error
func ErrInavlidForm() *aerr.AppError {
	return &aerr.AppError{
		Code:    406,
		Message: "Invalid form",
	}
}

// ErrBadRequest br error
func ErrBadRequest() *aerr.AppError {
	return &aerr.AppError{
		Code:    400,
		Message: "Bad Request",
	}
}

// ErrUnauthorized u error
func ErrUnauthorized() *aerr.AppError {
	return &aerr.AppError{
		Code:    401,
		Message: "Unauthorized",
	}
}

// Ok sends a 200 ok message
func Ok(c *gin.Context, d interface{}) {
	c.JSON(200, gin.H{"message": "Success", "data": d})
}
