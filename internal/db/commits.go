package db

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/redhatinsights/platform-changelog-go/internal/metrics"
	"github.com/redhatinsights/platform-changelog-go/internal/models"
)

func (conn *DBConnectorImpl) CreateCommitEntry(t []models.Timelines) error {
	callDurationTimer := prometheus.NewTimer(metrics.SqlCreateCommitEntry)
	defer callDurationTimer.ObserveDuration()

	for _, timeline := range t {
		conn.db.Create(&timeline)
	}

	return conn.db.Error
}

func (conn *DBConnectorImpl) GetCommitsAll(offset int, limit int) ([]models.Timelines, int64, error) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetCommitsAll)
	defer callDurationTimer.ObserveDuration()

	var count int64
	var commits []models.Timelines

	conn.db = conn.db.Model(models.Timelines{}).Where("timelines.type = ?", "commit")

	conn.db.Find(&commits).Count(&count)
	result := conn.db.Order("Timestamp desc").Limit(limit).Offset(offset).Find(&commits)

	return commits, count, result.Error
}

func (conn *DBConnectorImpl) GetCommitsByService(service models.Services, offset int, limit int) ([]models.Timelines, int64, error) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetCommitsByService)
	defer callDurationTimer.ObserveDuration()

	var count int64
	var commits []models.Timelines

	conn.db = conn.db.Model(models.Timelines{}).Where("timelines.service_id = ?", service.ID).Where("timelines.type = ?", "commit")

	conn.db.Find(&commits).Count(&count)
	result := conn.db.Order("Timestamp desc").Limit(limit).Offset(offset).Find(&commits)

	return commits, count, result.Error
}

func (conn *DBConnectorImpl) GetCommitByRef(ref string) (models.Timelines, int64, error) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetCommitByRef)
	defer callDurationTimer.ObserveDuration()
	var commit models.Timelines
	result := conn.db.Model(models.Timelines{}).Where("timelines.ref = ?", ref).Where("timelines.type = ?", "commit").Scan(&commit)

	return commit, result.RowsAffected, result.Error
}
