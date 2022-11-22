package db

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/redhatinsights/platform-changelog-go/internal/config"
	l "github.com/redhatinsights/platform-changelog-go/internal/logging"
	"github.com/redhatinsights/platform-changelog-go/internal/metrics"
	"github.com/redhatinsights/platform-changelog-go/internal/models"
	"github.com/redhatinsights/platform-changelog-go/internal/structs"
)

func (conn *DBConnectorImpl) CreateServiceTableEntry(name string, s config.Service) (service models.Services, err error) {
	newService := models.Services{Name: name, DisplayName: s.DisplayName, Tenant: s.Tenant, GHRepo: s.GHRepo, GLRepo: s.GLRepo, Branch: s.Branch, Namespace: s.Namespace, DeployFile: s.DeployFile}
	results := conn.db.Create(&newService)

	return newService, results.Error
}

func (conn *DBConnectorImpl) GetServicesAll(offset int, limit int) ([]structs.ExpandedServicesData, int64, error) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetServicesAll)
	defer callDurationTimer.ObserveDuration()

	var count int64
	var services []structs.ExpandedServicesData

	dbQuery := conn.db.Model(models.Services{})
	dbQuery.Find(&services).Count(&count)

	result := dbQuery.Limit(limit).Offset(offset).Find(&services)

	var servicesWithTimelines []structs.ExpandedServicesData
	for i := 0; i < len(services); i++ {
		s, _, _ := conn.GetLatest(services[i])

		servicesWithTimelines = append(servicesWithTimelines, s)
	}

	return servicesWithTimelines, count, result.Error
}

func (conn *DBConnectorImpl) GetLatest(service structs.ExpandedServicesData) (structs.ExpandedServicesData, error, error) {
	l.Log.Debugf("Query name: %s", service.Name)

	// TODO: Make one query to get the latest commit and deploy for each service
	comResult := conn.db.Model(models.Timelines{}).Select("*").Joins("JOIN services ON timelines.service_id = services.id").Where("services.name = ?", service.Name).Where("timelines.type = ?", "commit").Order("Timestamp desc").Limit(1).Find(&service.Commit)

	depResult := conn.db.Model(models.Timelines{}).Select("*").Joins("JOIN services ON timelines.service_id = services.id").Where("services.name = ?", service.Name).Where("timelines.type = ?", "deploy").Order("Timestamp desc").Limit(1).Find(&service.Deploy)

	return service, comResult.Error, depResult.Error
}

func (conn *DBConnectorImpl) GetServiceByName(name string) (structs.ServicesData, int64, error) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetServiceByName)
	defer callDurationTimer.ObserveDuration()
	var service structs.ServicesData
	result := conn.db.Model(models.Services{}).Where("name = ?", name).First(&service)
	return service, result.RowsAffected, result.Error
}

func (conn *DBConnectorImpl) GetServiceByGHRepo(service_url string) (structs.ServicesData, error) {
	var service structs.ServicesData
	result := conn.db.Model(models.Services{}).Where("gh_repo = ?", service_url).First(&service)

	return service, result.Error
}
