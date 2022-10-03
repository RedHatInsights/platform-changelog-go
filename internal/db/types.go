package db

import (
	"context"

	"github.com/redhatinsights/platform-changelog-go/internal/structs"
)

type DBConnector interface {
	GetCommitsAll(ctx context.Context) ([]structs.TimelinesData, error)
	GetCommitsByService(ctx context.Context, service structs.ServicesData) ([]structs.TimelinesData, error)
	GetCommitByRef(ctx context.Context, ref string) (structs.TimelinesData, error)
	CreateCommitEntry(ctx context.Context, commitData []structs.TimelinesData) error

	GetServicesAll(ctx context.Context) ([]structs.ServicesData, error)
	GetServiceByName(ctx context.Context, name string) (structs.ServicesData, error)
	GetLatest(ctx context.Context, service structs.ServicesData) (structs.ExpandedServicesData, error)
	GetServiceByGHRepo(ctx context.Context, repo string) (structs.ServicesData, error)
	CreateServiceEntry(ctx context.Context, service structs.ServicesData) error

	GetTimelinesAll(ctx context.Context) ([]structs.TimelinesData, error)
	GetTimelinesByService(ctx context.Context, service structs.ServicesData) ([]structs.TimelinesData, error)
	GetTimelineByRef(ctx context.Context, ref string) (structs.TimelinesData, error)

	GetDeploysAll(ctx context.Context) ([]structs.TimelinesData, error)
	GetDeploysByService(ctx context.Context, service structs.ServicesData) ([]structs.TimelinesData, error)
	GetDeployByRef(ctx context.Context, ref string) (structs.TimelinesData, error)
}

type MockDBConnector struct {
	timelines []structs.TimelinesData
	services  []structs.ServicesData
}

func (m *MockDBConnector) GetCommitsAll(ctx context.Context) ([]structs.TimelinesData, error) {
	commits := []structs.TimelinesData{}
	for _, timeline := range m.timelines {
		if timeline.Type == "commit" {
			commits = append(commits, timeline)
		}
	}
	return commits, nil
}

func (m *MockDBConnector) GetCommitsByService(ctx context.Context, service structs.ServicesData) ([]structs.TimelinesData, error) {
	commits := []structs.TimelinesData{}
	for _, timeline := range m.timelines {
		if timeline.Type == "commit" && timeline.ServiceID == service.ID {
			commits = append(commits, timeline)
		}
	}
	return commits, nil
}

func (m *MockDBConnector) GetCommitByRef(ctx context.Context, ref string) (structs.TimelinesData, error) {
	for _, timeline := range m.timelines {
		if timeline.Type == "commit" && timeline.Ref == ref {
			return timeline, nil
		}
	}
	return structs.TimelinesData{}, nil
}

func (m *MockDBConnector) CreateCommitEntry(ctx context.Context, commitData []structs.TimelinesData) error {
	m.timelines = append(m.timelines, commitData...)
	return nil
}

func (m *MockDBConnector) GetServicesAll(ctx context.Context) ([]structs.ServicesData, error) {
	return m.services, nil
}

func (m *MockDBConnector) GetServiceByName(ctx context.Context, name string) (structs.ServicesData, error) {
	for _, service := range m.services {
		if service.Name == name {
			return service, nil
		}
	}
	return structs.ServicesData{}, nil
}

func (m *MockDBConnector) GetLatest(ctx context.Context, service structs.ServicesData) (structs.ExpandedServicesData, error) {
	expandedService := structs.ExpandedServicesData{
		ServicesData: service,
	}

	// TODO: get latest commit and deploy

	return expandedService, nil
}

func (m *MockDBConnector) GetServiceByGHRepo(ctx context.Context, repo string) (structs.ServicesData, error) {
	for _, service := range m.services {
		if service.GHRepo == repo {
			return service, nil
		}
	}
	return structs.ServicesData{}, nil
}

func (m *MockDBConnector) CreateServiceEntry(ctx context.Context, service structs.ServicesData) error {
	m.services = append(m.services, service)
	return nil
}

func (m *MockDBConnector) GetTimelinesAll(ctx context.Context) ([]structs.TimelinesData, error) {
	return m.timelines, nil
}

func (m *MockDBConnector) GetTimelinesByService(ctx context.Context, service structs.ServicesData) ([]structs.TimelinesData, error) {
	timelines := []structs.TimelinesData{}
	for _, timeline := range m.timelines {
		if timeline.ServiceID == service.ID {
			timelines = append(timelines, timeline)
		}
	}
	return timelines, nil
}

func (m *MockDBConnector) GetTimelineByRef(ctx context.Context, ref string) (structs.TimelinesData, error) {
	for _, timeline := range m.timelines {
		if timeline.Ref == ref {
			return timeline, nil
		}
	}
	return structs.TimelinesData{}, nil
}

func (m *MockDBConnector) GetDeploysAll(ctx context.Context) ([]structs.TimelinesData, error) {
	deploys := []structs.TimelinesData{}
	for _, timeline := range m.timelines {
		if timeline.Type == "deploy" {
			deploys = append(deploys, timeline)
		}
	}
	return deploys, nil
}

func (m *MockDBConnector) GetDeploysByService(ctx context.Context, service structs.ServicesData) ([]structs.TimelinesData, error) {
	deploys := []structs.TimelinesData{}
	for _, timeline := range m.timelines {
		if timeline.Type == "deploy" && timeline.ServiceID == service.ID {
			deploys = append(deploys, timeline)
		}
	}
	return deploys, nil
}

func (m *MockDBConnector) GetDeployByRef(ctx context.Context, ref string) (structs.TimelinesData, error) {
	for _, timeline := range m.timelines {
		if timeline.Type == "deploy" && timeline.Ref == ref {
			return timeline, nil
		}
	}
	return structs.TimelinesData{}, nil
}
