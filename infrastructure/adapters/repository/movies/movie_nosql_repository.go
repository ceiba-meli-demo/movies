package movies

import (
	"context"
	"errors"
	"fmt"

	"github.com/ceiba-meli-demo/movies/domain/model"
	"github.com/ceiba-meli-demo/movies/infrastructure/adapters/repository/models"
	"github.com/ceiba-meli-demo/movies/infrastructure/mappers/movie_mapper"
	"github.com/ceiba-meli-demo/movies/infrastructure/utils/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	Schema = "movie_db"
	Table  = "movie"
)

type MovieNoSqlRepository struct {
	Connection *mongo.Client
}

func (MovieNoSqlRepository *MovieNoSqlRepository) GetAll() ([]model.Movie, error) {
	var moviesDb []models.MovieDb
	collection := MovieNoSqlRepository.Connection.Database(Schema).Collection(Table)
	result, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		logger.Error("collection is empty or dont exists", err)
		return []model.Movie{}, errors.New("collection is empty or dont exists")
	}
	for result.Next(context.TODO()) {
		var movie models.MovieDb
		err := result.Decode(&movie)
		if err != nil {
			logger.Error("error when trying to decode movie record", err)
			return []model.Movie{}, errors.New("error when trying to decode movie record")
		}
		moviesDb = append(moviesDb, movie)
	}
	if err := result.Err(); err != nil {
		logger.Error("error in result generated", err)
		return []model.Movie{}, errors.New("error in result generated")
	}
	movies := movie_mapper.MoviesDbToMovies(moviesDb)
	return movies, nil
}

func (MovieNoSqlRepository *MovieNoSqlRepository) GetById(movieId string) (model.Movie, error) {
	var movieDb models.MovieDb
	IDMovie, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": IDMovie}
	collection := MovieNoSqlRepository.Connection.Database(Schema).Collection(Table)
	if err := collection.FindOne(context.TODO(), filter).Decode(&movieDb); err != nil {
		logger.Error(fmt.Sprintf("Can't find this id %s", movieId), err)
		return model.Movie{}, fmt.Errorf("Can't find this id %s", movieId)
	}
	movie := movie_mapper.MovieDbToMovie(movieDb)
	return movie, nil
}

func (MovieNoSqlRepository *MovieNoSqlRepository) Save(movie *model.Movie) error {
	var movieDb models.MovieDb
	movieDb = movie_mapper.MovieToMovieDb(*movie)
	collection := MovieNoSqlRepository.Connection.Database(Schema).Collection(Table)
	result, err := collection.InsertOne(context.TODO(), movieDb)
	if err != nil {
		logger.Error(fmt.Sprintf("Can't work with %s", movieDb.Title), err)
		return fmt.Errorf("Can't work with %s", movieDb.Title)
	}
	movie.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return nil
}
