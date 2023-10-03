package metrics

import (
	"net/http"
	"strconv"

	p "github.com/prometheus/client_golang/prometheus"
	pa "github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	requests = pa.NewCounterVec(p.CounterOpts{
		Name: "platform_changelog_requests",
		Help: "Number of requests",
	}, []string{"path", "method", "user_agent"})

	webhooks = pa.NewCounterVec(p.CounterOpts{
		Name: "platform_changelog_webhooks",
		Help: "Number of incoming webhooks",
	}, []string{"source", "method", "user_agent"})

	webhookErrors = pa.NewCounterVec(p.CounterOpts{
		Name: "platform_changelog_webhooks_errors",
		Help: "Number of incoming webhook errors",
	}, []string{"source", "method", "user_agent"})

	jenkins = pa.NewCounterVec(p.CounterOpts{
		Name: "platform_changelog_jenkins",
		Help: "Number of jenkins endpoint hits",
	}, []string{"source", "method", "user_agent"})

	jenkinsErrors = pa.NewCounterVec(p.CounterOpts{
		Name: "platform_changelog_jenkins_errors",
		Help: "Number of jenkins endpoint errors",
	}, []string{"source", "method", "user_agent"})

	tekton = pa.NewCounterVec(p.CounterOpts{
		Name: "platform_changelog_tekton_runs",
		Help: "Number of tekton runs",
	}, []string{"method", "user_agent"})

	tektonErrors = pa.NewCounterVec(p.CounterOpts{
		Name: "platform_changelog_tekton_errors",
		Help: "Number of tekton errors",
	}, []string{"method", "user_agent"})

	responseCodes = pa.NewCounterVec(p.CounterOpts{
		Name: "platform_changelog_responses",
		Help: "Number of response codes",
	}, []string{"code"})

	SqlCreateCommitEntry = pa.NewHistogram(p.HistogramOpts{
		Name: "platform_changelog_sql_create_commit_entry_seconds",
		Help: "Elapsed time for sql creation of commit entry",
	})

	SqlGetServicesAll = pa.NewHistogram(p.HistogramOpts{
		Name: "platform_changelog_sql_get_services_all_seconds",
		Help: "Elapsed time for sql lookup of all services",
	})

	SqlGetTimelinesAll = pa.NewHistogram(p.HistogramOpts{
		Name: "platform_changelog_sql_get_timelines_all_seconds",
		Help: "Elapsed time for sql lookup of timeline entries",
	})

	SqlGetCommitsAll = pa.NewHistogram(p.HistogramOpts{
		Name: "platform_changelog_sql_get_commits_all_seconds",
		Help: "Elapsed time for sql lookup of all commits",
	})

	SqlGetDeploysAll = pa.NewHistogram(p.HistogramOpts{
		Name: "platform_changelog_sql_get_deploys_all_seconds",
		Help: "Elapsed time for sql lookup of all deploys",
	})

	SqlGetServiceByID = pa.NewHistogram(p.HistogramOpts{
		Name: "platform_changelog_sql_get_service_by_id_seconds",
		Help: "Elapsed time for sql lookup of services by name",
	})

	SqlGetServiceByName = pa.NewHistogram(p.HistogramOpts{
		Name: "platform_changelog_sql_get_service_by_name_seconds",
		Help: "Elapsed time for sql lookup of services by name",
	})

	SqlGetProjectByID = pa.NewHistogram(p.HistogramOpts{
		Name: "platform_changelog_sql_get_project_by_id_seconds",
		Help: "Elapsed time for sql lookup of projects by name",
	})

	SqlGetProjectByName = pa.NewHistogram(p.HistogramOpts{
		Name: "platform_changelog_sql_get_project_by_name_seconds",
		Help: "Elapsed time for sql lookup of projects by name",
	})

	SqlGetProjectsByService = pa.NewHistogram(p.HistogramOpts{
		Name: "platform_changelog_sql_get_project_by_service_seconds",
		Help: "Elapsed time for sql lookup of projects by service",
	})

	SqlGetCommitsByService = pa.NewHistogram(p.HistogramOpts{
		Name: "platform_changelog_sql_get_commits_by_service_seconds",
		Help: "Elapsed time for sql lookup of commits by service",
	})

	SqlGetDeploysByService = pa.NewHistogram(p.HistogramOpts{
		Name: "platform_changelog_sql_get_deploys_by_service_seconds",
		Help: "Elapsed time for sql lookup of deploys by service",
	})

	SqlGetTimelinesByService = pa.NewHistogram(p.HistogramOpts{
		Name: "platform_changelog_sql_get_timelines_by_service_seconds",
		Help: "Elapsed time for sql lookup of a service's timeline entries",
	})

	SqlGetCommitsByProject = pa.NewHistogram(p.HistogramOpts{
		Name: "platform_changelog_sql_get_commits_by_project_seconds",
		Help: "Elapsed time for sql lookup of commits by project",
	})

	SqlGetDeploysByProject = pa.NewHistogram(p.HistogramOpts{
		Name: "platform_changelog_sql_get_deploys_by_project_seconds",
		Help: "Elapsed time for sql lookup of deploys by project",
	})

	SqlGetTimelinesByProject = pa.NewHistogram(p.HistogramOpts{
		Name: "platform_changelog_sql_get_timelines_by_project_seconds",
		Help: "Elapsed time for sql lookup of timelines by project",
	})

	SqlGetCommitByRef = pa.NewHistogram(p.HistogramOpts{
		Name: "platform_changelog_sql_get_commit_by_ref_seconds",
		Help: "Elapsed time for sql lookup of commit by ref",
	})

	SqlGetDeployByRef = pa.NewHistogram(p.HistogramOpts{
		Name: "platform_changelog_sql_get_deploy_by_ref_seconds",
		Help: "Elapsed time for sql lookup of deploy by ref",
	})

	SqlGetTimelineByRef = pa.NewHistogram(p.HistogramOpts{
		Name: "platform_changelog_sql_get_timeline_by_ref_seconds",
		Help: "Elapsed time for sql lookup of timeline by ref",
	})
)

