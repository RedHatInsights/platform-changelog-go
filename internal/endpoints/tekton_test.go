package endpoints_test

import (
	"net/http"
	"net/http/httptest"
	"os"

	chi "github.com/go-chi/chi/v5"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redhatinsights/platform-changelog-go/internal/config"
	"github.com/redhatinsights/platform-changelog-go/internal/db"
	"github.com/redhatinsights/platform-changelog-go/internal/endpoints"
	"github.com/redhatinsights/platform-changelog-go/internal/logging"
)

var _ = Describe("Handler", func() {

	logging.InitLogger()

	// test the TektonTaskRun function
	DescribeTable("TektonTaskRun", func(expected_status int, message string, data_path string) {

		f, err := os.Open(data_path)
		Expect(err).To(BeNil())

		defer f.Close()

		// create a mock db connection & endpoint handler
		var cfg config.Config = config.Config{
			DBImpl: "mock",
		}
		dbConnector := db.NewMockDBConnector(&cfg)
		handler := endpoints.NewHandler(dbConnector)

		// create a request
		req, err := http.NewRequest("POST", "/api/v1/tekton", f)
		Expect(err).To(BeNil())

		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		router := chi.NewRouter()
		router.Post("/api/v1/tekton", handler.TektonTaskRun)

		router.ServeHTTP(rr, req)

		Expect(rr.Code).To(Equal(expected_status))
		Expect(rr.Body.String()).To(ContainSubstring(message))
	},
		Entry("Valid", http.StatusOK, "Tekton info received", "../../tests/tekton/valid.json"),
		Entry("Empty", http.StatusBadRequest, "json body is required", "../../tests/empty.json"),
		Entry("Missing timestamp", http.StatusBadRequest, "timestamp is required", "../../tests/tekton/missing_timestamp.json"),
		Entry("Missing app", http.StatusBadRequest, "app is required", "../../tests/tekton/missing_app.json"),
		Entry("Missing status", http.StatusBadRequest, "status is required", "../../tests/tekton/missing_status.json"),
	)
})
