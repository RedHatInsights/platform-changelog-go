package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/redhatinsights/platform-changelog-go/internal/config"
	l "github.com/redhatinsights/platform-changelog-go/internal/logging"
	"github.com/redhatinsights/platform-changelog-go/internal/models"
	"github.com/redhatinsights/platform-changelog-go/internal/structs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnectorImpl struct {
	db *gorm.DB
}

func NewDBConnector(cfg *config.Config) *DBConnectorImpl {
	dsn, err := buildPostgresDSN(cfg)
	if err != nil {
		l.Log.Fatal("Error building postgres DSN: ", err)
		return nil
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		l.Log.Fatal(err)
	}

	l.Log.Info("DB initialization complete")

	return &DBConnectorImpl{db: db}
}

func OpenPostgresDB(cfg *config.Config) (*sql.DB, error) {
	dsn, err := buildPostgresDSN(cfg)
	if err != nil {
		l.Log.Fatal("Error building postgres DSN: ", err)
		return nil, err
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		l.Log.Fatal("Error opening DB: ", err)
		return nil, err
	}

	l.Log.Info("DB initialization complete")

	return db, nil
}

func buildPostgresDSN(cfg *config.Config) (string, error) {
	var (
		user     = cfg.DatabaseConfig.DBUser
		password = cfg.DatabaseConfig.DBPassword
		dbname   = cfg.DatabaseConfig.DBName
		host     = cfg.DatabaseConfig.DBHost
		port     = cfg.DatabaseConfig.DBPort
	)

	sslConfigString, err := buildPostgresSslConfigString(cfg)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s %s", user, password, dbname, host, port, sslConfigString), nil
}

