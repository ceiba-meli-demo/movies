package model

import (
	"github.com/ceiba-meli-demo/movies/domain/exceptions"
	"github.com/ceiba-meli-demo/movies/domain/validators"
)

type Movie struct {
	ID       string `json:"Id"`
	Title    string `json:"Title"`
	Duration int64  `json:"Duration"`
	UrlImg   string `json:"UrlImg"`
	Synopsis string `json:"Synopsis"`
}

type Movies []Movie

//CreateMovie return a valid Movie
func (movie *Movie) CreateMovie(title string, duration int64, urlImg string, synopsis string) (Movie, error) {
	if err := validators.ValidateRequired(title, "Title should have some value"); err != nil {
		return Movie{}, err
	}
	if err := validators.ValidateRequiredDuration(duration, "Duration should have some value > 0 "); err != nil {
		return Movie{}, exceptions.InvalidDuration{ErrMessage: err.Error()}
	}
	if err := validators.ValidateRequired(urlImg, "url image should have some value"); err != nil {
		return Movie{}, err
	}
	if err := validators.ValidateRequired(synopsis, "Synopsis should have some value"); err != nil {
		return Movie{}, err
	}
	return Movie{
		Title:    title,
		Duration: duration,
		UrlImg:   urlImg,
		Synopsis: synopsis,
	}, nil
}
