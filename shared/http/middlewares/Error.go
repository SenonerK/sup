package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/senonerk/sup/shared/aerr"
	"github.com/senonerk/sup/shared/mode"
)

// ErrorReporter catches http errors and composes a json response
func ErrorReporter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		detectedErrors := c.Errors.ByType(gin.ErrorTypeAny)

		if len(detectedErrors) > 0 {
			err := detectedErrors[0].Err
			var parsedError *aerr.AppError
			switch err.(type) {
			case *aerr.AppError:
				parsedError = err.(*aerr.AppError)
			default:
				msg := aerr.FromErr(err).Error()
				if mode.Prod() {
					msg = "Internal Server Error"
				}
				parsedError = &aerr.AppError{
					Code:    http.StatusInternalServerError,
					Message: msg,
				}
			}

			c.JSON(parsedError.Code, gin.H{"message": parsedError.Message})
			c.Abort()
		}

	}
}
