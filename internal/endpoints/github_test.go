package endpoints_test

import (
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/redhatinsights/platform-changelog-go/internal/models"

	chi "github.com/go-chi/chi/v5"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redhatinsights/platform-changelog-go/internal/db"
	"github.com/redhatinsights/platform-changelog-go/internal/endpoints"
	"github.com/redhatinsights/platform-changelog-go/internal/logging"
)

var _ = Describe("Handler", func() {

	logging.InitLogger()

	Describe("Github Jenkins Run with empty body", func() {
		It("should return 400", func() {
			dbConnector := testDBImpl
			handler := endpoints.NewHandler(dbConnector)

			// create a request
			req, err := http.NewRequest("POST", "/api/v1/github", nil)
			Expect(err).To(BeNil())

			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()

			router := chi.NewRouter()
			router.Post("/api/v1/github", handler.TektonTaskRun)

			router.ServeHTTP(rr, req)

			Expect(rr.Code).To(Equal(http.StatusBadRequest))
			Expect(rr.Body.String()).To(Equal("json body required"))
		})
	})

	// test the TektonTaskRun function
	DescribeTable("Github Jenkins Run with JSON body", func(expected_status int, message string, data_path string) {

		f, err := os.Open(data_path)
		Expect(err).To(BeNil())

		defer f.Close()

		dbConnector := testDBImpl
		handler := endpoints.NewHandler(dbConnector)

		// create a request
		req, err := http.NewRequest("POST", "/api/v1/github", f)
		Expect(err).To(BeNil())

		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		router := chi.NewRouter()
		router.Post("/api/v1/github", handler.Github)

		router.ServeHTTP(rr, req)

		// Expect(rr.Code).To(Equal(expected_status))
		Expect(rr.Body.String()).To(ContainSubstring(message))
	},
		Entry("Valid; Project not in db", http.StatusOK, "Commit info received", "../../tests/jenkins/github_dump.json"),
		Entry("Valid; Project in db", http.StatusOK, "Commit info received", "../../tests/jenkins/github_dump.json"), // now onboarded
		Entry("Valid; Project and Service not in db", http.StatusOK, "Commit info received", "../../tests/jenkins/not_onboarded.json"),
		Entry("Empty", http.StatusBadRequest, "empty json body provided", "../../tests/empty.json"),
		Entry("Missing app", http.StatusBadRequest, "app is required", "../../tests/jenkins/missing_app.json"),
		Entry("Missing tenant", http.StatusBadRequest, "tenant is required", "../../tests/jenkins/missing_tenant.json"),
		Entry("Missing project", http.StatusBadRequest, "project is required", "../../tests/jenkins/missing_project.json"),
		Entry("Missing ref", http.StatusBadRequest, "ref is required", "../../tests/jenkins/missing_ref.json"),
		Entry("Missing branch", http.StatusBadRequest, "branch is required", "../../tests/jenkins/missing_branch.json"),
	)
})

func CreateService(conn db.DBConnector, name string, s models.Services) (service models.Services) {
	service, err := conn.CreateServiceTableEntry(s)

	Expect(err).To(BeNil())

	return service
}
