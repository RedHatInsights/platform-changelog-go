package endpoints_test

import (
	"fmt"
	"testing"

	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redhatinsights/platform-changelog-go/internal/config"
	"github.com/redhatinsights/platform-changelog-go/internal/db"
	"github.com/redhatinsights/platform-changelog-go/internal/models"
	"github.com/redhatinsights/platform-changelog-go/internal/utils"
	"gorm.io/gorm"
)

var (
	testDB     *embeddedpostgres.EmbeddedPostgres
	testGormDB *gorm.DB
	testDBImpl *db.DBConnectorImpl
)

func TestEndpoints(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Endpoints Suite")
}

var _ = BeforeSuite(func() {
	cfg := config.Get()

	var err error
	testDB, testGormDB, err = utils.CreateTestDB(cfg, "file://../../migrations")
	if err != nil {
		testDB.Stop()
	}

	Expect(err).To(BeNil())

	testDBImpl = db.SetDBConnector(testGormDB)

	seedServicesAndProjects(testDBImpl)
})

var _ = AfterSuite(func() {
	err := testDB.Stop()
	Expect(err).To(BeNil())
	fmt.Println("TEST DB STOPPED")
})

func seedServicesAndProjects(db *db.DBConnectorImpl) {
	ms := []models.Services{
		{ID: 1, Name: "platform-changelog", DisplayName: "Platform Changelog", Tenant: "Insights"},
		{ID: 2, Name: "insights-ingress", DisplayName: "Insights Ingress", Tenant: "Insights"},
		{ID: 3, Name: "rbac", DisplayName: "Insights RBAC", Tenant: "Insights"},
		{ID: 4, Name: "chrome-service", DisplayName: "Chrome Service", Tenant: "Insights"},
	}

	mp := []models.Projects{
		{
			ID:         1,
			ServiceID:  1,
			Name:       "platform-changelog-go",
			Repo:       "https://github.com/RedhatInsights/platform-changelog-go",
			Namespaces: []string{},
			Branches:   []string{},
		},
		{
			ID:         2,
			ServiceID:  1,
			Name:       "platform-changelog-ui",
			Repo:       "https://github.com/RedhatInsights/platform-changelog-ui",
			Namespaces: []string{},
			Branches:   []string{},
		},
		{
			ID:         3,
			ServiceID:  2,
			Name:       "insights-ingress-go",
			Repo:       "https://github.com/RedhatInsights/insights-ingress-go",
			Namespaces: []string{},
			Branches:   []string{},
		},
		{
			ID:         4,
			ServiceID:  3,
			Name:       "insights-rbac",
			Repo:       "https://github.com/RedhatInsights/insights-rbac",
			Namespaces: []string{},
			Branches:   []string{},
		},
	}

	CreateServices(db, ms)
	CreateProjects(db, mp)
}

func CreateServices(conn db.DBConnector, services []models.Services) {
	for _, s := range services {
		fmt.Printf(s.Name)
		_, err := conn.CreateServiceTableEntry(s)
		Expect(err).To(BeNil())
	}
}

func CreateProjects(conn db.DBConnector, projects []models.Projects) {
	for _, p := range projects {
		err := conn.CreateProjectTableEntry(p)
		Expect(err).To(BeNil())
	}
}
