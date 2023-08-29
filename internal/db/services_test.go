package db_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redhatinsights/platform-changelog-go/internal/config"
	"github.com/redhatinsights/platform-changelog-go/internal/logging"
	"github.com/redhatinsights/platform-changelog-go/internal/models"
)

var _ = Describe("Handler", func() {

	logging.InitLogger()

	Describe("Create and modify a service", func() {
		It("", func() {
			db := testDBImpl

			db.CreateServiceTableEntry(models.Services{
				ID:          1,
				Name:        "test-service",
				DisplayName: "Test Service",
				Tenant:      "test-tenant",
			})

			// Get service by name
			service, rowsAffected, err := db.GetServiceByName("test-service")
			Expect(err).To(BeNil())
			Expect(rowsAffected).To(Equal(int64(1)))

			Expect(service.Name).To(Equal("test-service"))
			Expect(service.DisplayName).To(Equal("Test Service"))
			Expect(service.Tenant).To(Equal("test-tenant"))

			// update the service
			updated_service, err := db.UpdateServiceTableEntry("test-service", config.Service{
				DisplayName: "Test Service",
				Tenant:      "test-tenant-new",
			})

			Expect(err).To(BeNil())
			Expect(updated_service.Tenant).To(Equal("test-tenant-new"))

			// just going to make sure the change persisted
			updated_service_1, rowsAffected, err := db.GetServiceByName("test-service")
			Expect(err).To(BeNil())
			Expect(rowsAffected).To(Equal(int64(1)))
			Expect(updated_service_1.Tenant).To(Equal("test-tenant-new"))

			// delete the service
			deleted_service_struct, err := db.DeleteServiceTableEntry("test-service")
			Expect(err).To(BeNil())
			Expect(deleted_service_struct.Name).To(Equal("test-service"))

			// make sure the service is gone
			deleted_service, rowsAffected, _ := db.GetServiceByName("test-service")
			Expect(rowsAffected).To(Equal(int64(0)))
			Expect(deleted_service.Name).To(Equal(""))
		})
	})
})
