package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/mwarnes/marklogic-go"
	"log"
)

func main() {

	// MarkLogic REST API Connection parameters
	conn := marklogic.Connection{
		Host:               "ml-node-1",
		Port:               8002,
		Username:           "admin",
		Password:           "admin",
		AuthenticationType: marklogic.DigestAuth,
	}

	// Create a new MarkLogic REST API client
	c := marklogic.RestClient(conn)

	// Get a list of all External Security Definitions stored in MarkLogic
	externalSecurityConfiguration, restError, resp := c.RestService.GetExternalSecurityConfiguration("SAML1")

	if resp.StatusCode == 200 {
		log.Println(spew.Sdump(externalSecurityConfiguration))
	} else {
		log.Println(spew.Sdump(restError))
	}

}
