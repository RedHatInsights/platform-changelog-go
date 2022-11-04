package db

import (
	"fmt"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/redhatinsights/platform-changelog-go/internal/metrics"
	"github.com/redhatinsights/platform-changelog-go/internal/models"
	"github.com/redhatinsights/platform-changelog-go/internal/structs"
	"gorm.io/gorm"
)

/**
 * GetTimeline returns a timeline of commits and deploys for a service
 */
func (conn *DBConnectorImpl) GetTimelinesAll(offset int, limit int) ([]models.Timelines, int64, error) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetTimelinesAll)
	defer callDurationTimer.ObserveDuration()

	var count int64
	var timelines []models.Timelines

	// Concatanate the timeline fields
	fields := fmt.Sprintf("%s,%s,%s", strings.Join(timelinesFields, ","), strings.Join(commitsFields, ","), strings.Join(deploysFields, ","))

	conn.db = conn.db.Model(models.Timelines{}).Select(fields).Session(&gorm.Session{})

	conn.db.Find(&timelines).Count(&count)
	result := conn.db.Order("Timestamp desc").Limit(limit).Offset(offset).Find(&timelines)

	return timelines, count, result.Error
}

func (conn *DBConnectorImpl) GetTimelinesByService(service structs.ServicesData, offset int, limit int) ([]models.Timelines, int64, error) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetTimelinesByService)
	defer callDurationTimer.ObserveDuration()

	var count int64
	var timelines []models.Timelines

	// Concatanate the timeline fields
	fields := fmt.Sprintf("%s,%s,%s", strings.Join(timelinesFields, ","), strings.Join(commitsFields, ","), strings.Join(deploysFields, ","))

	conn.db = conn.db.Model(models.Timelines{}).Select(fields).Where("service_id = ?", service.ID).Session(&gorm.Session{})

	conn.db.Find(&timelines).Count(&count)
	result := conn.db.Order("Timestamp desc").Limit(limit).Offset(offset).Find(&timelines)

	return timelines, count, result.Error
}

func (conn *DBConnectorImpl) GetTimelineByRef(ref string) (models.Timelines, int64, error) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetTimelineByRef)
	defer callDurationTimer.ObserveDuration()

	var timeline models.Timelines

	result := conn.db.Model(models.Timelines{}).Select("*").Where("timelines.ref = ?", ref).Find(&timeline).Session(&gorm.Session{})

	return timeline, result.RowsAffected, result.Error
}
