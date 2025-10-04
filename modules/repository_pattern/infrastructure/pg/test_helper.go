package pg

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PostgresContainer struct {
	container *postgres.PostgresContainer
	DSN       string
	Terminate func()
}

func WithPostgresContainer(ctx context.Context, fixture string, t *testing.T) (*PostgresContainer, error) {
	dbName := "testdb"
	dbUser := "testuser"
	dbPassword := "testpassword"

	scripts, err := readInitScripts()
	if err != nil {
		return nil, err
	}
	scripts = append(scripts, fixture)

	container, err := postgres.Run(ctx, "postgres:16-alpine",
		postgres.WithInitScripts(scripts...),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second),
		),
	)
	if err != nil {
		return nil, err
	}

	terminate := func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %v", err)
		}
	}

	connStr, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &PostgresContainer{
		container: container,
		DSN:       connStr,
		Terminate: terminate,
	}, nil
}

func readInitScripts() ([]string, error) {
	projectRoot := "../../../../.."
	scriptsDir := projectRoot + "/migrations/"

	entries, err := os.ReadDir(scriptsDir)
	if err != nil {
		return nil, err
	}

	scripts := make([]string, 0, len(entries))
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), "up.sql") {
			scripts = append(scripts, scriptsDir+e.Name())
		}
	}

	return scripts, nil
}
