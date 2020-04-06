package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var movie Movie

func TestCreateEmptyMovie(t *testing.T) {
	nMovie, err := movie.CreateMovie("", 0, "")
	assert.EqualValues(t, Movie{}, nMovie, "brand new movie is empty")
	assert.NotNil(t, err, "err is for title of movie is empty")
}

func TestCreateWithTitleMovie(t *testing.T) {
	nMovie, err := movie.CreateMovie("Code", 0, "")
	assert.EqualValues(t, Movie{}, nMovie, "brand new movie is empty")
	assert.NotNil(t, err, "err is for duration of movie is 0")
}

func TestCreateWithTitleAndNegativeDurationMovie(t *testing.T) {
	nMovie, err := movie.CreateMovie("Code", -10, "")
	assert.EqualValues(t, Movie{}, nMovie, "brand new movie is empty")
	assert.NotNil(t, err, "err is for duration of movie is minor than 0")
}

func TestCreateWithTitleAndDurationMovie(t *testing.T) {
	nMovie, err := movie.CreateMovie("Code", 100, "")
	assert.EqualValues(t, Movie{}, nMovie, "brand new movie is empty")
	assert.NotNil(t, err, "err is for synopsis of movie is empty")
}

func TestCreateWithAllMovie(t *testing.T) {
	nMovie, err := movie.CreateMovie("Code", 100, "Is a movie about coders")
	assert.EqualValues(t, Movie{Title: "Code", Duration: 100, Synopsis: "Is a movie about coders"}, nMovie, "brand new movie is correctly created")
	assert.Nil(t, err, "err is nil because movie is correctly created")
}
