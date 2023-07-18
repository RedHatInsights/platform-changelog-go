package main

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/redhatinsights/platform-changelog-go/internal/config"
	"github.com/redhatinsights/platform-changelog-go/internal/db"
	"github.com/redhatinsights/platform-changelog-go/internal/logging"
)

var (
	migrationsPath = "file://migrations"
)

func migrateDB(cfg *config.Config, direction string) error {
	logging.Log.Info("Migrating DB")

	postgres, err := db.OpenPostgresDB(cfg)
	if err != nil {
		logging.Log.Error(err)
		return err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logging.Log.Error("Error creating postgres driver: ", err)
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(migrationsPath, "postgres", driver)
	if err != nil {
		logging.Log.Fatal("Error creating migration instance: ", err)
		return err
	}

	m.Log = logging.Log

	if direction == "up" {
		err = m.Up()
	} else if direction == "down" {
		err = m.Steps(-1)
	} else {
		logging.Log.Fatal("Invalid migration direction: ", direction)
		return
	}

	if errors.Is(err, migrate.ErrNoChange) {
		logging.Log.Info("No migration changes")
	} else if err != nil {
		logging.Log.Error("DB migration resulted in an error: ", "error", err)
		return err
	}

	logging.Log.Info("DB Migration Complete")

	return nil
}
