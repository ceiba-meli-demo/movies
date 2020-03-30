package controllers

import (
	"net/http"
	"strconv"

	"github.com/ceiba-meli-demo/movies/application/commands"
	"github.com/ceiba-meli-demo/movies/application/usescases"
	"github.com/gin-gonic/gin"
)

type RedirectMovieHandler interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	FindById(c *gin.Context)
}

type Handler struct {
	GetMoviesUseCase    usescases.GetMovieUseCase
	GetMovieByIdUseCase usescases.GetMovieByIdUseCase
	CreateMovieUseCase  usescases.CreateMoviePort
}

//GetAll method, Find movies
func (handler *Handler) Get(c *gin.Context) {
	movies, err := handler.GetMoviesUseCase.Handler()
	if err != nil {
		//error
		c.String(501, err.Error())
		return
	}
	c.JSON(http.StatusOK, movies)
}

//GetById method, Find movies by id
func (handler *Handler) FindById(c *gin.Context) {
	movieId, idErr := strconv.ParseInt(c.Param("movie_id"), 10, 64)
	if idErr != nil {
		//error
		c.String(501, idErr.Error())
		return
	}
	movie, err := handler.GetMovieByIdUseCase.Handler(movieId)
	if err != nil {
		//error
		c.String(501, err.Error())
		return
	}
	c.JSON(http.StatusOK, movie)
}

//Create method controller, get json type movie
func (handler *Handler) Create(c *gin.Context) {
	var movieCommand commands.MovieCommand
	if err := c.ShouldBindJSON(&movieCommand); err != nil {
		//error
		c.String(501, err.Error())
		return
	}
	result, createMovieErr := handler.CreateMovieUseCase.Handler(movieCommand)
	if createMovieErr != nil {
		_ = c.Error(createMovieErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}
