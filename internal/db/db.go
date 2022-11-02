package db

import (
	"fmt"

	"github.com/redhatinsights/platform-changelog-go/internal/config"
	l "github.com/redhatinsights/platform-changelog-go/internal/logging"
	"github.com/redhatinsights/platform-changelog-go/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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

func (conn *DBConnectorImpl) AutoMigrate(serviceModel *models.Services, timelineModel *models.Timelines) error {
	return conn.db.AutoMigrate(serviceModel, timelineModel)
}

func (conn *DBConnectorImpl) Exec(sql string) error {
	return conn.db.Exec(sql).Error
}

type MockDBConnector struct {
	Timelines []models.Timelines
	Services  []models.Services
}

func NewMockDBConnector(cfg *config.Config) DBConnector {
	fmt.Println("Using MockDBConnector")

	conn := &MockDBConnector{
		Timelines: []models.Timelines{},
		Services:  []models.Services{},
	}

	// Add the data from the config file to the mock DB
	for key, service := range cfg.Services {
		_, rowsAffected, _ := conn.GetServiceByName(key)
		if rowsAffected == 0 {
			_, service := conn.CreateServiceTableEntry(key, service)
			l.Log.Info("Created service: ", service)
		} else {
			l.Log.Info("Service already exists: ", service.DisplayName)
		}
	}

	return conn
}

func (conn *MockDBConnector) AutoMigrate(serviceModel *models.Services, timelineModel *models.Timelines) error {
	return nil
}

func (conn *MockDBConnector) Exec(sql string) error {
	return nil
}

func (conn *MockDBConnector) CreateCommitEntry(timeline []models.Timelines) error {
	conn.Timelines = append(conn.Timelines, timeline...)
	return nil
}

func (conn *MockDBConnector) GetCommitsAll(offset int, limit int) ([]models.Timelines, int64, error) {
	var commits []models.Timelines
	for _, timeline := range conn.Timelines {
		if timeline.Type == "commit" {
			commits = append(commits, timeline)
		}
	}
	return commits, int64(len(commits)), nil
}

func (conn *MockDBConnector) GetCommitsByService(service models.Services, offset int, limit int) ([]models.Timelines, int64, error) {
	var commits []models.Timelines
	for _, timeline := range conn.Timelines {
		if timeline.Type == "commit" && timeline.ServiceID == service.ID {
			commits = append(commits, timeline)
		}
	}
	return commits, int64(len(commits)), nil
}

func (conn *MockDBConnector) GetCommitByRef(ref string) (models.Timelines, int64, error) {
	for _, timeline := range conn.Timelines {
		if timeline.Ref == ref {
			return timeline, 1, nil
		}
	}
	return models.Timelines{}, 0, nil
}

func (conn *MockDBConnector) CreateServiceTableEntry(name string, s config.Service) (models.Services, error) {
	newService := models.Services{
		Name:        name,
		DisplayName: s.DisplayName,
		GHRepo:      s.GHRepo,
		GLRepo:      s.GLRepo, Branch: s.Branch,
		Namespace:  s.Namespace,
		DeployFile: s.DeployFile,
	}

	conn.Services = append(conn.Services, newService)
	return newService, nil
}

func (conn *MockDBConnector) GetServicesAll(offset int, limit int) ([]models.ExpandedServices, int64, error) {
	servicesWithTimelines := []models.ExpandedServices{}

	for _, service := range conn.Services {
		serviceWithTimeline, _, _ := conn.GetLatest(models.ExpandedServices{
			Services: service,
		})
		servicesWithTimelines = append(servicesWithTimelines, serviceWithTimeline)
	}

	return servicesWithTimelines, int64(len(servicesWithTimelines)), nil
}

func (conn *MockDBConnector) GetLatest(service models.ExpandedServices) (models.ExpandedServices, error, error) {
	expandedService := models.ExpandedServices{
		Services: service.Services,
	}

	for _, timeline := range conn.Timelines {
		if timeline.ServiceID == service.ID && timeline.Type == "commit" {
			expandedService.Commit = timeline
		}
		if timeline.ServiceID == service.ID && timeline.Type == "deploy" {
			expandedService.Deploy = timeline
		}
	}

	return expandedService, nil, nil
}

func (conn *MockDBConnector) GetServiceByName(name string) (models.Services, int64, error) {
	for _, service := range conn.Services {
		if service.Name == name {
			return service, 1, nil
		}
	}
	return models.Services{}, 0, nil
}

func (conn *MockDBConnector) GetServiceByGHRepo(repo string) (models.Services, error) {
	for _, service := range conn.Services {
		if service.GHRepo == repo {
			return service, nil
		}
	}
	return models.Services{}, nil
}

func (conn *MockDBConnector) GetTimelinesAll(offset int, limit int) ([]models.Timelines, int64, error) {
	return conn.Timelines, int64(len(conn.Timelines)), nil
}

func (conn *MockDBConnector) GetTimelinesByService(service models.Services, offset int, limit int) ([]models.Timelines, int64, error) {
	var timelines []models.Timelines
	for _, timeline := range conn.Timelines {
		if timeline.ServiceID == service.ID {
			timelines = append(timelines, timeline)
		}
	}
	return timelines, int64(len(timelines)), nil
}

func (conn *MockDBConnector) GetTimelineByRef(ref string) (models.Timelines, int64, error) {
	for _, timeline := range conn.Timelines {
		if timeline.Ref == ref {
			return timeline, 1, nil
		}
	}
	return models.Timelines{}, 0, nil
}

func (conn *MockDBConnector) GetDeploysAll(offset int, limit int) ([]models.Timelines, int64, error) {
	deploys := []models.Timelines{}
	for _, timeline := range conn.Timelines {
		if timeline.Type == "deploy" {
			deploys = append(deploys, timeline)
		}
	}
	return deploys, int64(len(deploys)), nil
}

func (conn *MockDBConnector) GetDeploysByService(service models.Services, offset int, limit int) ([]models.Timelines, int64, error) {
	deploys := []models.Timelines{}
	for _, timeline := range conn.Timelines {
		if timeline.Type == "deploy" && timeline.ServiceID == service.ID {
			deploys = append(deploys, timeline)
		}
	}
	return deploys, int64(len(deploys)), nil
}

func (conn *MockDBConnector) GetDeployByRef(ref string) (models.Timelines, int64, error) {
	for _, timeline := range conn.Timelines {
		if timeline.Ref == ref {
			return timeline, 1, nil
		}
	}
	return models.Timelines{}, 0, nil
}
