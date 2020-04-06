package usescases

import (
	"context"
	"os"
	"testing"

	"github.com/ceiba-meli-demo/movies/domain/ports"
	"github.com/ceiba-meli-demo/movies/infrastructure/adapters/repository/movies"
	"github.com/ceiba-meli-demo/movies/infrastructure/database_client"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	useCaseForTest UseCaseMovieCreate
)

func TestMain(m *testing.M) {
	containerMockServer, ctx := load()
	code := m.Run()
	beforeAll(ctx, containerMockServer)
	os.Exit(code)
}
func load() (testcontainers.Container, context.Context) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "mongo:latest",
		ExposedPorts: []string{"27017/tcp"},
		Env: map[string]string{
			"MONGO_INITDB_ROOT_USERNAME": "mongousername",
			"MONGO_INITDB_ROOT_PASSWORD": "password",
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

	useCaseForTest = UseCaseMovieCreate{
		MovieRepository: getMovieRepository(),
	}
	return mongoC, ctx
}
func beforeAll(ctx context.Context, container testcontainers.Container) {
	_ = container.Terminate(ctx)
}

func TestHandlerCreateMovie(t *testing.T) {

}

func getMovieRepository() ports.MovieRepository {
	return &movies.MovieNoSQLRepository{
		Connection: database_client.GetDatabaseInstance(),
	}
}
