package db

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/redhatinsights/platform-changelog-go/internal/metrics"
	"github.com/redhatinsights/platform-changelog-go/internal/models"
	"github.com/redhatinsights/platform-changelog-go/internal/structs"
)

func (conn *DBConnectorImpl) CreateCommitEntry(t models.Timelines) error {
	callDurationTimer := prometheus.NewTimer(metrics.SqlCreateCommitEntry)
	defer callDurationTimer.ObserveDuration()

	conn.db.Create(&t)

	return evaluateError(conn.db.Error)
}

func (conn *DBConnectorImpl) BulkCreateCommitEntry(t []models.Timelines) error {
	callDurationTimer := prometheus.NewTimer(metrics.SqlCreateCommitEntry)
	defer callDurationTimer.ObserveDuration()

	for _, timeline := range t {
		conn.db.Create(&timeline)
	}

	return evaluateError(conn.db.Error)
}

func (conn *DBConnectorImpl) GetCommitsAll(offset int, limit int, q structs.Query) ([]models.Timelines, int64, error) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetCommitsAll)
	defer callDurationTimer.ObserveDuration()

	var count int64
	var commits []models.Timelines

	db := conn.db.Model(models.Timelines{}).Where("timelines.type = ?", "commit")

	if len(q.Repo) > 0 {
		db = db.Where("timelines.repo IN ?", q.Repo)
	}
	if len(q.Author) > 0 {
		db = db.Where("timelines.author IN ?", q.Author)
	}
	if len(q.MergedBy) > 0 {
		db = db.Where("timelines.merged_by IN ?", q.MergedBy)
	}
	if len(q.Ref) > 0 {
		db = db.Where("timelines.ref IN ?", q.Ref)
	}

	db = FilterTimelineByDate(db, q.StartDate, q.EndDate)

	db.Model(&commits).Count(&count)
	result := db.Order("Timestamp desc").Order("ID desc").Limit(limit).Offset(offset).Find(&commits)

	return commits, count, evaluateError(result.Error)
}

func (conn *DBConnectorImpl) GetCommitsByService(service models.Services, offset int, limit int, q structs.Query) ([]models.Timelines, int64, error) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetCommitsByService)
	defer callDurationTimer.ObserveDuration()

	var count int64
	var commits []models.Timelines

	db := conn.db.Model(models.Timelines{}).Where("timelines.service_id = ?", service.ID).Where("timelines.type = ?", "commit")

	db = FilterTimelineByDate(db, q.StartDate, q.EndDate)

	db.Model(&commits).Count(&count)
	result := db.Order("Timestamp desc").Order("ID desc").Limit(limit).Offset(offset).Find(&commits)

	return commits, count, evaluateError(result.Error)
}

func (conn *DBConnectorImpl) GetCommitsByProject(project models.Projects, offset int, limit int, q structs.Query) ([]models.Timelines, int64, error) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetCommitsByProject)
	defer callDurationTimer.ObserveDuration()

	var count int64
	var commits []models.Timelines

	db := conn.db.Model(models.Timelines{}).Where("timelines.project_id = ?", project.ID).Where("timelines.type = ?", "commit")

	db = FilterTimelineByDate(db, q.StartDate, q.EndDate)

	db.Model(&commits).Count(&count)
	result := db.Order("Timestamp desc").Order("ID desc").Limit(limit).Offset(offset).Find(&commits)

	return commits, count, evaluateError(result.Error)
}

func (conn *DBConnectorImpl) GetCommitByRef(ref string) (models.Timelines, int64, error) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetCommitByRef)
	defer callDurationTimer.ObserveDuration()
	var commit models.Timelines
	result := conn.db.Model(models.Timelines{}).Where("timelines.ref = ?", ref).Where("timelines.type = ?", "commit").Scan(&commit)

	return commit, result.RowsAffected, evaluateError(result.Error)
}
