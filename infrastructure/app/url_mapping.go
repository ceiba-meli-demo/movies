package app

import "github.com/ceiba-meli-demo/movies/infrastructure/controllers"

func mapUrls(handler controllers.RedirectMovieHandler) {
	router.GET("/movies", handler.Get)
	router.GET("/movies/:movie_id", handler.FindByID)
	router.POST("/movies", handler.Create)
}
