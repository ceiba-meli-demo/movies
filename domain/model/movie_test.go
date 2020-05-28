package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var movie Movie

func TestCreateMovie(t *testing.T) {
	nMovie, err := movie.CreateMovie("", 0, "", "")
	assert.EqualValues(t, Movie{}, nMovie, "brand new movie is empty")
	assert.True(t, err != nil, "err is for title of movie is empty")
	nMovie, err = movie.CreateMovie("Code", 0, "", "")
	assert.EqualValues(t, Movie{}, nMovie, "brand new movie is empty")
	assert.True(t, err != nil, "err is for duration of movie is 0")
	nMovie, err = movie.CreateMovie("Code", -10, "", "")
	assert.EqualValues(t, Movie{}, nMovie, "brand new movie is empty")
	assert.True(t, err != nil, "err is for duration of movie is minor than 0")
	nMovie, err = movie.CreateMovie("Code", 100, "", "")
	assert.EqualValues(t, Movie{}, nMovie, "brand new movie is empty")
	assert.True(t, err != nil, "err is for synopsis of movie is empty")
	nMovie, err = movie.CreateMovie("Code", 100, "", "Is a movie about coders")
	assert.EqualValues(t, Movie{Title: "Code", Duration: 100, Synopsis: "Is a movie about coders"}, nMovie, "brand new movie is correctly created")
	assert.True(t, err == nil, "err is nil because movie is correctly created")
}
