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
		AuthenticationType: marklogic.DigestAuth,
	}

	// Create a new MarkLogic Admin REST API client
	c := marklogic.MarkLogicRestClient(conn)

	// Issue Timestamp request
	timestamp, _ := c.RestService.TimestampHead()

	fmt.Println(timestamp)
}
