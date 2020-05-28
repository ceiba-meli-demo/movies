package validators

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateRequired(t *testing.T) {
	var field string = ""
	var message string = ""
	err := ValidateRequired(field, message)
	assert.NotNil(t, err)
	assert.Equal(t, err, errors.New(message))
}

func TestValidateRequiredDuration(t *testing.T) {
	var field int64 = 4
	var message string = ""
	err := ValidateRequiredDuration(field, message)
	assert.Nil(t, err)
}
