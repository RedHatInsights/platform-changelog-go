package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/redhatinsights/platform-changelog-go/internal/structs"

	"github.com/go-chi/chi/v5"
	"github.com/redhatinsights/platform-changelog-go/internal/db"
	"github.com/redhatinsights/platform-changelog-go/internal/metrics"
)

func GetCommitsAll(w http.ResponseWriter, r *http.Request) {
	metrics.IncRequests(r.URL.Path, r.Method, r.UserAgent())

	q, err := initQuery(r)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, "Invalid query")
		return
	}

	commits, count, err := db.GetCommitsAll(db.DB, q.Offset, q.Limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	commitsList := structs.TimelinesList{count, commits}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(commitsList)
}

func GetCommitsByService(w http.ResponseWriter, r *http.Request) {
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

	commits, count, err := db.GetCommitsByService(db.DB, service, q.Offset, q.Limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	commitsList := structs.TimelinesList{count, commits}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(commitsList)
}

func GetCommitByRef(w http.ResponseWriter, r *http.Request) {
	metrics.IncRequests(r.URL.Path, r.Method, r.UserAgent())
	ref := chi.URLParam(r, "ref")

	commit, rowsAffected, err := db.GetCommitByRef(db.DB, ref)
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
