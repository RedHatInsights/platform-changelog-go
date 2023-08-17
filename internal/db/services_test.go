package db_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redhatinsights/platform-changelog-go/internal/config"
	"github.com/redhatinsights/platform-changelog-go/internal/logging"
)

var _ = Describe("Handler", func() {

	logging.InitLogger()

	Describe("Create and modify a service", func() {
		It("", func() {
			// create a mock db connection & endpoint handler
			db := testDBImpl

			db.CreateServiceTableEntry("test-service", config.Service{
				DisplayName: "Test Service",
				Tenant:      "test-tenant",
				GHRepo:      "https://github.com/testOrg/test-repo",
				Branch:      "test-branch",
				Namespace:   "test-namespace",
			})

			// Get service by name
			service, rowsAffected, err := db.GetServiceByName("test-service")
			Expect(err).To(BeNil())
			Expect(rowsAffected).To(Equal(int64(1)))

			Expect(service.Name).To(Equal("test-service"))
			Expect(service.DisplayName).To(Equal("Test Service"))
			Expect(service.Tenant).To(Equal("test-tenant"))
			Expect(service.GHRepo).To(Equal("https://github.com/testOrg/test-repo"))
			Expect(service.Branch).To(Equal("test-branch"))
			Expect(service.Namespace).To(Equal("test-namespace"))
			Expect(service.DeployFile).To(Equal(""))
			Expect(service.GLRepo).To(Equal(""))

			// update the service
			updated_service, err := db.UpdateServiceTableEntry("test-service", config.Service{
				DisplayName: "Test Service",
				Tenant:      "test-tenant",
				GHRepo:      "https://github.com/testOrg/test-repo",
				Branch:      "test-branch",
				Namespace:   "test-namespace",
				DeployFile:  "test-deploy-file", // new
			})

			Expect(err).To(BeNil())
			Expect(updated_service.DeployFile).To(Equal("test-deploy-file"))

			// just going to make sure the change persisted
			updated_service_1, rowsAffected, err := db.GetServiceByName("test-service")
			Expect(err).To(BeNil())
			Expect(rowsAffected).To(Equal(int64(1)))
			Expect(updated_service_1.DeployFile).To(Equal("test-deploy-file"))

			// delete the service
			deleted_service_struct, err := db.DeleteServiceTableEntry("test-service")
			Expect(err).To(BeNil())
			Expect(deleted_service_struct.Name).To(Equal("test-service"))

			// make sure the service is gone
			deleted_service, rowsAffected, _ := db.GetServiceByName("test-service")
			// Expect(err).To(BeNil())
			Expect(rowsAffected).To(Equal(int64(0)))
			Expect(deleted_service.Name).To(Equal(""))
		})
	})
})
