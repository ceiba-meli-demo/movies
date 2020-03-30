package app

import (
	"github.com/maik101010/movies/infrastructure/controllers"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}

