package db

import (
	"fmt"

	"github.com/redhatinsights/platform-changelog-go/internal/config"
	l "github.com/redhatinsights/platform-changelog-go/internal/logging"
	"github.com/redhatinsights/platform-changelog-go/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnector interface {
	CreateCommitEntry(timeline []models.Timelines) error
	GetCommitsAll(offset int, limit int) ([]models.Timelines, int64, error)
	GetCommitsByService(service models.Services, offset int, limit int) ([]models.Timelines, int64, error)
	GetCommitByRef(ref string) (models.Timelines, int64, error)

	CreateServiceTableEntry(name string, s config.Service) (models.Services, error)
	GetServicesAll(offset int, limit int) ([]models.ExpandedServices, int64, error)
	GetLatest(service models.ExpandedServices) (models.ExpandedServices, error, error)
	GetServiceByName(name string) (models.Services, int64, error)
	GetServiceByGHRepo(repo string) (models.Services, error)

	GetTimelinesAll(offset int, limit int) ([]models.Timelines, int64, error)
	GetTimelinesByService(service models.Services, offset int, limit int) ([]models.Timelines, int64, error)
	GetTimelineByRef(ref string) (models.Timelines, int64, error)

	GetDeploysAll(offset int, limit int) ([]models.Timelines, int64, error)
	GetDeploysByService(service models.Services, offset int, limit int) ([]models.Timelines, int64, error)
	GetDeployByRef(ref string) (models.Timelines, int64, error)
}

type DBConnectorImpl struct {
	db *gorm.DB
}

func NewDBConnector(cfg *config.Config) DBConnector {
	var (
		user     = cfg.DatabaseConfig.DBUser
		password = cfg.DatabaseConfig.DBPassword
		dbname   = cfg.DatabaseConfig.DBName
		host     = cfg.DatabaseConfig.DBHost
		port     = cfg.DatabaseConfig.DBPort
	)
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, password, dbname, host, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		l.Log.Fatal(err)
	}

	l.Log.Info("DB initialization complete")

	return &DBConnectorImpl{db: db}
}
