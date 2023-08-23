package utils

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"

	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/redhatinsights/platform-changelog-go/internal/config"
	"github.com/redhatinsights/platform-changelog-go/internal/logging"
	"github.com/redhatinsights/platform-changelog-go/internal/migrate"
	"gorm.io/gorm"
)

var (
	dsn = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
)

func CreateTestDB(cfg *config.Config, migrationsPath string) (*embeddedpostgres.EmbeddedPostgres, *gorm.DB, error) {
	dbStartTime := time.Now()
	fmt.Println("STARTING TEST DB...")

	embeddedDB := embeddedpostgres.NewDatabase(embeddedpostgres.DefaultConfig().Port(5432).Logger(nil))
	if err := embeddedDB.Start(); err != nil {
		fmt.Println("Error starting embedded postgres: ", err)
		return nil, nil, err
	}

	gres, err := sql.Open("postgres", dsn)
	if err != nil {
		logging.Log.Fatal("Error opening DB: ", err)
		return nil, nil, err
	}

	err = migrate.Migrate(gres, migrationsPath, "up")
	if err != nil {
		fmt.Println("Database migration failed: ", err)
		return nil, nil, err
	}

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logging.Log.Fatal(err)
	}

	logging.Log.Info("DB initialization complete")
	fmt.Println("TEST DB STARTED IN: ", time.Since(dbStartTime))

	return embeddedDB, gormDB, nil
}
