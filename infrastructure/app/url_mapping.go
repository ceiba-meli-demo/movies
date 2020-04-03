package app

import (
	"github.com/ceiba-meli-demo/movies/infrastructure/app/middlewares/users"
	"github.com/ceiba-meli-demo/movies/infrastructure/controllers"
)

func mapUrls(handler controllers.RedirectMovieHandler) {
	router.GET("/movies", handler.Get)
	router.GET("/movies/:movie_id", handler.FindByID)
	router.Use(users.UserRequired())
	{
		router.POST("/movies", handler.Create)
	}
}
