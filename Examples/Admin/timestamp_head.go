package main

import (
	"fmt"
	"github.com/mwarnes/marklogic-go"
)

// Issue a Timestamp HEAD Request
// https://docs.marklogic.com/REST/HEAD/admin/v1/timestamp
func main() {

	// MarkLogic REST API Connection parameters
	conn := marklogic.Connection{
		Host:               "ml-node-1",
		Port:               8001,
		Username:           "admin",
		Password:           "admin",
		AuthenticationType: marklogic.None,
	}

	// Create a new MarkLogic Admin REST API client
	c := marklogic.MarkLogicAdminClient(conn)

	// Issue Timestamp request
	timestamp, _ := c.Admin.TimestampHead()

	fmt.Println(timestamp)
}
