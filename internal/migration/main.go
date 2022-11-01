package main

import (
	"fmt"

	"github.com/redhatinsights/platform-changelog-go/internal/config"
	"github.com/redhatinsights/platform-changelog-go/internal/db"
	"github.com/redhatinsights/platform-changelog-go/internal/logging"
)

func main() {
	logging.InitLogger()

	cfg := config.Get()

	var dbConnector db.DBConnector
	switch cfg.DatabaseConfig.DBImpl {
	case "mock":
		fmt.Println("Using mock database")
		dbConnector = db.NewMockDBConnector()
	default:
		dbConnector = db.NewDBConnector(cfg)
	}

	dbConnector.Migrate()

	logging.Log.Info("DB Migration Complete")

	reconcileServices(cfg, dbConnector)
}

func reconcileServices(cfg *config.Config, conn db.DBConnector) {
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
