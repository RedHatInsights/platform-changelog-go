package db_test

import (
	"github.com/redhatinsights/platform-changelog-go/internal/logging"
	"github.com/redhatinsights/platform-changelog-go/internal/models"
	"github.com/redhatinsights/platform-changelog-go/internal/structs"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	db_package "github.com/redhatinsights/platform-changelog-go/internal/db"
)

var (
	s models.Services
)

var _ = Describe("Handler", Ordered, func() {

	logging.InitLogger()

	BeforeAll(func() {
		db := testDBImpl

		// create a service entry for the project
		s = models.Services{
			Name:        "rhel",
			DisplayName: "RHEL",
			Tenant:      "rhel",
		}
		err := db.CreateServiceTableEntry(&s)
		Expect(s.ID).NotTo(Equal(0))
		Expect(err).To(BeNil())

		// create a project
		p := models.Projects{
			ServiceID: s.ID,
			Name:      "test-project",
			Repo:      "github.com/org/test-project",
			Namespace: "",
			Branch:    "origin/main",
		}

		err = db.CreateProjectTableEntry(&p)
		Expect(p.ID).NotTo(Equal(0))
		Expect(err).To(BeNil())
	})

	Describe("Manipulate project", func() {
		It("Get project by repo", func() {
			db := testDBImpl

			project, err := db.GetProjectByRepo("github.com/org/test-project")
			Expect(err).To(BeNil())

			Expect(project.ID).ToNot(Equal(0))
			Expect(project.ServiceID).To(Equal(s.ID))
			Expect(project.Name).To(Equal("test-project"))
			Expect(project.Repo).To(Equal("github.com/org/test-project"))
			Expect(project.DeployFile).To(Equal(""))
			Expect(project.Namespace).To(Equal(""))
			Expect(project.Branch).To(Equal("origin/main"))
		})

		It("Update project", func() {
			db := testDBImpl

			project, rowsAffected, err := db.GetProjectByName("test-project")
			Expect(err).To(BeNil())
			Expect(rowsAffected).To(Equal(int64(1)))
			Expect(project.ID).ToNot(Equal(0))

			project.Namespace = "insights-stage"
			project.DeployFile = "github.com/org/test-project/deploy.yml"

			err = db.UpdateProjectTableEntry(&project)
			Expect(err).To(BeNil())
			Expect(project.ID).ToNot(Equal(0))
			Expect(project.ServiceID).To(Equal(s.ID))
			Expect(project.Name).To(Equal("test-project"))
			Expect(project.Repo).To(Equal("github.com/org/test-project"))
			Expect(project.Namespace).To(Equal("insights-stage"))
			Expect(project.DeployFile).To(Equal("github.com/org/test-project/deploy.yml"))

			updated_project, rowsAffected, err := db.GetProjectByName("test-project")
			Expect(err).To(BeNil())
			Expect(rowsAffected).To(Equal(int64(1)))
			Expect(updated_project.Namespace).To(Equal("insights-stage"))
		})

		It("Get projects by service", func() {
			db := testDBImpl

			projects, rowsAffected, err := db.GetProjectsByService(s, 0, 10, structs.Query{})
			Expect(err).To(BeNil())
			Expect(rowsAffected).To(Equal(int64(1)))

			Expect(projects[0].ID).ToNot(Equal(0))
			Expect(projects[0].ServiceID).To(Equal(s.ID))
			Expect(projects[0].Name).To(Equal("test-project"))
			Expect(projects[0].Repo).To(Equal("github.com/org/test-project"))
			Expect(projects[0].DeployFile).To(Equal("github.com/org/test-project/deploy.yml"))
			Expect(projects[0].Namespace).To(Equal("insights-stage"))
			Expect(projects[0].Branch).To(Equal("origin/main"))

			// Adding another project
			p := models.Projects{
				ServiceID: s.ID,
				Name:      "test-project-2",
				Repo:      "github.com/org/test-project-2",
				Namespace: "",
				Branch:    "origin/main",
			}

			err = db.CreateProjectTableEntry(&p)
			Expect(p.ID).NotTo(Equal(0))
			Expect(err).To(BeNil())

			projects, rowsAffected, err = db.GetProjectsByService(s, 0, 10, structs.Query{})
			Expect(err).To(BeNil())
			Expect(rowsAffected).To(Equal(int64(2)))

			// taking advantage of sorting by ID desc
			Expect(projects[0].ID).ToNot(Equal(0))
			Expect(projects[0].ServiceID).To(Equal(s.ID))
			Expect(projects[0].Name).To(Equal("test-project-2"))
			Expect(projects[0].Repo).To(Equal("github.com/org/test-project-2"))
			Expect(projects[0].DeployFile).To(Equal(""))
			Expect(projects[0].Namespace).To(Equal(""))
			Expect(projects[0].Branch).To(Equal("origin/main"))

			Expect(projects[1].ID).ToNot(Equal(0))
			Expect(projects[1].ServiceID).To(Equal(s.ID))
			Expect(projects[1].Name).To(Equal("test-project"))
			Expect(projects[1].Repo).To(Equal("github.com/org/test-project"))
			Expect(projects[1].DeployFile).To(Equal("github.com/org/test-project/deploy.yml"))
			Expect(projects[1].Namespace).To(Equal("insights-stage"))
			Expect(projects[1].Branch).To(Equal("origin/main"))
		})

		It("Get project by name", func() {
			db := testDBImpl

			project, rowsAffected, err := db.GetProjectByName("test-project")
			Expect(err).To(BeNil())
			Expect(rowsAffected).To(Equal(int64(1)))

			Expect(project.ID).ToNot(Equal(0))
			Expect(project.ServiceID).To(Equal(s.ID))
			Expect(project.Name).To(Equal("test-project"))
			Expect(project.Repo).To(Equal("github.com/org/test-project"))
			Expect(project.Branch).To(Equal("origin/main"))
		})

		It("Duplicate Name", func() {
			db := testDBImpl

			duplicate := models.Projects{
				ServiceID: s.ID,
				Name:      "test-project",
				Repo:      "github.com/org-1/test-project",
				Namespace: "",
				Branch:    "origin/main",
			}

			err := db.CreateProjectTableEntry(&duplicate)
			Expect(err).To(BeNil())

			// try to get by project name
			project, rowsAffected, err := db.GetProjectByName("test-project")
			Expect(err).To(BeNil())
			// because of .First() function, the rowsAffected is either 0 or 1
			Expect(rowsAffected).To(Equal(int64(1)))

			// we need the project name to somehow be an unique identifier
			// maybe requiring the tenant as well? GetProjectByNameAndTenant

			Expect(project.ID).To(Equal(duplicate.ID))
			Expect(project.ServiceID).To(Equal(s.ID))
		})
	})

	Describe("Negative project tests", func() {
		It("Getting non-existent projects", func() {
			db := testDBImpl

			project, rowsAffected, err := db.GetProjectByName("non-existent-project")
			Expect(err).To(Equal(db_package.ErrNotFound))
			Expect(rowsAffected).To(Equal(int64(0)))
			Expect(project.ID).To(Equal(0))

			project, err = db.GetProjectByRepo("github.com/org/non-existent-project")
			Expect(err).To(Equal(db_package.ErrNotFound))
			Expect(rowsAffected).To(Equal(int64(0)))
			Expect(project.ID).To(Equal(0))
		})
	})
})