type MetricsTrackingResponseWriter struct {
	Wrapped   http.ResponseWriter
	UserAgent string
}

func IncRequests(path string, method string, userAgent string) {
	requests.With(p.Labels{"path": path, "method": method, "user_agent": userAgent}).Inc()
}

func IncWebhooks(source string, method string, userAgent string, err bool) {
	if !err {
		webhooks.With(p.Labels{"source": source, "method": method, "user_agent": userAgent}).Inc()
	} else {
		webhookErrors.With(p.Labels{"source": source, "method": method, "user_agent": userAgent}).Inc()
	}
}

func IncJenkins(source string, method string, userAgent string, err bool) {
	if !err {
		jenkins.With(p.Labels{"source": source, "method": method, "user_agent": userAgent}).Inc()
	} else {
		jenkinsErrors.With(p.Labels{"source": source, "method": method, "user_agent": userAgent}).Inc()
	}
}

func IncTekton(method string, userAgent string, err bool) {
	if !err {
		tekton.With(p.Labels{"method": method, "user_agent": userAgent}).Inc()
	} else {
		tektonErrors.With(p.Labels{"method": method, "user_agent": userAgent}).Inc()
	}
}

func (m *MetricsTrackingResponseWriter) Header() http.Header {
	return m.Wrapped.Header()
}

func (m *MetricsTrackingResponseWriter) WriteHeader(statusCode int) {
	responseCodes.With(p.Labels{"code": strconv.Itoa(statusCode)}).Inc()
	m.Wrapped.WriteHeader(statusCode)
}

func (m *MetricsTrackingResponseWriter) Write(b []byte) (int, error) {
	return m.Wrapped.Write(b)
}

func ResponseMetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mw := &MetricsTrackingResponseWriter{
			Wrapped:   w,
			UserAgent: r.Header.Get("User-Agent"),
		}
		next.ServeHTTP(mw, r)
	})
}
