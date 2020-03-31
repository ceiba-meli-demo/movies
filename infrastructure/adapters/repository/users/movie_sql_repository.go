package users

import (
	"github.com/ceiba-meli-demo/movies/domain/model"
	"github.com/ceiba-meli-demo/movies/infrastructure/adapters/repository/models"
	"github.com/ceiba-meli-demo/movies/infrastructure/mappers/movie_mapper"
)

type MovieSqlRepository struct {
	//Db *gorm.DB
}

func (movieSqlRepository *MovieSqlRepository) Get() ([]model.Movie, error) {
	var movieDb []models.MovieDb
	//Aquí irá el metodo que obtiene los registros en la base de datos, quitar if
	if movieDb==nil{
		return nil, nil
	}
	movies := movie_mapper.MoviesDbToMovies(movieDb)
	return movies, nil
}

func (movieSqlRepository *MovieSqlRepository) GetById(movieId int64) (model.User, error) {
	var movieDb models.MovieDb
	//Aquí irá el metodo que obtiene por id en la base de datos, quitar validación de if
	user := movie_mapper.MovieDbToMovie(movieDb)
	return user, nil
}

func (movieSqlRepository *MovieSqlRepository) Save(movie *model.Movie) error {
	var movieDb models.MovieDb
	movieDb = movie_mapper.MovieToMovieDb(*movie)
	//Aquí irá el metodo que guarda en la base de datos, quitar validación de if
	if movieDb.Title ==""{
		return nil
	}
	return nil
}