package middlewares

import (
	"github.com/gin-gonic/gin"

	"github.com/senonerk/sup/shared/aerr"
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
				parsedError = aerr.FromErr(err)
			}

			c.JSON(parsedError.Code, gin.H{"message": parsedError.Message})
			c.Abort()
		}

	}
}
