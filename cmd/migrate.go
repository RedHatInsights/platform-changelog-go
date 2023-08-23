package main

import (
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/redhatinsights/platform-changelog-go/internal/config"
	"github.com/redhatinsights/platform-changelog-go/internal/db"
	"github.com/redhatinsights/platform-changelog-go/internal/logging"
	m "github.com/redhatinsights/platform-changelog-go/internal/migrate"
)

var (
	migrationsPath = "file://migrations"
)

func migrateDB(cfg *config.Config, direction string) error {
	logging.Log.Info("Migrating DB")

	gres, err := db.OpenPostgresDB(cfg)
	if err != nil {
		logging.Log.Error(err)
		return err
	}

	return m.Migrate(gres, migrationsPath, direction)
}
