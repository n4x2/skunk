package handler

// EmptyValueError represents an error when a field is empty.
type EmptyValueError struct {
	Field string
}

// Error returns the error message for EmptyValueError.
func (e *EmptyValueError) Error() string {
	return "error: " + e.Field + " cannot be empty"
}
