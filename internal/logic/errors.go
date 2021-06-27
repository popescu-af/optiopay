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

// AlreadyFoundError should be returned when a specific resource
// already exists.
type AlreadyFoundError struct {
	message string
}

func (e *AlreadyFoundError) Error() string {
	return e.message
}

// NewAlreadyFoundError creates a custom AlreadyFoundError instance.
func NewAlreadyFoundError(message string) *AlreadyFoundError {
	return &AlreadyFoundError{message: message}
}

// ArgumentError should be returned when one or more arguments for an operation are invalid.
type ArgumentError struct {
	message string
}

func (e *ArgumentError) Error() string {
	return e.message
}

// ArgumentError creates a custom ArgumentError instance.
func NewArgumentError(message string) *ArgumentError {
	return &ArgumentError{message: message}
}
