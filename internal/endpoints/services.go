package endpoints

import (
	"encoding/json"
	"net/http"

	l "github.com/redhatinsights/platform-changelog-go/internal/logging"
	"github.com/redhatinsights/platform-changelog-go/internal/metrics"
	"github.com/redhatinsights/platform-changelog-go/internal/structs"
)

func (eh *EndpointHandler) GetServicesAll(w http.ResponseWriter, r *http.Request) {
	metrics.IncRequests(r.URL.Path, r.Method, r.UserAgent())

	q, err := initQuery(r)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, "Invalid query")
		return
	}

	servicesWithTimelines, count, err := eh.conn.GetServicesAll(q.Offset, q.Limit, q)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	servicesList := structs.ExpandedServicesList{Count: count, Data: servicesWithTimelines}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(servicesList)
}

func (eh *EndpointHandler) GetServiceByID(w http.ResponseWriter, r *http.Request) {
	metrics.IncRequests(r.URL.Path, r.Method, r.UserAgent())

	serviceID, err := getIDFromURL(r, "service_id")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid service ID"))
		return
	}

	service, _, err := eh.conn.GetServiceByID(serviceID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	/**
	 * the service doesn't exist
	 */
	if service.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Service not found"))
		return
	}

	serviceData := structs.ServicesData{
		ID:          service.ID,
		Name:        service.Name,
		DisplayName: service.DisplayName,
		Tenant:      service.Tenant,
		Projects:    convertProjectsToProjectsData(service.Projects),
	}

	l.Log.Debugf("URL Param: %d", serviceID)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(serviceData)
}
