package db

import (
	"context"

	"github.com/redhatinsights/platform-changelog-go/internal/structs"
)

type ServiceRepository interface {
	GetServicesAll(ctx context.Context, offset int, limit int) ([]structs.ServicesData, error)
	GetServiceByName(ctx context.Context, name string) (structs.ServicesData, error)
	GetLatest(ctx context.Context, service structs.ServicesData) (structs.ExpandedServicesData, error)
	GetServiceByGHRepo(ctx context.Context, repo string) (structs.ServicesData, error)
	CreateServiceEntry(ctx context.Context, service structs.ServicesData) error
}

type MockServiceRepository struct {
	services *[]structs.ServicesData
}

func (m *MockServiceRepository) GetServicesAll(ctx context.Context, offset int, limit int) ([]structs.ServicesData, error) {
	return *m.services, nil
}

func (m *MockServiceRepository) GetServiceByName(ctx context.Context, name string) (structs.ServicesData, error) {
	for _, service := range *m.services {
		if service.Name == name {
			return service, nil
		}
	}
	return structs.ServicesData{}, nil
}

func (m *MockServiceRepository) GetLatest(ctx context.Context, service structs.ServicesData) (structs.ExpandedServicesData, error) {
	expandedService := structs.ExpandedServicesData{
		ServicesData: service,
	}

	// TODO: get latest commit and deploy

	return expandedService, nil
}

func (m *MockServiceRepository) GetServiceByGHRepo(ctx context.Context, repo string) (structs.ServicesData, error) {
	for _, service := range *m.services {
		if service.GHRepo == repo {
			return service, nil
		}
	}
	return structs.ServicesData{}, nil
}

func (m *MockServiceRepository) CreateServiceEntry(ctx context.Context, service structs.ServicesData) error {
	*m.services = append(*m.services, service)
	return nil
}
