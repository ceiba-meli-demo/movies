package app

import (
	"github.com/ceiba-meli-demo/movies/infrastructure/controllers"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping)


}

