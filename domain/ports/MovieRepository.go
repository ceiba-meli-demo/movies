package ports

import "github.com/ceiba-meli-demo/movies/domain/model"

type MovieRepository interface {
	Save(movie *model.Movie) error
	GetByID(movieID string) (model.Movie, error)
	GetAll() ([]model.Movie, error)
}
