package endpoints

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/redhatinsights/platform-changelog-go/internal/structs"
)

func writeResponse(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	w.Write([]byte(message))
}

func initQuery(r *http.Request) (structs.Query, error) {
	q := structs.Query{
		Offset: 0,
		Limit:  10,

		StartDate: r.URL.Query().Get("start_date"),
		EndDate:   r.URL.Query().Get("end_date"),
	}

	// allowing multiple values for all the keys
	values := r.URL.Query()
	for k, v := range values {
		k = strings.ToLower(k)
		switch k {
		// timeline filters
		case "ref":
			q.Ref = v
		case "repo":
			q.Repo = v
		case "author":
			q.Author = v
		case "merged_by":
			q.MergedBy = v
		case "cluster":
			q.Cluster = v
		case "image":
			q.Image = v
		// service filters
		case "name":
			q.ServiceName = v
		case "display_name":
			q.ServiceDisplayName = v
		case "tenant":
			q.ServiceTenant = v
		case "namespace":
			q.ServiceNamespace = v
		case "branch":
			q.ServiceBranch = v
		}
	}

	var err error

	offset := r.URL.Query().Get("offset")
	limit := r.URL.Query().Get("limit")

	if offset != "" {
		q.Offset, err = strconv.Atoi(offset)
	}

	if limit != "" {
		q.Limit, err = strconv.Atoi(limit)
	}

	return q, err
}
