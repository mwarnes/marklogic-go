package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/mwarnes/marklogic-go"
	"log"
)

// Issue a instance-admin POST Request
// https://docs.marklogic.com/REST/POST/admin/v1/instance-admin
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
	c := marklogic.MarkLogicAdminClient(conn)

	secProps := marklogic.SecurityProperties{
		AdminUsername: "admin",
		AdminPassword: "admin",
		Realm:         "public",
	}

	// Initialize MarkLogic server Security database.
	restartResp, errorResp, resp := c.Admin.InstanceAdmin(secProps)

	if resp.StatusCode == 202 {
		log.Println("Security initialised restart required.")
		log.Println(spew.Sdump(restartResp))
	} else {
		log.Println(resp.Status)
		log.Println(spew.Sdump(errorResp))
	}

}
