package factory

import (
	"testing"

	"github.com/ceiba-meli-demo/movies/application/commands"
	"github.com/stretchr/testify/assert"
)

func TestGoodMovieFactory(t *testing.T) {
	var cmd commands.MovieCommand
	cmd.Movie.Title = "Eramos muchos y se preño la abuela"
	cmd.Movie.Duration = 120
	cmd.Movie.Synopsis = "Triste historia de la vida real"

	movie, err := CreateMovie(cmd)
	assert.NotNil(t, cmd)
	assert.NotNil(t, movie)
	assert.Nil(t, err, "err is equal to nil")
}

func TestBadMovieFactory(t *testing.T) {
	var cmd commands.MovieCommand
	cmd.Movie.Title = "Eramos muchos y se preño la abuela 2"
	cmd.Movie.Duration = 140

	movie, err := CreateMovie(cmd)
	assert.NotNil(t, cmd)
	assert.NotNil(t, movie)
	assert.NotNil(t, err, "err is present because synopsis is blank")
}
