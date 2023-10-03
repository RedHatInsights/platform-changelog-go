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

	var services []structs.ExpandedServicesData
	var count int64

	BeforeAll(func() {
		// get all services from the database
		db := testDBImpl

		services, rowsAffected, err := db.GetServicesAll(0, 100, structs.Query{})

		Expect(err).To(BeNil())

		count = rowsAffected

		Expect(len(services)).To(Equal(int(count)))
	})

	Describe("Get services", func() {
		It("should return all services", func() {
			dbConnector := testDBImpl
			handler := endpoints.NewHandler(dbConnector)

			// create a request
			req, err := http.NewRequest("GET", "/api/v1/services?limit=100&offset=0", nil)
			Expect(err).To(BeNil())

			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()

			router := chi.NewRouter()
			router.Get("/api/v1/services", handler.GetServicesAll)

			router.ServeHTTP(rr, req)

			Expect(rr.Code).To(Equal(http.StatusOK))
			body := rr.Body.String()

			Expect(body).To(ContainSubstring("\"count\":" + fmt.Sprint(count)))
			Expect(body).To(ContainSubstring("\"data\":[{"))

			for _, s := range services {
				Expect(body).To(ContainSubstring(s.Name))

				if (s.Commit != models.Timelines{}) {
					Expect(body).To(ContainSubstring(s.Commit.Ref))
				}
				if (s.Deploy != models.Timelines{}) {
					Expect(body).To(ContainSubstring(s.Deploy.Status))
				}
			}
		})

		It("should return service by name", func() {
			dbConnector := testDBImpl
			handler := endpoints.NewHandler(dbConnector)

			// create a request
			req, err := http.NewRequest("GET", "/services/1", nil) // platform-changelog
			Expect(err).To(BeNil())

			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()

			router := chi.NewRouter()
			router.Get("/services/{service_id}", handler.GetServiceByID)

			router.ServeHTTP(rr, req)

			Expect(rr.Body.String()).To(ContainSubstring("platform-changelog"))
			Expect(rr.Code).To(Equal(http.StatusOK))
			body := rr.Body.String()

			Expect(body).To(ContainSubstring("platform-changelog"))
			Expect(body).To(ContainSubstring("Insights"))

			Expect(body).To(ContainSubstring("\"projects\":[{\"id\""))
		})
	})
})
