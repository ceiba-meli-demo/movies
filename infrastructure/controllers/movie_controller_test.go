package controllers

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestHandler_Get(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "nginx",
		ExposedPorts: []string{"8080/tcp"},
		WaitingFor:   wait.ForHTTP("/"),
	}
	nginxC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Error(err)
	}
	defer nginxC.Terminate(ctx)
	ip, err := nginxC.Host(ctx)
	if err != nil {
		t.Error(err)
	}
	port, err := nginxC.MappedPort(ctx, "8080")
	if err != nil {
		t.Error(err)
	}
	resp, err := http.Get(fmt.Sprintf("http://%s:%s", ip, port.Port()))
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d. Got %d.", http.StatusOK, resp.StatusCode)
	}
}
