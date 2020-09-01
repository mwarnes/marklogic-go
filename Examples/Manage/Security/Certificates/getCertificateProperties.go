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

	certificates, restError, resp := c.Security.GetCertificates()
	if resp.StatusCode != 200 {
		log.Fatalln(resp.Status)
	}

	certificateProperties, restError, resp := c.Security.GetCertificateProperties(
		certificates.ListItems.ListItem[0].Idref)

	if resp.StatusCode == 200 {
		log.Println(spew.Sdump(certificateProperties))
	} else {
		log.Println(spew.Sdump(restError))
	}

}
