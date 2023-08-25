package db

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/redhatinsights/platform-changelog-go/internal/config"
	l "github.com/redhatinsights/platform-changelog-go/internal/logging"
	"github.com/redhatinsights/platform-changelog-go/internal/metrics"
	"github.com/redhatinsights/platform-changelog-go/internal/models"
	"github.com/redhatinsights/platform-changelog-go/internal/structs"
)

func (conn *DBConnectorImpl) CreateServiceTableEntry(s models.Services) (service models.Services, err error) {
	results := conn.db.Create(&s)

	return s, results.Error
}

func (conn *DBConnectorImpl) UpdateServiceTableEntry(name string, s config.Service) (service models.Services, err error) {
	newService := models.Services{Name: name, DisplayName: s.DisplayName, Tenant: s.Tenant}
	results := conn.db.Model(models.Services{}).Where("name = ?", name).Updates(&newService)

	return newService, results.Error
}

func (conn *DBConnectorImpl) DeleteServiceTableEntry(name string) (structs.ServicesData, error) {
	// save the service to delete the timelines
	service, _, _ := conn.GetServiceByName(name)

	results := conn.db.Model(models.Services{}).Where("name = ?", name).Delete(&models.Services{})
	if results.Error != nil {
		return structs.ServicesData{}, results.Error
	}

	// delete the timelines for the service
	err := conn.DeleteTimelinesByService(service)
	if err != nil {
		return structs.ServicesData{}, err
	}

	return service, results.Error
}

func (conn *DBConnectorImpl) GetServicesAll(offset int, limit int, q structs.Query) ([]structs.ExpandedServicesData, int64, error) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetServicesAll)
	defer callDurationTimer.ObserveDuration()

	var count int64
	var services []structs.ExpandedServicesData

	db := conn.db.Model(models.Services{})

	if len(q.Name) > 0 {
		db = db.Where("services.name IN ?", q.Name)
	}
	if len(q.DisplayName) > 0 {
		db = db.Where("services.display_name IN ?", q.DisplayName)
	}
	if len(q.Tenant) > 0 {
		db = db.Where("services.tenant IN ?", q.Tenant)
	}

	// Uses the Services model here to reflect the proper db relation
	db.Model(models.Services{}).Count(&count)

	// TODO: add a sort_by field to the query struct
	result := db.Order("ID desc").Limit(limit).Offset(offset).Find(&services)

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

func (conn *DBConnectorImpl) GetServiceNames() ([]string, error) {
	var names []string
	result := conn.db.Model(models.Services{}).Pluck("name", &names)
	return names, result.Error
}

func (conn *DBConnectorImpl) GetServiceByName(name string) (structs.ServicesData, int64, error) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetServiceByName)
	defer callDurationTimer.ObserveDuration()
	var service structs.ServicesData
	result := conn.db.Model(models.Services{}).Where("name = ?", name).First(&service)
	return service, result.RowsAffected, result.Error
}

func (conn *DBConnectorImpl) GetServiceByRepo(repo string) (structs.ServicesData, error) {
	var service structs.ServicesData
	result := conn.db.Model(models.Services{}).Joins("JOIN services on projects.service_id = services.id").Where("repo = ?", repo).First(&service)

	return service, result.Error
}
