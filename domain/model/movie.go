package model

import (
	"github.com/ceiba-meli-demo/movies/domain/exceptions"
	"github.com/ceiba-meli-demo/movies/domain/validators"
)

type Movie struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Duration int64  `json:"duration"`
	Synopsis string `json:"synopsis"`
}

//CreateMovie return a valid Movie
func (movie *Movie) CreateMovie(title string, duration int64, synopsis string) (Movie, error) {
	if err := validators.ValidateRequired(title, "Title should have some value"); err != nil {
		return Movie{}, err
	}
	if err := validators.ValidateRequiredDuration(duration, "Duration should have some value > 0 "); err != nil {
		return Movie{}, exceptions.InvalidDuration{ErrMessage: err.Error()}
	}
	if err := validators.ValidateRequired(synopsis, "Synopsis should have some value"); err != nil {
		return Movie{}, err
	}
	return Movie{
		Title:    title,
		Duration: duration,
		Synopsis: synopsis,
	}, nil
}
