package app

import (
	"github.com/ceiba-meli-demo/movies/application/usescases"
	"github.com/ceiba-meli-demo/movies/domain/ports"
	"github.com/ceiba-meli-demo/movies/infrastructure/controllers/movies"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	router = gin.Default()
)

type UserMysqlRepository struct {
	Db *gorm.DB
}


func StartApplication() {
	//logger.Info("about to start the application")
	router.Run()
	_ = router.Run(":8081")
	movieRepository := getMovieRepository()
	var handler = createHandler(movieRepository)
	mapUrls(handler)

}

func createHandler(movieRepository ports.MovieRepository) movies.RedirectUserHandler {

	return newHandler(newCreateMovieUseCase(movieRepository), newGetMoviesUseCase(movieRepository),
		newFindMovieByIdUseCase(movieRepository))
}

func newHandler(createMovie usescases.CreateMoviePort, getMoviesUseCase usescases.GetMovieUseCase, getMovieById usescases.GetMovieByIdUseCase) controllers.RedirectUserHandler {
	return &movies.Handler{
		CreateMovieUseCase: createMovie,
		GetMoviesUseCase: getMoviesUseCase,
		GetMovieByIdUseCase: getMovieById,
	}
}

func newCreateMovieUseCase(repository ports.MovieRepository) usescases.CreateMoviePort {
	return &usescases.UseCaseMovieCreate{
		MovieRepository: repository,
	}
}

func newGetMoviesUseCase(repository ports.MovieRepository) usescases.GetMovieUseCase {
	return &usescases.UseCaseGetMovie{
		MovieRepository: repository,
	}
}

func newFindMovieByIdUseCase(repository ports.MovieRepository) usescases.GetMovieByIdUseCase {
	return &usescases.UseCaseGetMovieById{
		MovieRepository: repository,
	}
}
func getMovieRepository() ports.MovieRepository {
	return nil
}
