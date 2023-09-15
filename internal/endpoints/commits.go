package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/redhatinsights/platform-changelog-go/internal/structs"

	"github.com/go-chi/chi/v5"
	"github.com/redhatinsights/platform-changelog-go/internal/metrics"
)

func (eh *EndpointHandler) GetCommitsAll(w http.ResponseWriter, r *http.Request) {
	metrics.IncRequests(r.URL.Path, r.Method, r.UserAgent())

	q, err := initQuery(r)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, "Invalid query")
		return
	}

	commits, count, err := eh.conn.GetCommitsAll(q.Offset, q.Limit, q)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	commitsList := structs.TimelinesList{Count: count, Data: commits}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(commitsList)
}

func (eh *EndpointHandler) GetCommitsByService(w http.ResponseWriter, r *http.Request) {
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

	commits, count, err := eh.conn.GetCommitsByService(service, q.Offset, q.Limit, q)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	commitsList := structs.TimelinesList{Count: count, Data: commits}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(commitsList)
}

func (eh *EndpointHandler) GetCommitsByProject(w http.ResponseWriter, r *http.Request) {
	metrics.IncRequests(r.URL.Path, r.Method, r.UserAgent())
	projectName := chi.URLParam(r, "project")

	q, err := initQuery(r)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, "Invalid query")
		return
	}

	project, _, err := eh.conn.GetProjectByName(projectName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Couldn't find the service"))
		return
	}

	commits, count, err := eh.conn.GetCommitsByProject(project, q.Offset, q.Limit, q)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	commitsList := structs.TimelinesList{Count: count, Data: commits}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(commitsList)
}

func (eh *EndpointHandler) GetCommitByRef(w http.ResponseWriter, r *http.Request) {
	metrics.IncRequests(r.URL.Path, r.Method, r.UserAgent())
	ref := chi.URLParam(r, "ref")

	commit, rowsAffected, err := eh.conn.GetCommitByRef(ref)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	if rowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Commit not found"))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(commit)
}
