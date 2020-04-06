package usescases

import (
	"github.com/ceiba-meli-demo/movies/domain/model"
	"github.com/ceiba-meli-demo/movies/domain/ports"
)

type GetMovieByIDUseCase interface {
	Handler(movieID string) (model.Movie, error)
}

type UseCaseGetMovieById struct {
	MovieRepository ports.MovieRepository
}

func (useCaseGetMovieById *UseCaseGetMovieById) Handler(movieID string) (model.Movie, error) {
	movie, err := useCaseGetMovieById.MovieRepository.GetByID(movieID)
	if err != nil {
		return model.Movie{}, err
	}
	return movie, nil
}
