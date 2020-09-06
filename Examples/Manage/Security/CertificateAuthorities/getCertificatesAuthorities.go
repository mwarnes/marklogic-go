package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/mwarnes/marklogic-go"
	"log"
)

// Issue a init POST Request
// https://docs.marklogic.com/REST/POST/admin/v1/init
func main() {

	// MarkLogic REST API Connection parameters
	conn := marklogic.Connection{
		Host:               "mwca",
		Port:               8002,
		Username:           "admin",
		Password:           "password",
		AuthenticationType: marklogic.DigestAuth,
	}

	// Create a new MarkLogic REST API client
	c := marklogic.MarkLogicRestClient(conn)

	// Get a list of all Certificates Authorities stored in MarkLogic
	certificateAuthoritiesList, restError, resp := c.RestService.GetCertificateAuthorities()
	log.Println(spew.Sdump(resp.Status))

	if resp.StatusCode == 200 {
		log.Println(spew.Sdump(certificateAuthoritiesList))
	} else {
		log.Println(spew.Sdump(restError))
	}

}
