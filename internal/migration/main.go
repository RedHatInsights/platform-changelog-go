package main

import (
	"github.com/redhatinsights/platform-changelog-go/internal/config"
	"github.com/redhatinsights/platform-changelog-go/internal/db"
	"github.com/redhatinsights/platform-changelog-go/internal/logging"
	"github.com/redhatinsights/platform-changelog-go/internal/models"

	"gorm.io/gorm"
)

func main() {
	logging.InitLogger()

	cfg := config.Get()

	db.DbConnect(cfg)

	// Set up TimelineType Enum (gorm doesn't have a function for this)
	db.DB.Exec("CREATE TYPE timeline_type AS ENUM ('unknown', 'commit', 'deploy')")

	db.DB.AutoMigrate(
		&models.Services{},
		&models.Timelines{},
	)

	logging.Log.Info("DB Migration Complete")

	reconcileServices(db.DB, cfg)
}

func reconcileServices(g *gorm.DB, cfg *config.Config) {
	for key, service := range cfg.Services {
		// Validate the tenant field exists in the config
		if !validateTenant(service.Tenant, cfg) {
			logging.Log.Error("Tenant not validated: ", service.Tenant)
			continue
		}

		res, _ := db.GetServiceByName(g, key)
		if res.RowsAffected == 0 {
			_, service := db.CreateServiceTableEntry(g, key, service)
			logging.Log.Info("Created service: ", service)
		} else {
			logging.Log.Info("Service already exists: ", service.DisplayName)
		}
	}
}

func validateTenant(tenant string, cfg *config.Config) bool {
	for _, t := range cfg.Tenants {
		if t.Name == tenant {
			return true
		}
	}
	return false
}
