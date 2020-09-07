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
		Password:           "admin",
		AuthenticationType: marklogic.DigestAuth,
	}

	// Create a new MarkLogic REST API client
	c := marklogic.RestClient(conn)

	certificateAuthority, restError, resp := c.RestService.GetCertificateAuthority("14539353274846980465")

	if resp.StatusCode == 200 {
		log.Println(spew.Sdump(certificateAuthority))
	} else {
		log.Println(spew.Sdump(restError))
	}

}
