package db

import (
	"context"

	"github.com/redhatinsights/platform-changelog-go/internal/structs"
)

type DeployRepository interface {
	GetDeploysAll(ctx context.Context, offset int, limit int) ([]structs.TimelinesData, error)
	GetDeploysByService(ctx context.Context, service structs.ServicesData, offset int, limit int) ([]structs.TimelinesData, error)
	GetDeployByRef(ctx context.Context, ref string) (structs.TimelinesData, error)
}

type MockDeployRepository struct {
	timelines *[]structs.TimelinesData
}

func (m *MockDeployRepository) GetDeploysAll(ctx context.Context, offset int, limit int) ([]structs.TimelinesData, error) {
	deploys := []structs.TimelinesData{}
	for _, timeline := range *m.timelines {
		if timeline.Type == "deploy" {
			deploys = append(deploys, timeline)
		}
	}

	return deploys, nil
}

func (m *MockDeployRepository) GetDeploysByService(ctx context.Context, service structs.ServicesData, offset int, limit int) ([]structs.TimelinesData, error) {
	deploys := []structs.TimelinesData{}
	for _, timeline := range *m.timelines {
		if timeline.Type == "deploy" && timeline.ServiceID == service.ID {
			deploys = append(deploys, timeline)
		}
	}
	return deploys, nil
}

func (m *MockDeployRepository) GetDeployByRef(ctx context.Context, ref string) (structs.TimelinesData, error) {
	for _, timeline := range *m.timelines {
		if timeline.Type == "deploy" && timeline.Ref == ref {
			return timeline, nil
		}
	}
	return structs.TimelinesData{}, nil
}
