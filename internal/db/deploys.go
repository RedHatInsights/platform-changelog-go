package db

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/redhatinsights/platform-changelog-go/internal/metrics"
	"github.com/redhatinsights/platform-changelog-go/internal/models"
	"github.com/redhatinsights/platform-changelog-go/internal/structs"
)

func (conn *DBConnectorImpl) GetDeploysAll(offset int, limit int, q structs.Query) ([]models.Timelines, int64, error) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetDeploysAll)
	defer callDurationTimer.ObserveDuration()

	var count int64
	var deploys []models.Timelines

	db := conn.db.Model(models.Timelines{}).Where("timelines.type = ?", "deploy")

	if len(q.Repo) > 0 {
		db = db.Where("timelines.repo IN ?", q.Repo)
	}
	if len(q.Ref) > 0 {
		db = db.Where("timelines.ref IN ?", q.Ref)
	}
	if len(q.Cluster) > 0 {
		db = db.Where("timelines.cluster IN ?", q.Cluster)
	}
	if len(q.Image) > 0 {
		db = db.Where("timelines.image IN ?", q.Image)
	}

	db = FilterTimelineByDate(db, q.StartDate, q.EndDate)

	db.Model(&deploys).Count(&count)
	result := db.Order("Timestamp desc").Limit(limit).Offset(offset).Find(&deploys)

	return deploys, count, result.Error
}

func (conn *DBConnectorImpl) GetDeploysByService(service structs.ServicesData, offset int, limit int, q structs.Query) ([]models.Timelines, int64, error) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetDeploysByService)
	defer callDurationTimer.ObserveDuration()

	var count int64
	var deploys []models.Timelines

	db := conn.db.Model(models.Timelines{}).Where("timelines.service_id = ?", service.ID).Where("timelines.type = ?", "deploy")

	db = FilterTimelineByDate(db, q.StartDate, q.EndDate)

	db.Model(&deploys).Count(&count)
	result := db.Order("Timestamp desc").Limit(limit).Offset(offset).Find(&deploys)

	return deploys, count, result.Error
}

func (conn *DBConnectorImpl) GetDeployByRef(ref string) (models.Timelines, int64, error) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetDeployByRef)
	defer callDurationTimer.ObserveDuration()
	var deploy models.Timelines
	result := conn.db.Model(models.Timelines{}).Where("timelines.ref = ?", ref).Where("timelines.type = ?", "deploy").Scan(&deploy)
	rowsAffected := result.RowsAffected

	return deploy, rowsAffected, result.Error
}
