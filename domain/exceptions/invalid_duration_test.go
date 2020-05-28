package exceptions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvalidDurationError(t *testing.T) {
	const message = "Invalid duration"
	invalidDuration := InvalidDuration{ErrMessage: message}

	assert.NotNil(t, invalidDuration)
	assert.Equal(t, message, invalidDuration.Error())

	assert.Equal(t, true, invalidDuration.IsBusinessLogic())
}
