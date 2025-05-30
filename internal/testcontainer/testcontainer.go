package testcontainer

import (
	"context"
	"log/slog"
	"os"

	"github.com/darrenvechain/thorgo/thorest"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func NewSolo() (*thorest.Client, func()) {
	ctx := context.Background()
	cmd := []string{"solo", "-api-addr", "0.0.0.0:8669", "-api-cors", "*", "-on-demand", "-api-allowed-tracers", "call"}

	imageEnv, ok := os.LookupEnv("SOLO_IMAGE")
	if !ok {
		imageEnv = "ghcr.io/vechain/thor:release-galactica-latest"
	}

	req := testcontainers.ContainerRequest{
		Image:        imageEnv,
		ExposedPorts: []string{"8669"},
		WaitingFor:   wait.ForLog("new block packed"),
		Cmd:          cmd,
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		slog.Error("failed to create solo container", "error", err)
		os.Exit(1)
	}
	cancel := func() {
		err := container.Terminate(ctx)
		if err != nil {
			slog.Warn("failed to terminate solo container", "error", err)
		}
	}

	endpoint, err := container.Endpoint(ctx, "")
	if err != nil {
		slog.Error("failed to get solo container endpoint", "error", err)
		os.Exit(1)
	}

	return thorest.NewClientFromURL("http://" + endpoint), cancel
}
