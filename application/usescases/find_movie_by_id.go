package usescases

import (
	"github.com/ceiba-meli-demo/movies/domain/model"
	"github.com/ceiba-meli-demo/movies/domain/ports"
)

type GetMovieByIdUseCase interface {
	Handler(movieId int64) (model.Movie, error)
}

type UseCaseGetMovieById struct {
	MovieRepository ports.MovieRepository
}

func (useCaseGetMovieById *UseCaseGetMovieById) Handler(userId int64) (model.Movie, error) {
	movie, err := useCaseGetMovieById.MovieRepository.GetById(userId)
	if err != nil {
		return model.Movie{}, err
	}
	return movie, nil
}
