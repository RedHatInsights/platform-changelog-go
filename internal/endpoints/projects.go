package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	l "github.com/redhatinsights/platform-changelog-go/internal/logging"
	"github.com/redhatinsights/platform-changelog-go/internal/metrics"
	"github.com/redhatinsights/platform-changelog-go/internal/models"
	"github.com/redhatinsights/platform-changelog-go/internal/structs"
)

func (eh *EndpointHandler) GetProjectsAll(w http.ResponseWriter, r *http.Request) {
	metrics.IncRequests(r.URL.Path, r.Method, r.UserAgent())

	q, err := initQuery(r)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, "Invalid query")
		return
	}

	projects, count, err := eh.conn.GetProjectsAll(q.Offset, q.Limit, q)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	projectsList := structs.ProjectsList{Count: count, Data: convertProjectsToProjectsData(projects)}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(projectsList)
}

func (eh *EndpointHandler) GetProjectByName(w http.ResponseWriter, r *http.Request) {
	metrics.IncRequests(r.URL.Path, r.Method, r.UserAgent())

	projectName := chi.URLParam(r, "project")
	project, _, err := eh.conn.GetProjectByName(projectName)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	/**
	 * the project doesn't exist
	 */
	if project.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Service not found"))
		return
	}

	projectData := convertProjectToProjectsData(project)

	l.Log.Debugf("URL Param: %s", projectName)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(projectData)
}

func (eh *EndpointHandler) GetProjectsByService(w http.ResponseWriter, r *http.Request) {
	metrics.IncRequests(r.URL.Path, r.Method, r.UserAgent())

	serviceName := chi.URLParam(r, "service")

	q, err := initQuery(r)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, "Invalid query")
		return
	}

	service, _, err := eh.conn.GetServiceByName(serviceName)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Couldn't find the service"))
		return
	}

	projects, count, err := eh.conn.GetProjectsByService(service, q.Offset, q.Limit, q)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error producing the projects"))
		w.Write([]byte(err.Error()))
	}

	projectsList := structs.ProjectsList{Count: count, Data: convertProjectsToProjectsData(projects)}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(projectsList)
}

func convertProjectsToProjectsData(projects []models.Projects) []structs.ProjectsData {
	var projectsData []structs.ProjectsData

	for _, project := range projects {
		projectsData = append(projectsData, convertProjectToProjectsData(project))
	}

	return projectsData
}

func convertProjectToProjectsData(project models.Projects) structs.ProjectsData {
	projectData := structs.ProjectsData{
		ID:         project.ID,
		ServiceID:  project.ServiceID,
		Name:       project.Name,
		Repo:       project.Repo,
		DeployFile: project.DeployFile,
		Namespace:  project.Namespace,
		Branch:     project.Branch,
	}

	return projectData
}
