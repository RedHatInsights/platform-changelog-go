package endpoints

import "github.com/redhatinsights/platform-changelog-go/internal/db"

type EndpointHandler struct {
	conn db.DBConnector
}

func NewHandler(conn db.DBConnector) EndpointHandler {
	return EndpointHandler{
		conn: conn,
	}
}
