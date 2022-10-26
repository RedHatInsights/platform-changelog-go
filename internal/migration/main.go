package main

import (
	"github.com/redhatinsights/platform-changelog-go/internal/config"
	"github.com/redhatinsights/platform-changelog-go/internal/db"
	"github.com/redhatinsights/platform-changelog-go/internal/logging"
)

var conn db.DBConnector

func main() {
	logging.InitLogger()

	cfg := config.Get()

	conn = db.NewDBConnector(cfg)

	conn.Migrate()

	logging.Log.Info("DB Migration Complete")

	reconcileServices(cfg)
}

func reconcileServices(cfg *config.Config) {
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
