package exceptions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidDurationError(t *testing.T) {
	const message = "Invalid duration"
	invalidDuration := InvalidDuration{ErrMessage: message}

	assert.NotNil(t, invalidDuration)
	assert.Equal(t, message, invalidDuration.Error())

	assert.Equal(t, true, invalidDuration.IsBusinessLogic())
}
