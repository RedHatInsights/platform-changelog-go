package endpoints_test

import (
	"net/http"
	"net/http/httptest"

	chi "github.com/go-chi/chi/v5"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("URL tests", func(url string, expectedStatusCode int) {

	router := chi.NewRouter()
	router.Get("/api/v1/services", func(w http.ResponseWriter, router *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// create a request
	req, _ := http.NewRequest("GET", url, nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	Expect(rr.Code).To(Equal(expectedStatusCode))

},
	Entry("Valid URL", "/api/v1/services", http.StatusOK),
	Entry("URL with extra '/'", "/api/v1/services/", http.StatusOK),
	Entry("Nonexistent URL", "/api/v1/ser", http.StatusNotFound),
)
