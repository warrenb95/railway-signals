package repository_test

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/go-pg/pg/v10"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/sirupsen/logrus"

	"github.com/warrenb95/railway-signals/internal/adapters/repository"
)

var (
	underlyingDB *pg.DB
	testDB       *repository.PostgresRepository
)

func TestMain(m *testing.M) {
	// Create a new Docker pool
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// Run a PostgreSQL container
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "15", // Change to your desired PostgreSQL version
		Env: []string{
			"POSTGRES_USER=testuser",
			"POSTGRES_PASSWORD=testpass",
			"POSTGRES_DB=testdb",
		},
	}, func(config *docker.HostConfig) {
		// Auto-remove the container after tests finish
		config.AutoRemove = true
		// Allow connections from the host machine
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// Retry connecting until the database is ready
	pool.MaxWait = 30 * time.Second
	err = pool.Retry(func() error {
		underlyingDB = pg.Connect(&pg.Options{
			// Addr:     "localhost:5432",
			Addr:     resource.GetHostPort("5432/tcp"),
			User:     "testuser",
			Password: "testpass",
			Database: "testdb",
		})
		if err != nil {
			return err
		}
		return underlyingDB.Ping(context.Background())
	})
	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	testDB, err = repository.NewPostgresRepository(underlyingDB, logrus.New())
	if err != nil {
		log.Fatalf("NewPostgresRepository: %s", err)
	}

	// Run tests
	code := m.Run()

	// Clean up
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}
