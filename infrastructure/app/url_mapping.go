package app

import "github.com/ceiba-meli-demo/movies/infrastructure/controllers/movies"

func mapUrls(handler movies.RedirectUserHandler) {
	router.GET("/movies", handler.Get)
	router.GET("/movies/:movie_id", handler.FindById)
	router.POST("/movies", handler.Create)
}


