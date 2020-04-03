package usescases

import (
	"github.com/ceiba-meli-demo/movies/domain/model"
	"github.com/ceiba-meli-demo/movies/domain/ports"
)

type GetMoviesUseCase interface {
	Handler() (model.Movies, error)
}

type UseCaseGetMovies struct {
	MovieRepository ports.MovieRepository
}

func (useCaseGetMovies *UseCaseGetMovies) Handler() (model.Movies, error) {
	movies, err := useCaseGetMovies.MovieRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return movies, nil
}
