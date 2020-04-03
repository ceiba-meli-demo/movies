package controllers

import (
	"encoding/hex"
	"net/http"

	"github.com/ceiba-meli-demo/movies/application/commands"
	"github.com/ceiba-meli-demo/movies/application/usescases"
	"github.com/ceiba-meli-demo/movies/infrastructure/app/middlewares/users"
	"github.com/ceiba-meli-demo/movies/infrastructure/utils/rest_errors"
	"github.com/gin-gonic/gin"
)

type RedirectMovieHandler interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	FindByID(c *gin.Context)
}

type Handler struct {
	GetMoviesUseCase    usescases.GetMoviesUseCase
	GetMovieByIDUseCase usescases.GetMovieByIDUseCase
	CreateMovieUseCase  usescases.CreateMoviePort
}

//Get All method, Find movies
func (handler *Handler) Get(c *gin.Context) {
	movies, err := handler.GetMoviesUseCase.Handler()
	if err != nil {
		restErr := rest_errors.NewBadRequestError(err.Error())
		c.JSON(restErr.Status(), restErr)
		return
	}
	c.JSON(http.StatusOK, movies)
}

//FindByID method, Find movies by id
func (handler *Handler) FindByID(c *gin.Context) {
	movieID := c.Param("movie_id")
	_, idErr := hex.DecodeString(movieID)
	if idErr != nil {
		restErr := rest_errors.NewBadRequestError("movie_id should be valid")
		c.JSON(restErr.Status(), restErr)
		return
	}
	movie, errGet := handler.GetMovieByIDUseCase.Handler(movieID)
	if errGet != nil {
		_ = c.Error(errGet)
		return
	}
	c.JSON(http.StatusOK, movie)
}

//Create method controller, get json type movie
func (handler *Handler) Create(c *gin.Context) {
	var movieCommand commands.MovieCommand
	if err := c.ShouldBindJSON(&movieCommand); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json")
		c.JSON(restErr.Status(), restErr)
		return
	}
	result, createMovieErr := handler.CreateMovieUseCase.Handler(movieCommand)
	if createMovieErr != nil {
		_ = c.Error(createMovieErr)
		return
	}
	if _, err := users.UpdateUser(movieCommand.Dni); err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, result)
}
