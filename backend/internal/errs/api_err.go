package errs

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type APIError struct {
	StatusCode int `json:"statusCode"`
	Message    any `json:"msg"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("api error: %d %v", e.StatusCode, e.Message)
}

func NewAPIError(statusCode int, err error) APIError {
	return APIError{
		StatusCode: statusCode,
		Message:    err.Error(),
	}
}

func BadRequest(err error) APIError {
	return NewAPIError(http.StatusBadRequest, err)
}

func InvalidJSON() APIError {
	return NewAPIError(http.StatusBadRequest, errors.New("invalid JSON request data"))
}

func NotFound(title string, withKey string, withValue any) APIError {
	return NewAPIError(http.StatusNotFound, fmt.Errorf("%s with %s='%s' not found", title, withKey, withValue))
}

func Conflict(title string, withKey string, withValue any) APIError {
	return NewAPIError(http.StatusConflict, fmt.Errorf("conflict: %s with %s='%s' already exists", title, withKey, withValue))
}

func InvalidRequestData(errors map[string]string) APIError {
	return APIError{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    errors,
	}
}

func InternalServerError() APIError {
	return NewAPIError(http.StatusInternalServerError, errors.New("internal server error"))
}

func GithubIntegrationError(err error) APIError {
	return NewAPIError(http.StatusInternalServerError, fmt.Errorf("error with github integration: %s", err.Error()))
}

func MissingAPIParamError(field string) APIError {
	return NewAPIError(http.StatusBadRequest, fmt.Errorf("missing request field: %s", field))
}

func AuthenticationError() APIError {
	return NewAPIError(http.StatusForbidden, fmt.Errorf("please authenticate properly"))
}

/* Post Requests Only */
func InvalidRequestBody(expected interface{}) APIError {
	fieldAcc := make([]string, 0, 10)

	// Use reflection to inspect the struct type
	t := reflect.TypeOf(expected)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" {
			fieldAcc = append(fieldAcc, jsonTag)
		}
	}

	msg := fmt.Sprintf("Expected Fields: %s", strings.Join(fieldAcc, ", "))
	return NewAPIError(http.StatusBadRequest, errors.New(msg))
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	var apiErr APIError
	if castedErr, ok := err.(APIError); ok {
		apiErr = castedErr
	} else {
		apiErr = InternalServerError()
	}

	slog.Error("HTTP API error", "err", err.Error(), "method", c.Method(), "path", c.Path())

	return c.Status(apiErr.StatusCode).JSON(apiErr)
}
