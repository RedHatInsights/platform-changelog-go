package db

import (
	"context"

	"github.com/redhatinsights/platform-changelog-go/internal/structs"
)

type TimelineRepository interface {
	GetTimelinesAll(ctx context.Context) ([]structs.TimelinesData, error)
	GetTimelinesByService(ctx context.Context, service structs.ServicesData) ([]structs.TimelinesData, error)
	GetTimelineByRef(ctx context.Context, ref string) (structs.TimelinesData, error)
}

type MockTimelineRepository struct {
	timelines *[]structs.TimelinesData
}

func (m *MockTimelineRepository) GetTimelinesAll(ctx context.Context) ([]structs.TimelinesData, error) {
	return *m.timelines, nil
}

func (m *MockTimelineRepository) GetTimelinesByService(ctx context.Context, service structs.ServicesData) ([]structs.TimelinesData, error) {
	timelines := []structs.TimelinesData{}
	for _, timeline := range *m.timelines {
		if timeline.ServiceID == service.ID {
			timelines = append(timelines, timeline)
		}
	}
	return timelines, nil
}

func (m *MockTimelineRepository) GetTimelineByRef(ctx context.Context, ref string) (structs.TimelinesData, error) {
	for _, timeline := range *m.timelines {
		if timeline.Ref == ref {
			return timeline, nil
		}
	}
	return structs.TimelinesData{}, nil
}
