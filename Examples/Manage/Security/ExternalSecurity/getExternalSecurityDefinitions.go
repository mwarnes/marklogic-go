package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/mwarnes/marklogic-go"
	"log"
)

func main() {

	// MarkLogic REST API Connection parameters
	conn := marklogic.Connection{
		Host:               "localhost",
		Port:               8002,
		Username:           "admin",
		Password:           "admin",
		AuthenticationType: marklogic.DigestAuth,
	}

	// Create a new MarkLogic REST API client
	c := marklogic.MarkLogicRestClient(conn)

	// Get a list of all External Security Definitions stored in MarkLogic
	externalSecurityList, restError, resp := c.RestService.GetExternalSecurity()

	if resp.StatusCode == 200 {
		log.Println(spew.Sdump(externalSecurityList))
	} else {
		log.Println(spew.Sdump(restError))
	}

}