func buildPostgresSslConfigString(cfg *config.Config) (string, error) {
	if cfg.DatabaseConfig.DBSSLMode == "disable" {
		return "sslmode=disable", nil
	} else if cfg.DatabaseConfig.DBSSLMode == "verify-full" {
		return "sslmode=verify-full sslrootcert=" + cfg.DatabaseConfig.RDSCa, nil
	} else {
		return "", errors.New("Invalid SSL configuration for database connection: " + cfg.DatabaseConfig.DBSSLMode)
	}
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
			service, _ := conn.CreateServiceTableEntry(key, service)
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

func (conn *MockDBConnector) GetCommitsAll(offset int, limit int, q structs.Query) ([]models.Timelines, int64, error) {
	commits := []models.Timelines{}
	for _, timeline := range conn.Timelines {
		if timeline.Type != "commit" || !filterCommit(timeline, q) {
			continue
		}

		commits = append(commits, timeline)
	}
	return commits, int64(len(commits)), nil
}

func (conn *MockDBConnector) GetCommitsByService(service structs.ServicesData, offset int, limit int, q structs.Query) ([]models.Timelines, int64, error) {
	commits := []models.Timelines{}
	for _, timeline := range conn.Timelines {
		if timeline.Type != "commit" || timeline.ServiceID != service.ID {
			continue
		}

		if !filterCommit(timeline, q) {
			continue
		}

		commits = append(commits, timeline)
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

func filterCommit(commit models.Timelines, q structs.Query) bool {
	if !filterByField(commit.Repo, q.Repo) {
		return false
	}

	if !filterByField(commit.Author, q.Author) {
		return false
	}

	if !filterByField(commit.MergedBy, q.MergedBy) {
		return false
	}

	if !filterByField(commit.Ref, q.Ref) {
		return false
	}

	return true
}

func (conn *MockDBConnector) CreateServiceTableEntry(name string, s config.Service) (models.Services, error) {
	newService := models.Services{
		ID:          len(conn.Services) + 1,
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

func (conn *MockDBConnector) UpdateServiceTableEntry(name string, s config.Service) (models.Services, error) {
	return models.Services{}, nil
}

func (conn *MockDBConnector) DeleteServiceTableEntry(name string) (structs.ServicesData, error) {
	return structs.ServicesData{}, nil
}

func (conn *MockDBConnector) GetServicesAll(offset int, limit int, q structs.Query) ([]structs.ExpandedServicesData, int64, error) {
	servicesWithTimelines := []structs.ExpandedServicesData{}

	for _, service := range conn.Services {
		if !filterService(service, q) {
			continue
		}
		serviceWithTimeline, _, _ := conn.GetLatest(structs.ExpandedServicesData{
			ServicesData: structs.ServicesData{
				ID:          service.ID,
				Name:        service.Name,
				DisplayName: service.DisplayName,
				GHRepo:      service.GHRepo,
				GLRepo:      service.GLRepo,
				DeployFile:  service.DeployFile,
				Namespace:   service.Namespace,
				Branch:      service.Branch,
			},
		})
		servicesWithTimelines = append(servicesWithTimelines, serviceWithTimeline)
	}

	return servicesWithTimelines, int64(len(servicesWithTimelines)), nil
}

func (conn *MockDBConnector) GetLatest(service structs.ExpandedServicesData) (structs.ExpandedServicesData, error, error) {
	expandedService := structs.ExpandedServicesData{
		ServicesData: service.ServicesData,
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

func (conn *MockDBConnector) GetServiceByName(name string) (structs.ServicesData, int64, error) {
	for _, service := range conn.Services {
		if service.Name == name {
			return structs.ServicesData{
				ID:          service.ID,
				Name:        service.Name,
				DisplayName: service.DisplayName,
				GHRepo:      service.GHRepo,
				GLRepo:      service.GLRepo,
				DeployFile:  service.DeployFile,
				Namespace:   service.Namespace,
				Branch:      service.Branch,
			}, 1, nil
		}
	}
	return structs.ServicesData{}, 0, nil
}

func (conn *MockDBConnector) GetServiceByGHRepo(repo string) (structs.ServicesData, error) {
	for _, service := range conn.Services {
		if service.GHRepo == repo {
			return structs.ServicesData{
				ID:          service.ID,
				Name:        service.Name,
				DisplayName: service.DisplayName,
				GHRepo:      service.GHRepo,
				GLRepo:      service.GLRepo,
				DeployFile:  service.DeployFile,
				Namespace:   service.Namespace,
				Branch:      service.Branch,
			}, nil
		}
	}
	return structs.ServicesData{}, nil
}

func filterService(service models.Services, q structs.Query) bool {
	if !filterByField(service.Name, q.ServiceName) {
		return false
	}

	if !filterByField(service.DisplayName, q.ServiceDisplayName) {
		return false
	}

	if !filterByField(service.Tenant, q.ServiceTenant) {
		return false
	}

	if !filterByField(service.Branch, q.ServiceBranch) {
		return false
	}

	if !filterByField(service.Namespace, q.ServiceNamespace) {
		return false
	}

	return true
}

func (conn *MockDBConnector) GetTimelinesAll(offset int, limit int, q structs.Query) ([]models.Timelines, int64, error) {
	var timelines []models.Timelines
	for _, timeline := range conn.Timelines {
		if !filterTimeline(timeline, q) {
			continue
		}

		timelines = append(timelines, timeline)
	}
	return timelines, int64(len(timelines)), nil
}

func (conn *MockDBConnector) GetTimelinesByService(service structs.ServicesData, offset int, limit int, q structs.Query) ([]models.Timelines, int64, error) {
	var timelines []models.Timelines
	for _, timeline := range conn.Timelines {
		if timeline.ServiceID != service.ID {
			continue
		}

		if !filterTimeline(timeline, q) {
			continue
		}

		timelines = append(timelines, timeline)
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

func filterTimeline(timeline models.Timelines, q structs.Query) bool {
	if !filterByField(timeline.Repo, q.Repo) {
		return false
	}

	if !filterByField(timeline.Ref, q.Ref) {
		return false
	}

	return true
}

func (conn *MockDBConnector) DeleteTimelinesByService(services structs.ServicesData) error {
	// no need to implement now
	return nil
}

func (conn *MockDBConnector) CreateDeployEntry(timeline models.Timelines) error {
	conn.Timelines = append(conn.Timelines, timeline)
	return nil
}

func (conn *MockDBConnector) GetDeploysAll(offset int, limit int, q structs.Query) ([]models.Timelines, int64, error) {
	deploys := []models.Timelines{}
	for _, timeline := range conn.Timelines {
		if timeline.Type != "deploy" {
			continue
		}

		if !filterDeploy(timeline, q) {
			continue
		}

		deploys = append(deploys, timeline)
	}
	return deploys, int64(len(deploys)), nil
}

func (conn *MockDBConnector) GetDeploysByService(service structs.ServicesData, offset int, limit int, q structs.Query) ([]models.Timelines, int64, error) {
	deploys := []models.Timelines{}
	for _, timeline := range conn.Timelines {
		if timeline.Type != "deploy" || timeline.ServiceID != service.ID {
			continue
		}

		if !filterDeploy(timeline, q) {
			continue
		}

		deploys = append(deploys, timeline)
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

func filterDeploy(deploy models.Timelines, q structs.Query) bool {
	if !filterByField(deploy.Repo, q.Repo) {
		return false
	}

	if !filterByField(deploy.Ref, q.Ref) {
		return false
	}

	if !filterByField(deploy.Cluster, q.Cluster) {
		return false
	}

	if !filterByField(deploy.Image, q.Image) {
		return false
	}

	return true
}
