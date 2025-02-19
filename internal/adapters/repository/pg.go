package repository

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq" // Required for PostgreSQL
)

// PostgresRepo struct for database interactions
type PostgresRepo struct {
	db     *pg.DB
	logger *logrus.Logger
}

// NewRepository initializes a new repository
func NewRepository(db *pg.DB, logger *logrus.Logger) (*PostgresRepo, error) {
	p := &PostgresRepo{db: db, logger: logger}
	return p, p.runMigrations()
}

func (r *PostgresRepo) runMigrations() error {
	// run migrations
	collection := migrations.NewCollection()
	err := collection.DiscoverSQLMigrations("migrations")
	if err != nil {
		return err
	}

	// start the migrations
	_, _, err = collection.Run(r.db, "init")
	if err != nil {
		r.logger.WithError(err).Error("Starting PostgreSQL migrations")
		return fmt.Errorf("starting PostgreSQL migrations: %w", err)
	}

	oldVersion, newVersion, err := collection.Run(r.db, "up")
	if err != nil {
		r.logger.WithError(err).Error("Running PostgreSQL migrations")
		return fmt.Errorf("running PostgreSQL migrations: %w", err)
	}

	if newVersion != oldVersion {
		r.logger.WithFields(logrus.Fields{
			"old_version": oldVersion, "new_version": newVersion,
		}).Info("new database migration")
	} else {
		r.logger.WithField("old_version", oldVersion).Info("migration version")
	}

	return err
}
