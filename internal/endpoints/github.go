package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/redhatinsights/platform-changelog-go/internal/db"
	l "github.com/redhatinsights/platform-changelog-go/internal/logging"
	"github.com/redhatinsights/platform-changelog-go/internal/metrics"
	"github.com/redhatinsights/platform-changelog-go/internal/models"
	"github.com/redhatinsights/platform-changelog-go/internal/structs"
)

// This endpoint is different than the github and gitlab endpoints
// This will be used as a part of the Jenkins pipeline
// on each push to a monitored branch (configured in app-interface)

type GithubPayload *struct {
	App     string `json:"app"`
	Project string `json:"project"`
	Tenant  string `json:"tenant"`
	Repo    string `json:"repo,omitempty"`
	Branch  string `json:"branch"`
	Ref     string `json:"ref"`
}

func decodeGithubJSONBody(w http.ResponseWriter, r *http.Request) (GithubPayload, error) {
	if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		return nil, fmt.Errorf("invalid Content-Type header: '%s' should contain 'application/json'", r.Header.Get("Content-Type"))
	}

	if r.Body == nil {
		return nil, fmt.Errorf("json body required")
	}

	var payload GithubPayload

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if !dec.More() {
		return nil, fmt.Errorf("empty json body provided")
	}

	err := dec.Decode(&payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func (eh *EndpointHandler) Github(w http.ResponseWriter, r *http.Request) {
	metrics.IncJenkins("github", r.Method, r.UserAgent(), false)

	// log everything for now
	l.Log.Info("Github Jenkins run received")
	l.Log.Info(r.Body)

	payload, err := decodeGithubJSONBody(w, r)
	if err != nil {
		l.Log.Error(err)
		metrics.IncJenkins("github", r.Method, r.UserAgent(), true)
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	err = validateGithubPayload(payload)

	if err != nil {
		l.Log.Error(err)
		metrics.IncJenkins("github", r.Method, r.UserAgent(), true)
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	service, project, err1, err2 := getServiceAndProject(eh.conn, payload)
	if err1 != nil || err2 != nil {
		if service.ID == 0 { // how do I compare the structs completely?
			service, err = createNewService(eh.conn, payload)
			if err != nil {
				l.Log.Error(err)
				metrics.IncJenkins("github", r.Method, r.UserAgent(), true)
				writeResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
		}
		if project.ID == 0 {
			project, err = createNewProject(eh.conn, payload, service)
			if err != nil {
				l.Log.Error(err)
				metrics.IncJenkins("github", r.Method, r.UserAgent(), true)
				writeResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
		}
	}

	commit, err := convertGithubPayloadToTimelines(eh.conn, payload, service, project)
	if err != nil {
		l.Log.Error(err)
		metrics.IncJenkins("github", r.Method, r.UserAgent(), true)
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = eh.conn.CreateCommitEntry(commit)

	if err != nil {
		l.Log.Error(err)
		metrics.IncJenkins("github", r.Method, r.UserAgent(), true)
		writeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeResponse(w, http.StatusOK, `{"msg": "Commit info received"}`)
}

// Validate the payload contains necessary data
func validateGithubPayload(payload GithubPayload) error {
	// timestamp no longer required since we will be getting it from the github api soon

	if payload.App == "" {
		return fmt.Errorf("app is required")
	}

	if payload.Project == "" {
		return fmt.Errorf("project is required")
	}

	if payload.Tenant == "" {
		return fmt.Errorf("tenant is required")
	}

	if payload.Repo == "" {
		return fmt.Errorf("repo is required")
	}

	if payload.Branch == "" {
		return fmt.Errorf("branch is required")
	}

	if payload.Ref == "" {
		return fmt.Errorf("ref is required")
	}

	return nil
}

func getServiceAndProject(conn db.DBConnector, payload GithubPayload) (service structs.ServicesData, project structs.ProjectsData, err1 error, err2 error) {
	service, _, err1 = conn.GetServiceByName(payload.App)
	project, _, err2 = conn.GetProjectByName(payload.Project)

	return
}

func createNewService(conn db.DBConnector, payload GithubPayload) (service structs.ServicesData, err error) {
	// couldn't find service; create it, then handle the project
	s := models.Services{
		Name:        payload.App,
		DisplayName: payload.App,
		Tenant:      payload.Tenant,
	}

	_, err = conn.CreateServiceTableEntry(s)
	if err != nil {
		return structs.ServicesData{}, fmt.Errorf("problem creating service %s", payload.App)
	}

	service, _, err = conn.GetServiceByName(payload.App)
	return
}

func createNewProject(conn db.DBConnector, payload GithubPayload, service structs.ServicesData) (project structs.ProjectsData, err error) {
	p := models.Projects{
		ServiceID: service.ID,
		Name:      payload.Project,
		Repo:      payload.Repo,
		Branch:    payload.Branch,
	}

	err = conn.CreateProjectTableEntry(p)
	if err != nil {
		return structs.ProjectsData{}, fmt.Errorf("problem creating project %s", payload.Project)
	}

	project, _, err = conn.GetProjectByName(payload.Project)
	return
}

// Converting from GithubPayload struct to Timeline model
func convertGithubPayloadToTimelines(conn db.DBConnector, payload GithubPayload, service structs.ServicesData, project structs.ProjectsData) (commit models.Timelines, err error) {
	// author, timestamp, mergedby, and message will be updated with information from github api

	t := models.Timelines{
		ServiceID: service.ID,
		ProjectID: project.ID,
		Timestamp: time.Now(),
		Type:      "commit",
		Repo:      payload.Repo,
		Ref:       payload.Ref,
	}

	return t, nil
}
