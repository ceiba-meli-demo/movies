package usescases

import (
	"github.com/ceiba-meli-demo/movies/domain/model"
	"github.com/ceiba-meli-demo/movies/domain/ports"
)

type GetMovieUseCase interface {
	Handler() ([]model.Movie, error)
}

type UseCaseGetMovie struct {
	MovieRepository ports.MovieRepository
}

func (useCaseGetMovie *UseCaseGetMovie) Handler() ([]model.Movie, error) {
	movies, err := useCaseGetMovie.MovieRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return movies, nil
}

