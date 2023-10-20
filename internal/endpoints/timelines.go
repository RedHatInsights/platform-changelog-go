package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/redhatinsights/platform-changelog-go/internal/metrics"
	"github.com/redhatinsights/platform-changelog-go/internal/structs"
)

func (eh *EndpointHandler) GetTimelinesAll(w http.ResponseWriter, r *http.Request) {
	metrics.IncRequests(r.URL.Path, r.Method, r.UserAgent())

	q, err := initQuery(r)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, "Invalid query")
		return
	}

	timeline, count, err := eh.conn.GetTimelinesAll(q.Offset, q.Limit, q)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error producing the timeline"))
		w.Write([]byte(err.Error()))
		return
	}

	timelinesList := structs.TimelinesList{Count: count, Data: timeline}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(timelinesList)
}

func (eh *EndpointHandler) GetTimelinesByService(w http.ResponseWriter, r *http.Request) {
	metrics.IncRequests(r.URL.Path, r.Method, r.UserAgent())

	serviceID, err := getIDFromURL(r, "service_id")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid service ID"))
		return
	}

	q, err := initQuery(r)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, "Invalid query")
		return
	}

	service, _, err := eh.conn.GetServiceByID(serviceID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Couldn't find the service"))
		return
	}

	timeline, count, err := eh.conn.GetTimelinesByService(service, q.Offset, q.Limit, q)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error producing the timeline"))
		w.Write([]byte(err.Error()))
	}

	timelinesList := structs.TimelinesList{Count: count, Data: timeline}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(timelinesList)
}

func (eh *EndpointHandler) GetTimelinesByProject(w http.ResponseWriter, r *http.Request) {
	metrics.IncRequests(r.URL.Path, r.Method, r.UserAgent())

	projectID, err := getIDFromURL(r, "project_id")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid project ID"))
		return
	}

	q, err := initQuery(r)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, "Invalid query")
		return
	}

	project, _, err := eh.conn.GetProjectByID(projectID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Couldn't find the project"))
		return
	}

	timeline, count, err := eh.conn.GetTimelinesByProject(project, q.Offset, q.Limit, q)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error producing the timeline"))
		w.Write([]byte(err.Error()))
	}

	timelinesList := structs.TimelinesList{Count: count, Data: timeline}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(timelinesList)
}

func (eh *EndpointHandler) GetTimelineByRef(w http.ResponseWriter, r *http.Request) {
	metrics.IncRequests(r.URL.Path, r.Method, r.UserAgent())
	ref := chi.URLParam(r, "ref")

	timeline, rowsAffected, err := eh.conn.GetTimelineByRef(ref)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error producing the timeline"))
		w.Write([]byte(err.Error()))
		return
	}

	if rowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Timeline not found"))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(timeline)
}
