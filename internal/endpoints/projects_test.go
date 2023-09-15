package endpoints_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/go-chi/chi/v5"
	"github.com/redhatinsights/platform-changelog-go/internal/endpoints"
	"github.com/redhatinsights/platform-changelog-go/internal/logging"
	"github.com/redhatinsights/platform-changelog-go/internal/models"
	"github.com/redhatinsights/platform-changelog-go/internal/structs"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handler", Ordered, func() {

	logging.InitLogger()

	var projects []models.Projects
	var count int64

	BeforeAll(func() {
		// get all services from the database
		db := testDBImpl

		projects, rowsAffected, err := db.GetProjectsAll(0, 100, structs.Query{})
		Expect(err).To(BeNil())

		count = rowsAffected

		Expect(len(projects)).To(Equal(int(count)))
	})

	// test get all projects
	Describe("Get projects", func() {
		It("should return all projects", func() {
			dbConnector := testDBImpl
			handler := endpoints.NewHandler(dbConnector)

			// create a request
			req, err := http.NewRequest("GET", "/projects?limit=100&offset=0", nil)
			Expect(err).To(BeNil())

			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()

			router := chi.NewRouter()
			router.Get("/projects", handler.GetServicesAll)

			router.ServeHTTP(rr, req)

			Expect(rr.Code).To(Equal(http.StatusOK))
			body := rr.Body.String()

			Expect(body).To(ContainSubstring(fmt.Sprintf("\"count\":%d,", count)))
			Expect(body).To(ContainSubstring("\"data\":[{"))

			for _, p := range projects {
				Expect(body).To(ContainSubstring("\"name\":\"" + p.Name + "\""))
			}
		})

		It("should return project by service", func() {
			dbConnector := testDBImpl
			handler := endpoints.NewHandler(dbConnector)

			// create a request
			req, err := http.NewRequest("GET", "/services/platform-changelog/projects", nil)
			Expect(err).To(BeNil())

			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()

			router := chi.NewRouter()
			router.Get("/services/{service}/projects", handler.GetProjectsByService)

			router.ServeHTTP(rr, req)

			Expect(rr.Code).To(Equal(http.StatusOK))
			body := rr.Body.String()

			Expect(body).To(ContainSubstring("\"count\":2"))
			Expect(body).To(ContainSubstring("\"data\":[{"))

			Expect(body).To(ContainSubstring("\"name\":\"platform-changelog-go\""))
			Expect(body).To(ContainSubstring("\"name\":\"platform-changelog-ui\""))
		})
	})
})
