package endpoints

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/google/go-github/v50/github"
	"github.com/redhatinsights/platform-changelog-go/internal/config"
	l "github.com/redhatinsights/platform-changelog-go/internal/logging"
	"github.com/redhatinsights/platform-changelog-go/internal/metrics"
	"github.com/redhatinsights/platform-changelog-go/internal/models"
	"github.com/redhatinsights/platform-changelog-go/internal/utils"
)

// GithubWebhook gets data from the webhook and enters it into the DB
func (eh *EndpointHandler) GithubWebhook(w http.ResponseWriter, r *http.Request) {

	var err error
	var payload []byte

	metrics.IncWebhooks("github", r.Method, r.UserAgent(), false)

	if config.Get().SkipWebhookValidation {
		l.Log.Info("skipping webhook validation")
		payload, err = ioutil.ReadAll(r.Body)
	} else {
		if config.Get().GithubWebhookSecretKey == "" {
			l.Log.Error("invalid or missing github webhook secret key")
			writeResponse(w, http.StatusInternalServerError, `{"msg": "server has an invalid or missing github webhook secret key"}`)
			metrics.IncWebhooks("github", r.Method, r.UserAgent(), true)
			return
		}

		payload, err = github.ValidatePayload(r, []byte(config.Get().GithubWebhookSecretKey))
	}

	if err != nil {
		l.Log.Error(err)
		writeResponse(w, http.StatusUnauthorized, fmt.Sprintf(`{"msg": "%s"}`, err.Error()))
		metrics.IncWebhooks("github", r.Method, r.UserAgent(), true)
		return
	}
	defer r.Body.Close()

	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		l.Log.Errorf("could not parse webhook: err=%s\n", err)
		writeResponse(w, http.StatusBadRequest, `{"msg": "could not parse webhook"}`)
		metrics.IncWebhooks("github", r.Method, r.UserAgent(), true)
		return
	}

	switch e := event.(type) {
	case *github.PingEvent:
		writeResponse(w, http.StatusOK, `{"msg": "ok"}`)
		return
	case *github.PushEvent:
		repo := e.Repo.GetURL()
		project, err := eh.conn.GetProjectByRepo(repo)
		if err != nil {
			// project not onboarded; build project and find service if available

			// Due to the webhook events not being connected to app-interface,
			// The service name and project name will be the same.
			// Also, the tenant will not be specified.
			// A user could override these by modifying the service.yml.
			service, _, err := eh.conn.GetServiceByName(e.Repo.GetName())

			if err != nil { // service not found
				// create service too
				service = models.Services{
					Name:        e.Repo.GetName(),
					DisplayName: e.Repo.GetName(),
					Tenant:      "undefined",
				}
				eh.conn.CreateServiceTableEntry(&service)

				service, _, err = eh.conn.GetServiceByName(e.Repo.GetName())
				if err != nil {
					// Failed to create service entry, something must be wrong with db
					l.Log.Errorf("Failed to insert new service: %v", err)
					metrics.IncWebhooks("github", r.Method, r.UserAgent(), true)
					writeResponse(w, http.StatusInternalServerError, `{"msg": "Failed to insert new service"}`)
					return
				}
			}

			project = models.Projects{
				ServiceID: service.ID,
				Name:      e.Repo.GetName(),
				Repo:      repo,
				Branch:    strings.Split(utils.DerefString(e.Ref), "/")[2],
			}

			err = eh.conn.CreateProjectTableEntry(&project)

			if err != nil {
				l.Log.Info("Failed to insert project: ", project)
				metrics.IncWebhooks("github", r.Method, r.UserAgent(), true)
				writeResponse(w, http.StatusInternalServerError, `{"msg": "Failed to insert new project"}`)
			}
		}

		commitData := getCommitData(e, project)

		err = eh.conn.BulkCreateCommitEntry(commitData)
		if err != nil {
			l.Log.Errorf("Failed to insert webhook data: %v", err)
			metrics.IncWebhooks("github", r.Method, r.UserAgent(), true)
			writeResponse(w, http.StatusInternalServerError, `{"msg": "Failed to insert webhook data"}`)
			return
		}

		l.Log.Infof("Created %d commit entries for %s", len(commitData), project.Name)
		writeResponse(w, http.StatusOK, `{"msg": "ok"}`)
		return

	default:
		l.Log.Errorf("Event type %T not supported", e)
		writeResponse(w, http.StatusOK, `{"msg": "Event from this repo is not a push event"}`)
		return
	}
}

func getCommitData(g *github.PushEvent, p models.Projects) []models.Timelines {
	var commits []models.Timelines
	for _, commit := range g.Commits {
		record := models.Timelines{
			ServiceID: p.ServiceID,
			ProjectID: p.ID,
			Repo:      utils.DerefString(g.GetRepo().Name),
			Ref:       commit.GetID(),
			Type:      "commit",
			Timestamp: commit.Timestamp.Time,
			Author:    utils.DerefString(commit.GetAuthor().Login),
			MergedBy:  g.Pusher.GetName(),
			Message:   commit.GetMessage(),
		}
		commits = append(commits, record)
	}

	return commits
}
