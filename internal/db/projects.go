package db

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/redhatinsights/platform-changelog-go/internal/metrics"
	"github.com/redhatinsights/platform-changelog-go/internal/models"
	"github.com/redhatinsights/platform-changelog-go/internal/structs"
)

func (conn *DBConnectorImpl) CreateProjectTableEntry(p models.Projects) (err error) {
	results := conn.db.Create(&p)

	return results.Error
}

func (conn *DBConnectorImpl) UpdateProjectTableEntry(p structs.ProjectsData) (project models.Projects, err error) {
	project = models.Projects{Name: p.Name, Repo: p.Repo, DeployFile: p.DeployFile, Namespaces: p.Namespaces, Branches: p.Branches}

	results := conn.db.Model(models.Projects{}).Where("name = ?", p.Name).Updates(&project)

	return project, results.Error
}

func (conn *DBConnectorImpl) GetProjectsAll(offset int, limit int, q structs.Query) ([]structs.ProjectsData, int64, error) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetServicesAll)
	defer callDurationTimer.ObserveDuration()

	var count int64
	var projects []structs.ProjectsData

	db := conn.db.Model(models.Projects{})

	if len(q.Name) > 0 {
		db = db.Where("projects.name IN ?", q.Name)
	}

	// will these be able to find the namespace/branch in an array properly?
	if len(q.Namespace) > 0 {
		db = db.Where("projects.namespaces IN ?", q.Namespace)
	}
	if len(q.Branch) > 0 {
		db = db.Where("projects.branches IN ?", q.Branch)
	}

	// Uses the Projects model here to reflect the proper db relation
	db.Model(models.Projects{}).Count(&count)

	// TODO: add a sort_by field to the query struct
	result := db.Order("ID desc").Limit(limit).Offset(offset).Find(&projects)

	return projects, count, result.Error
}

func (conn *DBConnectorImpl) GetProjectsByService(service structs.ServicesData, offset int, limit int, q structs.Query) ([]structs.ProjectsData, int64, error) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetProjectsByService)
	defer callDurationTimer.ObserveDuration()

	var count int64
	var projects []structs.ProjectsData

	db := conn.db.Model(models.Projects{}).Select("*").Where("service_id = ?", service.ID)

	db = FilterTimelineByDate(db, q.StartDate, q.EndDate)

	db.Model(models.Projects{}).Count(&count)
	result := db.Order("Timestamp desc").Order("ID desc").Limit(limit).Offset(offset).Find(&projects)

	return projects, count, result.Error
}

func (conn *DBConnectorImpl) GetProjectByName(name string) (structs.ProjectsData, int64, error) {
	callDurationTimer := prometheus.NewTimer(metrics.SqlGetProjectByName)
	defer callDurationTimer.ObserveDuration()
	var project structs.ProjectsData
	result := conn.db.Model(models.Projects{}).Where("name = ?", name).First(&project)
	return project, result.RowsAffected, result.Error
}

func (conn *DBConnectorImpl) GetProjectByRepo(repo string) (structs.ProjectsData, error) {
	var project structs.ProjectsData
	result := conn.db.Model(models.Projects{}).Where("repo = ?", repo).First(&project)

	return project, result.Error
}
