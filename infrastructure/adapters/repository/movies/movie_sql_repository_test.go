package movies

import (
	"context"
	"fmt"
	"github.com/ceiba-meli-demo/movies/domain/model"
	_ "github.com/ceiba-meli-demo/movies/infrastructure/adapters/repository/models"
	"github.com/ceiba-meli-demo/movies/infrastructure/database_client"
	_ "github.com/ceiba-meli-demo/movies/infrastructure/mappers/movie_mapper"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"os"
	"testing"
)

var (
	movieSqlRepository MovieSqlRepository
)

func TestMain(m *testing.M) {
	fmt.Println("about to start oauth tests")
	containerMockServer, ctx := load()
	code := m.Run()
	beforeAll(containerMockServer, ctx)
	os.Exit(code)
}
func load() (testcontainers.Container, context.Context) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "mongo:latest",
		ExposedPorts: []string{"27017/tcp"},
		Env: map[string]string{
			"MONGO_INITDB_ROOT_USERNAME": "mongousername",
			"MONGO_INITDB_ROOT_PASSWORD":      "password",
		},
		WaitingFor: wait.ForLog("Listening on 0.0.0.0"),
	}
	mongoC, _ := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	host, _ := mongoC.Host(ctx)
	p, _ := mongoC.MappedPort(ctx, "27017/tcp")
	port := p.Port()
	_ = os.Setenv("MONGODB_HOST", host)
	_ = os.Setenv("MONGODB_PORT", port)
	_ = os.Setenv("MONGODB_USERNAME", "mongousername")
	_ = os.Setenv("MONGODB_PASSWORD", "password")

	movieSqlRepository = MovieSqlRepository{
		Connection:database_client.GetDatabaseInstance(),
	}
	return mongoC, ctx
}
func beforeAll(container testcontainers.Container, ctx context.Context) {
	_ = container.Terminate(ctx)
}
func TestMovieSqlRepository_Save(t *testing.T) {
	var movie model.Movie
	movie, _ = movie.CreateMovie("Mi peli", 2, "synopsit")
	err := movieSqlRepository.Save(&movie)
	assert.Nil(t, err)
	assert.EqualValues(t, movie.Title, "Mi peli")
	assert.NotEqual(t, movie.Title, "sistemas31")
}

func TestMovieSqlRepository_GetById(t *testing.T) {
	var movieModel model.Movie
	movieModel, _ = movieModel.CreateMovie("Mi peli", 2, "synopsit")
	if err := movieSqlRepository.Save(&movieModel); err != nil {
		assert.Fail(t, err.Error())
	}
	movieResult, err := movieSqlRepository.GetById(movieModel.ID)
	assert.Nil(t, err)
	assert.NotNil(t, movieResult)
	assert.NotEqual(t, movieResult.ID, 5)
	assert.EqualValues(t, movieResult.ID, movieModel.ID)
}
