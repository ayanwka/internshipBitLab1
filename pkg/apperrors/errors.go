package apperrors

import "fmt"

type ValidationError struct {
	Message string
}

func (e ValidationError) Error() string {
	return e.Message
}

func NewValidationError(message string) ValidationError {
	return ValidationError{Message: message}
}

type NotFoundError struct {
	ID       any
	Resource string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("%s with err ID %v not found", e.Resource, e.ID)
}

func NewNotFoundError(resource string, id any) NotFoundError {
	return NotFoundError{Resource: resource, ID: id}
}

type ConflictError struct {
	Reason string
}

func (e ConflictError) Error() string {
	return e.Reason
}

func NewConflictError(reason string) ConflictError {
	return ConflictError{Reason: reason}
}
