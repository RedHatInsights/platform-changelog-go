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
func GetTimelinesAll(db *gorm.DB, offset int, limit int, q structs.Query) (*gorm.DB, []models.Timelines, int64) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetTimelinesAll)
	defer callDurationTimer.ObserveDuration()

	var count int64
	var timelines []models.Timelines

	// Concatanate the timeline fields
	fields := fmt.Sprintf("%s,%s,%s", strings.Join(timelinesFields, ","), strings.Join(commitsFields, ","), strings.Join(deploysFields, ","))

	db = db.Model(models.Timelines{}).Select(fields)

	if len(q.Repo) > 0 {
		db = db.Where("timelines.repo IN ?", q.Repo)
	}
	if len(q.Ref) > 0 {
		db = db.Where("timelines.ref IN ?", q.Ref)
	}

	db = FilterTimelineByDate(db, q.Start_Date, q.End_Date)

	db.Find(&timelines).Count(&count)
	result := db.Order("Timestamp desc").Limit(limit).Offset(offset).Find(&timelines)

	return result, timelines, count
}

func GetTimelinesByService(db *gorm.DB, service structs.ServicesData, offset int, limit int, q structs.Query) (*gorm.DB, []models.Timelines, int64) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetTimelinesByService)
	defer callDurationTimer.ObserveDuration()

	var count int64
	var timelines []models.Timelines

	// Concatanate the timeline fields
	fields := fmt.Sprintf("%s,%s,%s", strings.Join(timelinesFields, ","), strings.Join(commitsFields, ","), strings.Join(deploysFields, ","))

	db = db.Model(models.Timelines{}).Select(fields).Where("service_id = ?", service.ID)

	db = FilterTimelineByDate(db, q.Start_Date, q.End_Date)

	db.Find(&timelines).Count(&count)
	result := db.Order("Timestamp desc").Limit(limit).Offset(offset).Find(&timelines)

	return result, timelines, count
}

func GetTimelineByRef(db *gorm.DB, ref string) (*gorm.DB, models.Timelines) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetTimelineByRef)
	defer callDurationTimer.ObserveDuration()

	var timeline models.Timelines

	result := db.Model(models.Timelines{}).Select("*").Where("timelines.ref = ?", ref).Find(&timeline)

	return result, timeline
}
