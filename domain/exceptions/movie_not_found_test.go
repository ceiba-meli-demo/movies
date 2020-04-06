package exceptions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovieNotFoundError(t *testing.T) {
	const message = "not found"
	movieNotFound := MovieNotFound{ErrMessage: message}

	assert.NotNil(t, movieNotFound)
	assert.Equal(t, message, movieNotFound.Error())
}
