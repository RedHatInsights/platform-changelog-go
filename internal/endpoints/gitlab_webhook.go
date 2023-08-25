package endpoints

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"fmt"

	"github.com/redhatinsights/platform-changelog-go/internal/config"
	l "github.com/redhatinsights/platform-changelog-go/internal/logging"
	"github.com/redhatinsights/platform-changelog-go/internal/metrics"
	m "github.com/redhatinsights/platform-changelog-go/internal/models"
	"github.com/redhatinsights/platform-changelog-go/internal/structs"
	"github.com/xanzy/go-gitlab"
)

func getURL(p *gitlab.PushEvent) string {
	if p == nil || p.Repository == nil {
		return ""
	}
	return p.Repository.Homepage
}

func getRepo(p *gitlab.PushEvent) *gitlab.Repository {
	if p == nil || p.Repository == nil {
		return nil
	}
	return p.Repository
}

type RepInfo *struct {
	ID        string     `json:"id"`
	Message   string     `json:"message"`
	Title     string     `json:"title"`
	Timestamp *time.Time `json:"timestamp"`
	URL       string     `json:"url"`
	Author    struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"author"`
	Added    []string `json:"added"`
	Modified []string `json:"modified"`
	Removed  []string `json:"removed"`
}

func getID(p RepInfo) string {
	if p == nil {
		return ""
	}
	return p.ID
}

func getTime(p *time.Time) time.Time {
	if p == nil {
		return time.Time{}
	}
	return *p
}

type PingEvent struct {
	// Random string of GitHub zen.
	Zen *string `json:"zen,omitempty"`
	// The ID of the webhook that triggered the ping.
	HookID *int64 `json:"hook_id,omitempty"`
	// The webhook configuration.
	Hook *gitlab.Hook `json:"hook,omitempty"`
	//Installation *Installation `json:"installation,omitempty"`
}

func getAuthor(p RepInfo) string {
	if p == nil {
		return ""
	}
	return p.Author.Name
}

func getName(p *gitlab.PushEvent) string {
	if p == nil {
		return ""
	}
	return p.UserName
}

func getMessage(p RepInfo) string {
	if p == nil {
		return ""
	}
	return p.Message
}

// GitlabWebhook gets data from the webhook and enters it into the DB
func (eh *EndpointHandler) GitlabWebhook(w http.ResponseWriter, r *http.Request) {

	var err error
	var payload []byte

	metrics.IncWebhooks("gitlab", r.Method, r.UserAgent(), false)

	if config.Get().SkipWebhookValidation {
		l.Log.Info("skipping webhook validation")
	} else {
		if config.Get().GitlabWebhookSecretKey == "" {
			l.Log.Error("missing gitlab webhook secret key")
			writeResponse(w, http.StatusInternalServerError, `{"msg": "server is missing gitlab webhook secret key"}`)
			metrics.IncWebhooks("gitlab", r.Method, r.UserAgent(), true)
			return
		}

		token := r.Header.Get("X-Gitlab-Token")

		if token == "" || token != config.Get().GitlabWebhookSecretKey {
			l.Log.Error("invalid or missing X-Gitlab-Token")
			writeResponse(w, http.StatusBadRequest, `{"msg": "invalid or missing X-Gitlab-Token"}`)
			metrics.IncWebhooks("gitlab", r.Method, r.UserAgent(), true)
			return
		}
	}

	payload, err = ioutil.ReadAll(r.Body)

	if err != nil {
		l.Log.Error(err)
		writeResponse(w, http.StatusBadRequest, fmt.Sprintf(`{"msg": "%s"}`, err.Error()))
		metrics.IncWebhooks("gitlab", r.Method, r.UserAgent(), true)
		return
	}
	defer r.Body.Close()

	event, err := gitlab.ParseWebhook(gitlab.WebhookEventType(r), payload)
	if err != nil {
		l.Log.Errorf("could not parse webhook: err=%s\n", err)
		writeResponse(w, http.StatusBadRequest, `{"msg": "Could not parse webhook"}`)
		metrics.IncWebhooks("gitlab", r.Method, r.UserAgent(), true)
		return
	}

	switch e := event.(type) {

	case PingEvent:
		writeResponse(w, http.StatusOK, `{"msg": "ok"}`)
		return

	case *gitlab.PushEvent:
		repo := getURL(e)
		project, err := eh.conn.GetProjectByRepo(repo)
		if err != nil {
			// project not onboarded; build project and find service if available

			// Due to the webhook events not being connected to app-interface,
			// The service name and project name will be the same.
			// Also, the tenant will not be specified.
			// A user could override these by modifying the service.yml.
			service, _, err := eh.conn.GetServiceByName(getRepo(e).Name)

			if err != nil { // service not found
				// create service too
				newService := m.Services{
					Name:        getRepo(e).Name,
					DisplayName: getRepo(e).Name,
					Tenant:      "undefined",
				}
				eh.conn.CreateServiceTableEntry(newService)

				service, _, err = eh.conn.GetServiceByName(getRepo(e).Name)
				if err != nil {
					// Failed to create service entry, something must be wrong with db
					l.Log.Errorf("Failed to insert new service: %v", err)
					metrics.IncWebhooks("gitlab", r.Method, r.UserAgent(), true)
					writeResponse(w, http.StatusInternalServerError, `{"msg": "Failed to insert new service"}`)
					return
				}
			}

			newProject := m.Projects{
				ServiceID:  service.ID,
				Name:       getRepo(e).Name,
				Repo:       repo,
				Namespaces: []string{},
				Branches:   []string{strings.Split(e.Ref, "/")[2]},
			}

			err = eh.conn.CreateProjectTableEntry(newProject)

			if err != nil {
				l.Log.Info("Failed to create project: ", newProject)
			}

			// retry to get the project (for the id)
			project, err = eh.conn.GetProjectByRepo(repo)
			if err != nil {
				l.Log.Errorf("Failed to insert new project: %v", err)
				metrics.IncWebhooks("gitlab", r.Method, r.UserAgent(), true)
				writeResponse(w, http.StatusInternalServerError, `{"msg": "Failed to insert new project"}`)
				return
			}
		}

		commitData := getCommitData2(e, project)

		err = eh.conn.BulkCreateCommitEntry(commitData)
		if err != nil {
			l.Log.Errorf("Failed to insert webhook data: %v", err)
			metrics.IncWebhooks("gitlab", r.Method, r.UserAgent(), true)
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

func getCommitData2(g *gitlab.PushEvent, p structs.ProjectsData) []m.Timelines {
	var commits []m.Timelines
	for _, commit := range g.Commits {
		record := m.Timelines{
			ServiceID: p.ID,
			Repo:      getRepo(g).Name,
			Ref:       getID(commit),
			Type:      "commit",
			Timestamp: getTime(commit.Timestamp),
			Author:    getAuthor(commit),
			MergedBy:  getName(g),
			Message:   getMessage(commit),
		}
		commits = append(commits, record)
	}

	return commits
}
