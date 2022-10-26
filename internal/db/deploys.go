package db

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/redhatinsights/platform-changelog-go/internal/metrics"
	"github.com/redhatinsights/platform-changelog-go/internal/models"
	"gorm.io/gorm"
)

func GetDeploysAll(db *gorm.DB, offset int, limit int) ([]models.Timelines, int64, error) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetDeploysAll)
	defer callDurationTimer.ObserveDuration()

	var count int64
	var deploys []models.Timelines

	db = db.Model(models.Timelines{}).Where("timelines.type = ?", "deploy")

	db.Find(&deploys).Count(&count)
	result := db.Order("Timestamp desc").Limit(limit).Offset(offset).Find(&deploys)

	return deploys, count, result.Error
}

func GetDeploysByService(db *gorm.DB, service models.Services, offset int, limit int) ([]models.Timelines, int64, error) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetDeploysByService)
	defer callDurationTimer.ObserveDuration()

	var count int64
	var deploys []models.Timelines

	db = db.Model(models.Timelines{}).Where("timelines.service_id = ?", service.ID).Where("timelines.type = ?", "deploy")

	db.Find(&deploys).Count(&count)
	result := db.Order("Timestamp desc").Limit(limit).Offset(offset).Find(&deploys)

	return deploys, count, result.Error
}

func GetDeployByRef(db *gorm.DB, ref string) (models.Timelines, int64, error) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetDeployByRef)
	defer callDurationTimer.ObserveDuration()
	var deploy models.Timelines
	result := db.Model(models.Timelines{}).Where("timelines.ref = ?", ref).Where("timelines.type = ?", "deploy").Scan(&deploy)
	rowsAffected := result.RowsAffected

	return deploy, rowsAffected, result.Error
}
