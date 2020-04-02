package factory

import (
	"github.com/ceiba-meli-demo/movies/application/commands"
	"github.com/ceiba-meli-demo/movies/domain/model"
)

func CreateMovie(movieCommand commands.MovieCommand) (model.Movie, error) {
	var movie model.Movie
	movie, err := movie.CreateMovie(movieCommand.Movie.Title, movieCommand.Movie.Duration, movieCommand.Movie.Synopsis)
	return movie, err
}
