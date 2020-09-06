package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/mwarnes/marklogic-go"
	"github.com/mwarnes/marklogic-go/Structures"
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

	parms := Structures.ExternalSecurityParameters{
		Format: "json",
	}

	// Get a list of all External Security Definitions stored in MarkLogic
	externalSecurityProperties, restError, resp := c.RestService.GetExternalSecurityProperties("KerberosExtSec", parms)

	if resp.StatusCode == 200 {
		log.Println(spew.Sdump(externalSecurityProperties))
	} else {
		log.Println(spew.Sdump(restError))
	}

}
