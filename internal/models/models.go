package models

import (
	"time"
)

type Services struct {
	ID          int         `gorm:"primary_key;autoincremement"`
	Name        string      `gorm:"not null"`
	DisplayName string      `gorm:"not null;unique"`
	Tenant      string      `gorm:"not null"`
	Projects    []Projects  `gorm:"foreignkey:ServiceID"`
	Timelines   []Timelines `gorm:"foreignkey:ServiceID"`
}

type Projects struct {
	ID         int    `gorm:"primary_key;autoincremement"`
	ServiceID  int    `gorm:"not null" json:"service_id"`
	Name       string `gorm:"not null"`
	Repo       string `gorm:"not null"`
	DeployFile string
	Namespaces []string    `gorm:"type:text[]"`
	Branches   []string    `gorm:"type:text[]"`
	Timelines  []Timelines `gorm:"foreignkey:ProjectID"`
}

type timelineType string

const (
	commit timelineType = "commit"
	deploy timelineType = "deploy"
)

type Timelines struct {
	ID              int          `gorm:"primary_key;autoincrement" json:"id"`
	ServiceID       int          `gorm:"not null" json:"service_id"`
	ProjectID       int          `gorm:"not null" json:"project_id"`
	Timestamp       time.Time    `gorm:"not null" json:"timestamp"`
	Type            timelineType `gorm:"not null" json:"type" sql:"type:timeline_type"`
	Repo            string       `gorm:"not null" json:"repo"`
	Ref             string       `json:"ref,omitempty"`
	Author          string       `json:"author,omitempty"`
	MergedBy        string       `json:"merged_by,omitempty"`
	Message         string       `json:"message,omitempty"`
	DeployNamespace string       `json:"namespace,omitempty"`
	Cluster         string       `json:"cluster,omitempty"`
	Image           string       `json:"image,omitempty"`
	TriggeredBy     string       `json:"triggered_by,omitempty"`
	Status          string       `json:"status,omitempty"`
}
