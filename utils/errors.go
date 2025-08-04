package utils

import (
	"fmt"
	"net/http"
	"strings"
)

// ErrorType represents different types of errors
type ErrorType string

const (
	ErrorTypeNetwork    ErrorType = "network"
	ErrorTypeValidation ErrorType = "validation"
	ErrorTypePermission ErrorType = "permission"
	ErrorTypeNotFound   ErrorType = "not_found"
	ErrorTypeRateLimit  ErrorType = "rate_limit"
	ErrorTypeInternal   ErrorType = "internal"
	ErrorTypeUser       ErrorType = "user"
)

// CLIError represents a user-friendly error
type CLIError struct {
	Type        ErrorType
	Message     string
	Details     string
	Suggestions []string
	Original    error
}

// NewCLIError creates a new CLI error
func NewCLIError(errType ErrorType, message, details string, suggestions []string, original error) *CLIError {
	return &CLIError{
		Type:        errType,
		Message:     message,
		Details:     details,
		Suggestions: suggestions,
		Original:    original,
	}
}

// Error implements the error interface
func (e *CLIError) Error() string {
	return e.Message
}

// Print prints the error in a user-friendly format
func (e *CLIError) Print() {
	PrintError(e.Message)

	if e.Details != "" {
		fmt.Printf("   Details: %s\n", e.Details)
	}

	if len(e.Suggestions) > 0 {
		fmt.Printf("   Suggestions:\n")
		for i, suggestion := range e.Suggestions {
			fmt.Printf("   %d. %s\n", i+1, suggestion)
		}
	}
}

// WrapNetworkError wraps network-related errors with user-friendly messages
func WrapNetworkError(err error, operation string) *CLIError {
	if err == nil {
		return nil
	}

	message := fmt.Sprintf("Failed to %s", operation)
	details := err.Error()

	suggestions := []string{
		"Check your internet connection",
		"Verify the URL is accessible",
		"Try again in a few moments",
	}

	return NewCLIError(ErrorTypeNetwork, message, details, suggestions, err)
}

// WrapHTTPError wraps HTTP errors with user-friendly messages
func WrapHTTPError(statusCode int, operation string) *CLIError {
	var message, details string
	var suggestions []string

	switch statusCode {
	case http.StatusNotFound:
		message = fmt.Sprintf("Resource not found during %s", operation)
		details = fmt.Sprintf("HTTP %d: The requested resource could not be found", statusCode)
		suggestions = []string{
			"Verify the search query is correct",
			"Check if the resource still exists",
			"Try a different search term",
		}
	case http.StatusTooManyRequests:
		message = fmt.Sprintf("Rate limit exceeded during %s", operation)
		details = fmt.Sprintf("HTTP %d: Too many requests, please wait before trying again", statusCode)
		suggestions = []string{
			"Wait a few minutes before trying again",
			"Reduce the number of concurrent requests",
			"Consider using a different search strategy",
		}
	case http.StatusForbidden:
		message = fmt.Sprintf("Access denied during %s", operation)
		details = fmt.Sprintf("HTTP %d: You don't have permission to access this resource", statusCode)
		suggestions = []string{
			"Check if you need authentication",
			"Verify your permissions",
			"Contact support if this persists",
		}
	case http.StatusInternalServerError:
		message = fmt.Sprintf("Server error during %s", operation)
		details = fmt.Sprintf("HTTP %d: The server encountered an internal error", statusCode)
		suggestions = []string{
			"Try again in a few moments",
			"Check if the service is experiencing issues",
			"Contact support if the problem persists",
		}
	default:
		message = fmt.Sprintf("Unexpected error during %s", operation)
		details = fmt.Sprintf("HTTP %d: An unexpected error occurred", statusCode)
		suggestions = []string{
			"Try again later",
			"Check the service status",
			"Contact support if the problem persists",
		}
	}

	return NewCLIError(ErrorTypeNetwork, message, details, suggestions, nil)
}

