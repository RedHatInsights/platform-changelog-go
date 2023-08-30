package structs

import (
	"github.com/redhatinsights/platform-changelog-go/internal/models"
)

type Query struct {
	Offset      int
	Limit       int
	Ref         []string
	Repo        []string
	Author      []string
	MergedBy    []string
	Cluster     []string
	Image       []string
	Name        []string // service and project filters
	DisplayName []string
	Tenant      []string
	Namespace   []string
	Branch      []string
	StartDate   string
	EndDate     string
}

// Add Link object to these structs for more clear pagination
// https://jsonapi.org/format/#fetching-pagination

// That would include adding a middleware and changing all these List structs
// to be covered by one ResponseData struct
type ServicesList struct {
	Count int64          `json:"count"`
	Data  []ServicesData `json:"data"`
}

type ExpandedServicesList struct {
	Count int64                  `json:"count"`
	Data  []ExpandedServicesData `json:"data"`
}

type TimelinesList struct {
	Count int64              `json:"count"`
	Data  []models.Timelines `json:"data"`
}

type ProjectsList struct {
	Count int64          `json:"count"`
	Data  []ProjectsData `json:"data"`
}

type ProjectsData struct {
	ID         int    `json:"id"`
	ServiceID  int    `json:"service_id"`
	Name       string `json:"name"`
	Repo       string `json:"repo"`
	DeployFile string `json:"deploy_file"`
	Namespace  string `json:"namespace"`
	Branch     string `json:"branch"`
}

type ServicesData struct {
	ID          int            `json:"id"`
	Name        string         `json:"name"`
	DisplayName string         `json:"display_name"`
	Tenant      string         `json:"tenant"`
	Projects    []ProjectsData `json:"projects"`
}

type ExpandedServicesData struct {
	ID          int              `json:"id"`
	Name        string           `json:"name"`
	DisplayName string           `json:"display_name"`
	Tenant      string           `json:"tenant"`
	Commit      models.Timelines `json:"latest_commit" gorm:"foreignkey:ID"`
	Deploy      models.Timelines `json:"latest_deploy" gorm:"foreignkey:ID"`
}
