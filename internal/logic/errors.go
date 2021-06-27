package logic

// NotFoundError should be returned when a specific resource
// requested by the client of the logic does not exist.
type NotFoundError struct {
	message string
}

func (e *NotFoundError) Error() string {
	return e.message
}

// NewNotFoundError creates a custom NotFoundError instance.
func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{message: message}
}
