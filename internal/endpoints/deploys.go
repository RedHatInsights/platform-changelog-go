package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/redhatinsights/platform-changelog-go/internal/db"
	"github.com/redhatinsights/platform-changelog-go/internal/metrics"
	"github.com/redhatinsights/platform-changelog-go/internal/structs"
)

func GetDeploysAll(w http.ResponseWriter, r *http.Request) {
	metrics.IncRequests(r.URL.Path, r.Method, r.UserAgent())

	q, err := initQuery(r)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, "Invalid query")
		return
	}

	deploys, count, err := db.GetDeploysAll(db.DB, q.Offset, q.Limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	deploysList := structs.TimelinesList{count, deploys}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deploysList)
}

func GetDeploysByService(w http.ResponseWriter, r *http.Request) {
	metrics.IncRequests(r.URL.Path, r.Method, r.UserAgent())
	serviceName := chi.URLParam(r, "service")

	q, err := initQuery(r)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, "Invalid query")
		return
	}

	service, _, err := db.GetServiceByName(db.DB, serviceName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Couldn't find the service"))
		return
	}

	deploys, count, err := db.GetDeploysByService(db.DB, service, q.Offset, q.Limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	deploysList := structs.TimelinesList{count, deploys}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deploysList)
}

func GetDeployByRef(w http.ResponseWriter, r *http.Request) {
	metrics.IncRequests(r.URL.Path, r.Method, r.UserAgent())
	ref := chi.URLParam(r, "ref")

	deploy, rowsAffected, err := db.GetDeployByRef(db.DB, ref)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	if rowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Deploy not found"))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deploy)
}
