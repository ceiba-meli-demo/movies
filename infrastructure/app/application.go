package app

import (
	"github.com/ceiba-meli-demo/movies/application/usescases"
	"github.com/ceiba-meli-demo/movies/domain/ports"
	"github.com/ceiba-meli-demo/movies/infrastructure/adapters/repository/movies"
	"github.com/ceiba-meli-demo/movies/infrastructure/app/middlewares/error_handler"
	"github.com/ceiba-meli-demo/movies/infrastructure/controllers"
	"github.com/ceiba-meli-demo/movies/infrastructure/database_client"

	//Guión bajo usado por propositos de testear la conexión. Modificar cuando se vaya a implementar
	_ "github.com/ceiba-meli-demo/movies/infrastructure/database_client"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	router = gin.Default()
)

type MovieMysqlRepository struct {
	Db *gorm.DB
}

func StartApplication() {
	router.Use(error_handler.ErrorHandler())
	movieRepository := getMovieRepository()
	var handler = createHandler(movieRepository)
	mapUrls(handler)

	//logger.Info("about to start the application")
	_ = router.Run(":8081")
}

func createHandler(movieRepository ports.MovieRepository) controllers.RedirectMovieHandler {
	return newHandler(newCreateMovieUseCase(movieRepository), newGetMoviesUseCase(movieRepository),
		newFindMovieByIdUseCase(movieRepository))
}

func newHandler(createMovie usescases.CreateMoviePort, getMoviesUseCase usescases.GetMovieUseCase, getMovieByID usescases.GetMovieByIDUseCase) controllers.RedirectMovieHandler {
	return &controllers.Handler{
		CreateMovieUseCase:  createMovie,
		GetMoviesUseCase:    getMoviesUseCase,
		GetMovieByIDUseCase: getMovieByID,
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

func newFindMovieByIdUseCase(repository ports.MovieRepository) usescases.GetMovieByIDUseCase {
	return &usescases.UseCaseGetMovieById{
		MovieRepository: repository,
	}
}
func getMovieRepository() ports.MovieRepository {
	return &movies.MovieSqlRepository{
		Connection: database_client.Client,
	}
}
