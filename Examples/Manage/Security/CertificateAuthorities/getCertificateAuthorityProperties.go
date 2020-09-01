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
		Host:               "ml-node-1",
		Port:               8002,
		Username:           "admin",
		Password:           "admin",
		AuthenticationType: marklogic.DigestAuth,
	}

	// Create a new MarkLogic REST API client
	c := marklogic.MarkLogicManageClient(conn)

	// Get List of Certificate Authorities
	certificateAuthorities, restError, resp := c.Security.GetCertificateAuthorities()
	if resp.StatusCode != 200 {
		log.Fatalln(resp.Status)
	}

	// Get properties for first Certificate Authority in list
	certificateAuthorityProperties, restError, resp := c.Security.GetCertificateAuthorityProperties(
		certificateAuthorities.ListItems.ListItem[0].Idref)

	if resp.StatusCode == 200 {
		log.Println(spew.Sdump(certificateAuthorityProperties))
	} else {
		log.Println(spew.Sdump(restError))
	}

}
