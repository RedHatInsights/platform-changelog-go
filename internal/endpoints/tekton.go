package endpoints

import (
	"io/ioutil"
	"net/http"

	l "github.com/redhatinsights/platform-changelog-go/internal/logging"
	"github.com/redhatinsights/platform-changelog-go/internal/metrics"
)

// Since the Tekton task has some unknowns still
// (what data will be available and the format of that data),
// this endpoint will log the request body and return a 200.
func TektonTaskRun(w http.ResponseWriter, r *http.Request) {
	var err error
	var payload []byte

	metrics.IncTekton(r.Method, r.UserAgent(), false)

	// log the request body
	l.Log.Info("Tekton TaskRun request body: ", r.Body)

	// read the info from the request body
	payload, err = ioutil.ReadAll(r.Body)
	if err != nil {
		l.Log.Error(err)
		metrics.IncTekton(r.Method, r.UserAgent(), true)
		return
	}

	defer r.Body.Close()

	// parse the tekton info
	_, err = ParseTektonTaskRun(payload)

	if err != nil {
		l.Log.Error(err)
		metrics.IncTekton(r.Method, r.UserAgent(), true)
		return
	}

	writeResponse(w, http.StatusOK, `{"msg": "Tekton info received"}`)
}

func ParseTektonTaskRun(payload []byte) (interface{}, error) {
	return nil, nil
}
