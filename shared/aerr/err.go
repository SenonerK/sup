package aerr

// AppError is a general type for error handling
type AppError struct {
	Code    int
	Message string
	Source  string
}

func (e *AppError) Error() string {
	return e.Message
}

func FromErr(err error) error {
	return &AppError{
		Code:    500,
		Message: err.Error(),
	}
}
