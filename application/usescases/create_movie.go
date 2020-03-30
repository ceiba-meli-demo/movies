package usescases

import (
	"github.com/ceiba-meli-demo/movies/application/commands"
	"github.com/ceiba-meli-demo/movies/application/factory"
	"github.com/ceiba-meli-demo/movies/domain/model"
	"github.com/ceiba-meli-demo/movies/domain/ports"
)

type CreateMoviePort interface {
	Handler(userCommand commands.MovieCommand) (model.Movie, error)
}

type UseCaseMovieCreate struct {
	MovieRepository ports.MovieRepository
}

func (createMovieUseCase *UseCaseMovieCreate) Handler(movieCommand commands.MovieCommand) (model.Movie, error) {
	movie, err := factory.CreateMovie(movieCommand)
	if err !=nil {
		return model.Movie{}, err
	}
	createUserErr := createMovieUseCase.MovieRepository.Save(&movie)
	if createUserErr != nil {
		return model.Movie{}, createUserErr
	}
	return movie, nil

}
