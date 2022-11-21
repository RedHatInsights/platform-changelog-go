package main

import (
	"github.com/redhatinsights/platform-changelog-go/internal/config"
	"github.com/redhatinsights/platform-changelog-go/internal/db"
	"github.com/redhatinsights/platform-changelog-go/internal/logging"
	"github.com/redhatinsights/platform-changelog-go/internal/models"
)

func main() {
	logging.InitLogger()

	cfg := config.Get()

	dbConnector := *db.NewDBConnector(cfg)

	Migrate(dbConnector)

	logging.Log.Info("DB Migration Complete")

	ReconcileServices(cfg, dbConnector)
}

func Migrate(conn db.DBConnectorImpl) {
	conn.Exec("CREATE TYPE timeline_type AS ENUM ('unknown', 'commit', 'deploy')")

	conn.AutoMigrate(
		&models.Services{},
		&models.Timelines{},
	)

	logging.Log.Info("DB Migration Complete")
}

func ReconcileServices(cfg *config.Config, conn db.DBConnectorImpl) {
	for key, service := range cfg.Services {
		_, rowsAffected, _ := conn.GetServiceByName(key)
		if rowsAffected == 0 {
			_, service := conn.CreateServiceTableEntry(key, service)
			logging.Log.Info("Created service: ", service)
		} else {
			logging.Log.Info("Service already exists: ", service.DisplayName)
		}
	}
}
