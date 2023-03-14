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

	services := config.Get().Services

	if config.Get().SkipWebhookValidation {
		l.Log.Info("skipping webhook validation")
	} else {
		if config.Get().GitlabWebhookSecretKey == "" {
			l.Log.Error("missing gitlab webhook secret key")
			writeResponse(w, http.StatusInternalServerError, `{"msg": "missing github webhook secret key"}`)
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
		for key, service := range services {
			if service.GLRepo == getURL(e) {
				s, _, _ := eh.conn.GetServiceByName(key)
				if s.Branch != strings.Split((e.Ref), "/")[2] {
					l.Log.Info("Branch mismatch: ", s.Branch, " != ", strings.Split((e.Ref), "/")[2])
					writeResponse(w, http.StatusOK, `{"msg": "Not a monitored branch"}`)
					return
				}
				commitData := getCommitData2(e, s)
				err := eh.conn.CreateCommitEntry(commitData)
				if err != nil {
					l.Log.Errorf("Failed to insert webhook data: %v", err)
					metrics.IncWebhooks("gitlab", r.Method, r.UserAgent(), true)
					writeResponse(w, http.StatusInternalServerError, `{"msg": "Failed to insert webhook data"}`)
					return
				}
				l.Log.Infof("Created %d commit entries for %s", len(commitData), key)
				writeResponse(w, http.StatusOK, `{"msg": "ok"}`)
				return
			}
		}
		// catch for if the service is not registered
		l.Log.Infof("Service not found for %s", getURL(e))
		fmt.Println(getURL(e))

		writeResponse(w, http.StatusOK, `{"msg": "The service is not registered"}`)
		return
	default:
		l.Log.Errorf("Event type %T not supported", e)
		writeResponse(w, http.StatusOK, `{"msg": "Event from this repo is not a push event"}`)
		return
	}
}

func getCommitData2(g *gitlab.PushEvent, s structs.ServicesData) []m.Timelines {
	var commits []m.Timelines
	for _, commit := range g.Commits {
		record := m.Timelines{
			ServiceID: s.ID,
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
