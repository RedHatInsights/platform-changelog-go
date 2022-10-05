package db

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/redhatinsights/platform-changelog-go/internal/config"
	l "github.com/redhatinsights/platform-changelog-go/internal/logging"
	"github.com/redhatinsights/platform-changelog-go/internal/metrics"
	"github.com/redhatinsights/platform-changelog-go/internal/models"
	"github.com/redhatinsights/platform-changelog-go/internal/structs"
	"gorm.io/gorm"
)

func CreateServiceTableEntry(db *gorm.DB, name string, s config.Service) (service models.Services, err error) {
	newService := models.Services{Name: name, DisplayName: s.DisplayName, GHRepo: s.GHRepo, GLRepo: s.GLRepo, Branch: s.Branch, Namespace: s.Namespace, DeployFile: s.DeployFile}
	results := db.Create(&newService)

	return newService, results.Error
}

func GetServicesAll(db *gorm.DB, offset int, limit int) ([]structs.ExpandedServicesData, int64, error) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetServicesAll)
	defer callDurationTimer.ObserveDuration()

	var count int64
	var services []structs.ExpandedServicesData

	dbQuery := db.Model(models.Services{})
	dbQuery.Find(&services).Count(&count)

	result := dbQuery.Limit(limit).Offset(offset).Find(&services)

	var servicesWithTimelines []structs.ExpandedServicesData
	for i := 0; i < len(services); i++ {
		s, _, _ := GetLatest(db, services[i])

		servicesWithTimelines = append(servicesWithTimelines, s)
	}

	return servicesWithTimelines, count, result.Error
}

func GetLatest(db *gorm.DB, service structs.ExpandedServicesData) (structs.ExpandedServicesData, error, error) {
	l.Log.Debugf("Query name: %s", service.Name)

	// TODO: Make one query to get the latest commit and deploy for each service
	comResult := db.Model(models.Timelines{}).Select("*").Joins("JOIN services ON timelines.service_id = services.id").Where("services.name = ?", service.Name).Where("timelines.type = ?", "commit").Order("Timestamp desc").Limit(1).Find(&service.Commit)

	depResult := db.Model(models.Timelines{}).Select("*").Joins("JOIN services ON timelines.service_id = services.id").Where("services.name = ?", service.Name).Where("timelines.type = ?", "deploy").Order("Timestamp desc").Limit(1).Find(&service.Deploy)

	return service, comResult.Error, depResult.Error
}

func GetServiceByName(db *gorm.DB, name string) (structs.ServicesData, int64, error) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetServiceByName)
	defer callDurationTimer.ObserveDuration()
	var service structs.ServicesData
	result := db.Model(models.Services{}).Where("name = ?", name).First(&service)
	return service, result.RowsAffected, result.Error
}

func GetServiceByGHRepo(db *gorm.DB, service_url string) (structs.ServicesData, error) {
	var service structs.ServicesData
	result := db.Model(models.Services{}).Where("gh_repo = ?", service_url).First(&service)

	return service, result.Error
}
