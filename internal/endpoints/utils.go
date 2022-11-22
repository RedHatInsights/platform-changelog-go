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

		Start_Date: r.URL.Query().Get("start_date"),
		End_Date:   r.URL.Query().Get("end_date"),
	}

	// allowing multiple values for all the keys
	values := r.URL.Query()
	for k, v := range values {
		k = strings.ToLower(k)

		// timeline filters
		if k == "ref" {
			q.Ref = v
		} else if k == "repo" {
			q.Repo = v
		} else if k == "author" {
			q.Author = v
		} else if k == "merged_by" {
			q.Merged_By = v
		} else if k == "cluster" {
			q.Cluster = v
		} else if k == "image" {
			q.Image = v
		}

		// service filters
		if k == "name" {
			q.Service_Name = v
		} else if k == "display_name" {
			q.Service_Display_Name = v
		} else if k == "tenant" {
			q.Service_Tenant = v
		} else if k == "namespace" {
			q.Service_Namespace = v
		} else if k == "branch" {
			q.Service_Branch = v
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
