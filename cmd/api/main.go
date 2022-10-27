package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/redhatinsights/platform-changelog-go/internal/config"
	"github.com/redhatinsights/platform-changelog-go/internal/db"
	"github.com/redhatinsights/platform-changelog-go/internal/endpoints"
	"github.com/redhatinsights/platform-changelog-go/internal/logging"
	"github.com/redhatinsights/platform-changelog-go/internal/metrics"
)

func lubdub(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("lubdub"))
}

func openAPIHandler(cfg *config.Config) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(cfg.OpenAPISpec)
	})
}

func main() {

	logging.InitLogger()

	cfg := config.Get()

	// // TODO: Switch between mock and impl based on config
	dbConnector := db.NewDBConnector(cfg)
	handler := endpoints.NewHandler(dbConnector)

	r := chi.NewRouter()
	mr := chi.NewRouter()
	sub := chi.NewRouter().With(metrics.ResponseMetricsMiddleware)

	// Mount the root of the api router on /api/v1
	r.Mount("/api/v1", sub)
	r.Get("/", lubdub)

	mr.Get("/", lubdub)
	mr.Get("/healthz", lubdub)
	mr.Handle("/metrics", promhttp.Handler())

	sub.Post("/github-webhook", handler.GithubWebhook)
	sub.Post("/gitlab-webhook", handler.GitlabWebhook)

	sub.Get("/services", handler.GetServicesAll)
	sub.Get("/timelines", handler.GetTimelinesAll)
	sub.Get("/commits", handler.GetCommitsAll)
	sub.Get("/deploys", handler.GetDeploysAll)

	sub.Get("/services/{service}", handler.GetServiceByName)
	sub.Get("/services/{service}/timelines", handler.GetTimelinesByService)
	sub.Get("/services/{service}/commits", handler.GetCommitsByService)
	sub.Get("/services/{service}/deploys", handler.GetDeploysByService)

	sub.Get("/timelines/{ref}", handler.GetTimelineByRef)
	sub.Get("/commits/{ref}", handler.GetCommitByRef)
	sub.Get("/deploys/{ref}", handler.GetDeployByRef)

	sub.Get("/openapi.json", openAPIHandler(cfg))

	srv := http.Server{
		Addr:    ":" + cfg.PublicPort,
		Handler: r,
	}

	msrv := http.Server{
		Addr:    ":" + cfg.MetricsPort,
		Handler: mr,
	}

	go func() {
		if err := msrv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
