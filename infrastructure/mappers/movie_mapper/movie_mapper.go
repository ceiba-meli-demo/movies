package movie_mapper

import (
	"github.com/ceiba-meli-demo/movies/domain/model"
	"github.com/ceiba-meli-demo/movies/infrastructure/adapters/repository/models"
)

func MovieToMovieDb(movie model.Movie) models.MovieDb {
	var movieDb models.MovieDb
	movieDb.Title = movie.Title
	movieDb.Duration = movie.Duration
	movieDb.UrlImg = movie.UrlImg
	movieDb.Synopsis = movie.Synopsis
	return movieDb
}

func MovieDbToMovie(movieDb models.MovieDb) model.Movie {
	var movie model.Movie
	movie.ID = movieDb.ID
	movie.Title = movieDb.Title
	movie.Duration = movieDb.Duration
	movie.UrlImg = movieDb.UrlImg
	movie.Synopsis = movieDb.Synopsis
	return movie
}

func MoviesDbToMovies(moviesDb []models.MovieDb) []model.Movie {
	var movies []model.Movie
	for _, movieDb := range moviesDb {
		movie := MovieDbToMovie(movieDb)
		movies = append(movies, movie)
	}
	return movies
}