// WrapValidationError wraps validation errors
func WrapValidationError(field, value, reason string) *CLIError {
	message := fmt.Sprintf("Invalid %s: %s", field, value)
	details := reason

	suggestions := []string{
		fmt.Sprintf("Provide a valid %s", field),
		"Check the format requirements",
		"Refer to the help documentation",
	}

	return NewCLIError(ErrorTypeValidation, message, details, suggestions, nil)
}

// WrapPermissionError wraps permission-related errors
func WrapPermissionError(operation, resource string) *CLIError {
	message := fmt.Sprintf("Permission denied: cannot %s %s", operation, resource)
	details := "You don't have the required permissions for this operation"

	suggestions := []string{
		"Check your user permissions",
		"Contact your administrator",
		"Verify the resource exists and is accessible",
	}

	return NewCLIError(ErrorTypePermission, message, details, suggestions, nil)
}

// WrapNotFoundError wraps not found errors
func WrapNotFoundError(resourceType, identifier string) *CLIError {
	message := fmt.Sprintf("%s not found: %s", resourceType, identifier)
	details := fmt.Sprintf("The requested %s could not be found", resourceType)

	suggestions := []string{
		"Verify the identifier is correct",
		"Check if the resource still exists",
		"Try searching with different terms",
	}

	return NewCLIError(ErrorTypeNotFound, message, details, suggestions, nil)
}

// WrapRateLimitError wraps rate limit errors
func WrapRateLimitError(operation string, retryAfter string) *CLIError {
	message := fmt.Sprintf("Rate limit exceeded for %s", operation)
	details := fmt.Sprintf("Too many requests. Please wait %s before trying again", retryAfter)

	suggestions := []string{
		"Wait before making another request",
		"Reduce the frequency of requests",
		"Consider implementing request throttling",
	}

	return NewCLIError(ErrorTypeRateLimit, message, details, suggestions, nil)
}

// WrapInternalError wraps internal errors
func WrapInternalError(operation string, err error) *CLIError {
	message := fmt.Sprintf("Internal error during %s", operation)
	details := "An unexpected error occurred in the application"

	suggestions := []string{
		"Try again in a few moments",
		"Check if the issue persists",
		"Contact support with error details",
	}

	return NewCLIError(ErrorTypeInternal, message, details, suggestions, err)
}

// WrapUserError wraps user-caused errors
func WrapUserError(message, details string, suggestions []string) *CLIError {
	return NewCLIError(ErrorTypeUser, message, details, suggestions, nil)
}

// IsNetworkError checks if an error is network-related
func IsNetworkError(err error) bool {
	if cliErr, ok := err.(*CLIError); ok {
		return cliErr.Type == ErrorTypeNetwork
	}
	return strings.Contains(err.Error(), "network") ||
		strings.Contains(err.Error(), "connection") ||
		strings.Contains(err.Error(), "timeout")
}

// IsValidationError checks if an error is validation-related
func IsValidationError(err error) bool {
	if cliErr, ok := err.(*CLIError); ok {
		return cliErr.Type == ErrorTypeValidation
	}
	return strings.Contains(err.Error(), "validation") ||
		strings.Contains(err.Error(), "invalid")
}

// PrintErrorWithContext prints an error with additional context
func PrintErrorWithContext(err error, context string) {
	if cliErr, ok := err.(*CLIError); ok {
		fmt.Printf("Error in %s: ", context)
		cliErr.Print()
	} else {
		PrintError(fmt.Sprintf("Error in %s: %v", context, err))
	}
}

// HandleError handles an error with appropriate user feedback
func HandleError(err error, operation string) {
	if err == nil {
		return
	}

	// Try to wrap the error if it's not already a CLIError
	var cliErr *CLIError
	if existingCLIErr, ok := err.(*CLIError); ok {
		cliErr = existingCLIErr
	} else {
		// Wrap unknown errors as internal errors
		cliErr = WrapInternalError(operation, err)
	}

	cliErr.Print()
}
