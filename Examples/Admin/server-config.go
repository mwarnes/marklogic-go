package main

import (
	"github.com/mwarnes/marklogic-go"
	"log"
)

// Issue a instance-admin POST Request
// https://docs.marklogic.com/REST/GET/admin/v1/server-config
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

	// Get MarkLogic Server configuration.
	configuration, resp := c.RestService.GetServerConfig()

	if resp.StatusCode == 200 {
		log.Println("Successfully retrieved MarkLogic server configuration.")
		log.Println(configuration)
	} else {
		log.Println("Error retrieving configuration.")
	}

}
