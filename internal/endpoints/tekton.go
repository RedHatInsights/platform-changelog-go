package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/redhatinsights/platform-changelog-go/internal/config"
	"github.com/redhatinsights/platform-changelog-go/internal/db"
	l "github.com/redhatinsights/platform-changelog-go/internal/logging"
	"github.com/redhatinsights/platform-changelog-go/internal/metrics"
	"github.com/redhatinsights/platform-changelog-go/internal/models"
)

type TektonPayload *struct {
	Status      string     `json:"status"`
	Timestamp   *time.Time `json:"timestamp"`
	App         string     `json:"app"`
	Env         string     `json:"env"`
	TriggeredBy string     `json:"triggered_by"`
	Ref         string     `json:"ref,omitempty"`
}

func decodeJSONBody(w http.ResponseWriter, r *http.Request) (TektonPayload, error) {
	if r.Header.Get("Content-Type") != "application/json" {
		return nil, fmt.Errorf("invalid Content-Type")
	}

	var payload TektonPayload

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func TektonTaskRun(w http.ResponseWriter, r *http.Request) {
	metrics.IncTekton(r.Method, r.UserAgent(), false)

	payload, err := decodeJSONBody(w, r)
	if err != nil {
		l.Log.Error(err)
		metrics.IncTekton(r.Method, r.UserAgent(), true)
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	deploy, err := ConvertTektonPayloadToTimeline(payload)

	if err != nil {
		l.Log.Error(err)
		metrics.IncTekton(r.Method, r.UserAgent(), true)
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	result := db.CreateDeployEntry(db.DB, deploy)

	if result.Error != nil {
		l.Log.Error(result.Error)
		metrics.IncTekton(r.Method, r.UserAgent(), true)
		writeResponse(w, http.StatusInternalServerError, result.Error.Error())
		return
	}

	writeResponse(w, http.StatusOK, `{"msg": "Tekton info received"}`)
}

// Converting from TektonPayload struct to Timeline model
func ConvertTektonPayloadToTimeline(payload TektonPayload) (models.Timelines, error) {
	services := config.Get().Services

	var deploy models.Timelines
	// Validate that the app specified is onboarded
	for key, service := range services {
		if service.Namespace == payload.App {
			_, s := db.GetServiceByName(db.DB, key)

			deploy = models.Timelines{
				ServiceID:       s.ID,
				Timestamp:       *payload.Timestamp,
				Type:            "deploy",
				Repo:            s.Name,
				Ref:             payload.Ref,
				DeployNamespace: payload.App,
				Cluster:         payload.Env,
				TriggeredBy:     payload.TriggeredBy,
				Status:          payload.Status,
			}
		}
	}

	return deploy, nil
}
