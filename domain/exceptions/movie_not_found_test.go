package exceptions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMovieNotFound_Error(t *testing.T) {
	const message = "not found"
	movieNotFound := MovieNotFound{ErrMessage: message}

	assert.NotNil(t, movieNotFound)
	assert.Equal(t, message, movieNotFound.Error())
}