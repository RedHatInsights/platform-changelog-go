package db

import (
	"github.com/redhatinsights/platform-changelog-go/internal/config"
	"github.com/redhatinsights/platform-changelog-go/internal/models"
	"github.com/redhatinsights/platform-changelog-go/internal/structs"
)

type DBConnector interface {
	Exec(string) error
	AutoMigrate(*models.Services, *models.Timelines) error

	CreateCommitEntry(timeline []models.Timelines) error
	GetCommitsAll(offset int, limit int) ([]models.Timelines, int64, error)
	GetCommitsByService(service structs.ServicesData, offset int, limit int) ([]models.Timelines, int64, error)
	GetCommitByRef(ref string) (models.Timelines, int64, error)

	CreateServiceTableEntry(name string, s config.Service) (models.Services, error)
	GetServicesAll(offset int, limit int) ([]structs.ExpandedServicesData, int64, error)
	GetLatest(service structs.ExpandedServicesData) (structs.ExpandedServicesData, error, error)
	GetServiceByName(name string) (structs.ServicesData, int64, error)
	GetServiceByGHRepo(repo string) (structs.ServicesData, error)

	GetTimelinesAll(offset int, limit int) ([]models.Timelines, int64, error)
	GetTimelinesByService(service structs.ServicesData, offset int, limit int) ([]models.Timelines, int64, error)
	GetTimelineByRef(ref string) (models.Timelines, int64, error)

	GetDeploysAll(offset int, limit int) ([]models.Timelines, int64, error)
	GetDeploysByService(service structs.ServicesData, offset int, limit int) ([]models.Timelines, int64, error)
	GetDeployByRef(ref string) (models.Timelines, int64, error)
}
