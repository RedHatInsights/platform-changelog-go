package db

import (
	"github.com/redhatinsights/platform-changelog-go/internal/config"
	"github.com/redhatinsights/platform-changelog-go/internal/models"
	"github.com/redhatinsights/platform-changelog-go/internal/structs"
)

type DBConnector interface {
	CreateServiceTableEntry(s models.Services) (models.Services, error)
	UpdateServiceTableEntry(name string, s config.Service) (service models.Services, err error)
	DeleteServiceTableEntry(name string) (models.Services, error)
	GetServicesAll(offset int, limit int, q structs.Query) ([]structs.ExpandedServicesData, int64, error)
	GetLatest(service structs.ExpandedServicesData) (structs.ExpandedServicesData, error, error)
	GetServiceByName(name string) (models.Services, int64, error)
	GetServiceByRepo(repo string) (structs.ServicesData, error)

	CreateProjectTableEntry(p models.Projects) (err error)
	UpdateProjectTableEntry(p structs.ProjectsData) (project models.Projects, err error)
	GetProjectsAll(offset int, limit int, q structs.Query) ([]structs.ProjectsData, int64, error)
	GetProjectsByService(service models.Services, offset int, limit int, q structs.Query) ([]structs.ProjectsData, int64, error)
	GetProjectByName(name string) (structs.ProjectsData, int64, error)
	GetProjectByRepo(repo string) (structs.ProjectsData, error)

	GetTimelinesAll(offset int, limit int, q structs.Query) ([]models.Timelines, int64, error)
	GetTimelinesByService(service models.Services, offset int, limit int, q structs.Query) ([]models.Timelines, int64, error)
	GetTimelinesByProject(project structs.ProjectsData, offset int, limit int, q structs.Query) ([]models.Timelines, int64, error)
	GetTimelineByRef(ref string) (models.Timelines, int64, error)
	DeleteTimelinesByService(service models.Services) error

	CreateCommitEntry(timeline models.Timelines) error
	BulkCreateCommitEntry(timeline []models.Timelines) error
	GetCommitsAll(offset int, limit int, q structs.Query) ([]models.Timelines, int64, error)
	GetCommitsByService(service models.Services, offset int, limit int, q structs.Query) ([]models.Timelines, int64, error)
	GetCommitsByProject(project structs.ProjectsData, offset int, limit int, q structs.Query) ([]models.Timelines, int64, error)
	GetCommitByRef(ref string) (models.Timelines, int64, error)

	CreateDeployEntry(timeline models.Timelines) error
	GetDeploysAll(offset int, limit int, q structs.Query) ([]models.Timelines, int64, error)
	GetDeploysByService(service models.Services, offset int, limit int, q structs.Query) ([]models.Timelines, int64, error)
	GetDeploysByProject(project structs.ProjectsData, offset int, limit int, q structs.Query) ([]models.Timelines, int64, error)
	GetDeployByRef(ref string) (models.Timelines, int64, error)
}
