package db

import (
	"context"

	"github.com/redhatinsights/platform-changelog-go/internal/structs"
)

type CommitRepository interface {
	GetCommitsAll(ctx context.Context) ([]structs.TimelinesData, error)
	GetCommitsByService(ctx context.Context, service structs.ServicesData) ([]structs.TimelinesData, error)
	GetCommitByRef(ctx context.Context, ref string) (structs.TimelinesData, error)
	CreateCommitEntry(ctx context.Context, commitData []structs.TimelinesData) error
}

type MockCommitRepository struct {
	timelines *[]structs.TimelinesData
}

func (m *MockCommitRepository) GetCommitsAll(ctx context.Context) ([]structs.TimelinesData, error) {
	commits := []structs.TimelinesData{}
	for _, timeline := range *m.timelines {
		if timeline.Type == "commit" {
			commits = append(commits, timeline)
		}
	}
	return commits, nil
}

func (m *MockCommitRepository) GetCommitsByService(ctx context.Context, service structs.ServicesData) ([]structs.TimelinesData, error) {
	commits := []structs.TimelinesData{}
	for _, timeline := range *m.timelines {
		if timeline.Type == "commit" && timeline.ServiceID == service.ID {
			commits = append(commits, timeline)
		}
	}
	return commits, nil
}

func (m *MockCommitRepository) GetCommitByRef(ctx context.Context, ref string) (structs.TimelinesData, error) {
	for _, timeline := range *m.timelines {
		if timeline.Type == "commit" && timeline.Ref == ref {
			return timeline, nil
		}
	}
	return structs.TimelinesData{}, nil
}

func (m *MockCommitRepository) CreateCommitEntry(ctx context.Context, commitData []structs.TimelinesData) error {
	*m.timelines = append(*m.timelines, commitData...)
	return nil
}
