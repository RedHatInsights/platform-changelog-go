package main

import (
	"fmt"

	"github.com/redhatinsights/platform-changelog-go/internal/models"
	"github.com/redhatinsights/platform-changelog-go/internal/structs"

	"github.com/redhatinsights/platform-changelog-go/internal/config"
	"github.com/redhatinsights/platform-changelog-go/internal/db"
	"github.com/redhatinsights/platform-changelog-go/internal/logging"
)

func seedDB(cfg *config.Config) error {
	var dbConnector db.DBConnector
	switch cfg.DBImpl {
	case "mock":
		fmt.Println("Using mock database")
		dbConnector = *db.NewMockDBConnector(cfg)
	default:
		dbConnector = db.NewDBConnector(cfg)
	}

	// Remove services that are no longer in the config
	cleanupServices(cfg, dbConnector)

	reconcileServices(cfg, dbConnector)
}

func cleanupServices(cfg *config.Config, conn db.DBConnectorImpl) {
	services, _ := conn.GetServicesAll()
	for _, service := range services {
		if _, ok := cfg.Services[service.Name]; !ok {
			conn.DeleteServiceTableEntry(service.Name)
		}
	}
}

func reconcileServices(cfg *config.Config, conn db.DBConnectorImpl) {
	for key, service := range cfg.Services {
		// Validate the tenant field exists in the config
		if !validateTenant(service.Tenant, cfg) {
			logging.Log.Error("Tenant not validated: ", service.Tenant)
			continue
		}

		serviceData, rowsAffected, err := conn.GetServiceByName(key)
		if err != nil {
			logging.Log.Error("Error getting service: ", err)
		}

		if rowsAffected == 0 {
			_, err := conn.CreateServiceTableEntry(key, service)
			if err != nil {
				logging.Log.Error("Error creating service: ", err)
				continue // skip to the next service
			}

			serviceData, _, err = conn.GetServiceByName(key) // get the service we just created

			logging.Log.Info("Created service: ", service)
		} else {
			logging.Log.Info("Service already exists: ", service.DisplayName)
		}

		// update the service if fields have changed
		if serviceData == nil {
			logging.Log.Error("Failed to retrieve service data")
			continue
		}

		err = compareService(serviceData, service, conn)
		if err != nil {
			logging.Log.Error("Error comparing and updating service: ", err)
			continue
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

func compareService(fromDB structs.ServicesData, fromCfg models.Services, conn db.DBConnectorImp) error {
	// compare the fields
	if fromDB.Name != fromCfg.Name ||
		fromDB.DisplayName != fromCfg.DisplayName ||
		fromDB.Tenant != fromCfg.Tenant ||
		fromDB.GHRepo != fromCfg.GHRepo ||
		fromDB.GLRepo != fromCfg.GLRepo ||
		fromDB.Branch != fromCfg.Branch ||
		fromDB.Namespace != fromCfg.Namespace ||
		fromDB.DeployFile != fromCfg.DeployFile {
		// update the service
		_, err := conn.UpdateServiceTableEntry(fromCfg.Name, fromCfg)
		if err != nil {
			logging.Log.Error("Error updating service: ", err)
			return err
		}

		logging.Log.Info("Updated service: ", fromCfg.DisplayName)
	}

	return nil
}
