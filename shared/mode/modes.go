package mode

import (
	"os"
	"strings"
)

var mode string

func init() {
	mode = strings.ToUpper(os.Getenv("SUP_MODE"))
}

// Dev is app in developement
func Dev() bool {
	return mode == "DEV" || mode == "DEVELOPEMENT" || mode == ""
}

// Prod is app in production
func Prod() bool {
	return mode == "PROD" || mode == "PRODUCTION"
}

// Test is app in testing
func Test() bool {
	return mode == "TEST" || mode == "TESTING"
}
