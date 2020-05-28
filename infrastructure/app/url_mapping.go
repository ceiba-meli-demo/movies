package app

import (
	"github.com/ceiba-meli-demo/movies/infrastructure/app/middlewares/users"
	"github.com/ceiba-meli-demo/movies/infrastructure/controllers"
	"github.com/gin-contrib/cors"
)

func mapUrls(handler controllers.RedirectMovieHandler) {

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	router.Use(cors.New(config))

	router.GET("/movies", handler.Get)
	router.GET("/movies/:movie_id", handler.FindByID)
	router.Use(users.UserRequired())
	{
		router.POST("/movies", handler.Create)
	}
}
