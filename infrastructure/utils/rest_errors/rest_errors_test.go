package rest_errors

import (
	"errors"
	"gotest.tools/assert"
	"strings"
	"testing"
)

func TestNewBadRequestError(t *testing.T) {
	const message = "bad request"
	error := NewBadRequestError(message)
	assert.Equal(t, message, error.Message())
	assert.Equal(t, strings.Contains(error.Error(), "bad_request"), true)
}

func TestNewInternalServerError(t *testing.T) {
	const message = "internal server error"
	error := NewInternalServerError(message, errors.New("internal server error"))
	assert.Equal(t, message, error.Message())
	assert.Equal(t, strings.Contains(error.Error(), "internal_server_error"), true)
}

func TestNewNotFoundError(t *testing.T) {
	const message = "not found error"
	error := NewNotFoundError(message)
	assert.Equal(t, message, error.Message())
	assert.Equal(t, strings.Contains(error.Error(), "not_found"), true)
}

func TestNewRestError(t *testing.T) {
	restError := NewRestError("error", 200, "error", nil)
	_, ok := restError.(restErr)
	assert.Equal(t, true, ok)
}
