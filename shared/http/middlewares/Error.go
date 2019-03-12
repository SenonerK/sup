package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/senonerk/sup/shared/http/util"
)

func ErrorReporter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		detectedErrors := c.Errors.ByType(gin.ErrorTypeAny)

		if len(detectedErrors) > 0 {
			err := detectedErrors[0].Err
			var parsedError *util.AppError
			switch err.(type) {
			case *util.AppError:
				parsedError = err.(*util.AppError)
			default:
				parsedError = &util.AppError{
					Code:    http.StatusInternalServerError,
					Message: "Internal Server Error",
				}
			}

			c.JSON(parsedError.Code, gin.H{"message": parsedError.Message})
			c.Abort()
		}

	}
}
