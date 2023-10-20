package db

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/redhatinsights/platform-changelog-go/internal/metrics"
	"github.com/redhatinsights/platform-changelog-go/internal/models"
	"github.com/redhatinsights/platform-changelog-go/internal/structs"
)

func (conn *DBConnectorImpl) CreateProjectTableEntry(p *models.Projects) error {
	results := conn.db.Create(p)

	return evaluateError(results.Error)
}

func (conn *DBConnectorImpl) UpdateProjectTableEntry(p *models.Projects) error {
	results := conn.db.Model(models.Projects{}).Where("name = ?", p.Name).Updates(p)

	return evaluateError(results.Error)
}

func (conn *DBConnectorImpl) GetProjectsAll(offset int, limit int, q structs.Query) ([]models.Projects, int64, error) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetServicesAll)
	defer callDurationTimer.ObserveDuration()

	var count int64
	var projects []models.Projects

	db := conn.db.Model(models.Projects{})

	if len(q.Name) > 0 {
		db = db.Where("projects.name IN ?", q.Name)
	}
	if len(q.Namespace) > 0 {
		db = db.Where("projects.namespace IN ?", q.Namespace)
	}
	if len(q.Branch) > 0 {
		db = db.Where("projects.branch IN ?", q.Branch)
	}

	// Uses the Projects model here to reflect the proper db relation
	db.Model(models.Projects{}).Count(&count)

	// TODO: add a sort_by field to the query struct
	result := db.Order("ID desc").Limit(limit).Offset(offset).Find(&projects)

	return projects, count, evaluateError(result.Error)
}

func (conn *DBConnectorImpl) GetProjectsByService(service models.Services, offset int, limit int, q structs.Query) ([]models.Projects, int64, error) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetProjectsByService)
	defer callDurationTimer.ObserveDuration()

	var count int64
	var projects []models.Projects

	db := conn.db.Model(models.Projects{}).Select("*").Where("service_id = ?", service.ID)
	db.Model(models.Projects{}).Count(&count)
	result := db.Order("ID desc").Limit(limit).Offset(offset).Find(&projects)

	return projects, count, evaluateError(result.Error)
}

func (conn *DBConnectorImpl) GetProjectByID(id int) (models.Projects, int64, error) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetProjectByID)
	defer callDurationTimer.ObserveDuration()
	var project models.Projects
	result := conn.db.Model(models.Projects{}).Where("id = ?", id).First(&project)
	return project, result.RowsAffected, evaluateError(result.Error)
}

func (conn *DBConnectorImpl) GetProjectByName(name string) (models.Projects, int64, error) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetProjectByName)
	defer callDurationTimer.ObserveDuration()
	var project models.Projects
	result := conn.db.Model(models.Projects{}).Order("ID desc").Where("name = ?", name).First(&project)
	return project, result.RowsAffected, evaluateError(result.Error)
}

func (conn *DBConnectorImpl) GetProjectByRepo(repo string) (models.Projects, error) {
	var project models.Projects
	result := conn.db.Model(models.Projects{}).Where("repo = ?", repo).First(&project)

	return project, evaluateError(result.Error)
}
