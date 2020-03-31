package movies

import (
	"context"
	"errors"
	"fmt"
	"github.com/ceiba-meli-demo/movies/domain/model"
	"github.com/ceiba-meli-demo/movies/infrastructure/adapters/repository/models"
	"github.com/ceiba-meli-demo/movies/infrastructure/database_client"
	"github.com/ceiba-meli-demo/movies/infrastructure/mappers/movie_mapper"
	"github.com/ceiba-meli-demo/movies/infrastructure/utils/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type MovieSqlRepository struct {
	Connection *mongo.Client
}

func (movieSqlRepository *MovieSqlRepository) GetAll() ([]model.Movie, error) {
	var moviesDb []models.MovieDb
	collection := movieSqlRepository.Connection.Database(database_client.Schema).Collection("movies")
	result, err := collection.Find(context.TODO(), moviesDb)
	if err != nil {
		log.Fatal(err)
	}
	for result.Next(context.TODO()) {
		var movie models.MovieDb
		err := result.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		moviesDb = append(moviesDb,movie)
	}
	if err := result.Err(); err != nil {
		log.Fatal(err)
	}
	movies := movie_mapper.MoviesDbToMovies(moviesDb)
	return movies, nil
}

func (movieSqlRepository *MovieSqlRepository) GetById(movieId int64) (model.Movie, error) {
	var movieDb models.MovieDb
	collection := movieSqlRepository.Connection.Database(database_client.Schema).Collection("movies")
	if err:= collection.FindOne(context.TODO(), movieId).Decode(&movieDb); err!= nil {
		log.Fatal(err)
	}
	movie := movie_mapper.MovieDbToMovie(movieDb)
	return movie, nil
}

func (movieSqlRepository *MovieSqlRepository) Save(movie *model.Movie) error{
	var movieDb models.MovieDb
	movieDb = movie_mapper.MovieToMovieDb(*movie)
	collection := movieSqlRepository.Connection.Database(database_client.Schema).Collection("movies")
	if _, err := collection.InsertOne(context.TODO(), movieDb); err !=nil{
		logger.Error(fmt.Sprintf("Can't work with %s", movieDb.Title), err)
		return errors.New(fmt.Sprint("Can't work with #{movieDb.Title}"))
	}
	return nil
}