package errors

import (
	"fmt"
	"time"
)

// APIError represents an error response from the Directus API
type APIError struct {
	StatusCode int                    `json:"status_code"`
	Message    string                 `json:"message"`
	Errors     []FieldError           `json:"errors,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
}

// FieldError represents a validation error for a specific field
type FieldError struct {
	Message string `json:"message"`
	Field   string `json:"field"`
	Code    string `json:"code"`
}

// Error implements the error interface
func (e *APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("API error %d: %s", e.StatusCode, e.Message)
	}
	return fmt.Sprintf("API error %d", e.StatusCode)
}

// ValidationError represents a validation error
type ValidationError struct {
	Message string
	Field   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s", e.Message)
}

// AuthenticationError represents an authentication error
type AuthenticationError struct {
	Message string
}

func (e *AuthenticationError) Error() string {
	return fmt.Sprintf("authentication error: %s", e.Message)
}

// AuthorizationError represents an authorization error
type AuthorizationError struct {
	Message string
}

func (e *AuthorizationError) Error() string {
	return fmt.Sprintf("authorization error: %s", e.Message)
}

// NotFoundError represents a resource not found error
type NotFoundError struct {
	Resource string
	ID       string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s with ID %s not found", e.Resource, e.ID)
}

// RateLimitError represents a rate limit error
type RateLimitError struct {
	RetryAfter time.Duration
}

func (e *RateLimitError) Error() string {
	return fmt.Sprintf("rate limit exceeded, retry after %v", e.RetryAfter)
}

// NewAPIError creates a new API error
func NewAPIError(statusCode int, message string) *APIError {
	return &APIError{
		StatusCode: statusCode,
		Message:    message,
	}
}

// NewValidationError creates a new validation error
func NewValidationError(message, field string) *ValidationError {
	return &ValidationError{
		Message: message,
		Field:   field,
	}
}

// NewAuthenticationError creates a new authentication error
func NewAuthenticationError(message string) *AuthenticationError {
	return &AuthenticationError{
		Message: message,
	}
}

// NewAuthorizationError creates a new authorization error
func NewAuthorizationError(message string) *AuthorizationError {
	return &AuthorizationError{
		Message: message,
	}
}

// NewNotFoundError creates a new not found error
func NewNotFoundError(resource, id string) *NotFoundError {
	return &NotFoundError{
		Resource: resource,
		ID:       id,
	}
}

// NewRateLimitError creates a new rate limit error
func NewRateLimitError(retryAfter time.Duration) *RateLimitError {
	return &RateLimitError{
		RetryAfter: retryAfter,
	}
}
