package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	//logger.Info("about to start the application")
	router.Run()
	_ = router.Run(":8081")
}