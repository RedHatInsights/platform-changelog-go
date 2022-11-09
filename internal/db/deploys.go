package db

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/redhatinsights/platform-changelog-go/internal/metrics"
	"github.com/redhatinsights/platform-changelog-go/internal/models"
	"github.com/redhatinsights/platform-changelog-go/internal/structs"
	"gorm.io/gorm"
)

func GetDeploysAll(db *gorm.DB, offset int, limit int, q structs.Query) (*gorm.DB, []models.Timelines, int64) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetDeploysAll)
	defer callDurationTimer.ObserveDuration()

	var count int64
	var deploys []models.Timelines

	db = db.Model(models.Timelines{}).Where("timelines.type = ?", "deploy")

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

	db = FilterTimelineByDate(db, q.Start_Date, q.End_Date)

	db.Find(&deploys).Count(&count)
	result := db.Order("Timestamp desc").Limit(limit).Offset(offset).Find(&deploys)

	return result, deploys, count
}

func GetDeploysByService(db *gorm.DB, service structs.ServicesData, offset int, limit int, q structs.Query) (*gorm.DB, []models.Timelines, int64) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetDeploysByService)
	defer callDurationTimer.ObserveDuration()

	var count int64
	var deploys []models.Timelines

	db = db.Model(models.Timelines{}).Where("timelines.service_id = ?", service.ID).Where("timelines.type = ?", "deploy")

	db = FilterTimelineByDate(db, q.Start_Date, q.End_Date)

	db.Find(&deploys).Count(&count)
	result := db.Order("Timestamp desc").Limit(limit).Offset(offset).Find(&deploys)

	return result, deploys, count
}

func GetDeployByRef(db *gorm.DB, ref string) (*gorm.DB, models.Timelines) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetDeployByRef)
	defer callDurationTimer.ObserveDuration()
	var deploy models.Timelines
	result := db.Model(models.Timelines{}).Where("timelines.ref = ?", ref).Where("timelines.type = ?", "deploy").Scan(&deploy)
	return result, deploy
}
