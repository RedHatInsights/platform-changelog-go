package endpoints

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/redhatinsights/platform-changelog-go/internal/config"
	"gopkg.in/yaml.v2"
)

func (eh *EndpointHandler) OpenAPIHandler(cfg *config.Config) http.HandlerFunc {

	openAPISpec := readOpenAPISpec(cfg.OpenAPIPath)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(openAPISpec)
	})
}

func readOpenAPISpec(path string) []byte {
	var openAPISpec config.OpenAPISpec
	openAPISpecFile, err := os.Open("schema/openapi.yaml")
	if err != nil {
		panic(err)
	}
	defer openAPISpecFile.Close()
	decoder := yaml.NewDecoder(openAPISpecFile)
	err = decoder.Decode(&openAPISpec)
	if err != nil {
		panic(err)
	}
	openAPISpecJSON, err := json.Marshal(openAPISpec)
	if err != nil {
		panic(err)
	}
	return openAPISpecJSON
}
