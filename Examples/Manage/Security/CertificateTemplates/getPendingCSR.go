package main

import (
	"github.com/mwarnes/marklogic-go"
	"github.com/mwarnes/marklogic-go/Structures"
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

	operation := Structures.CertificateTemplateOperation{
		Operation:  "get-pending-certificate-request",
		CommonName: "ml-node-2",
	}

	content, resp := c.Security.PerformCertificateTemplateOperation("ssl1", operation)

	if resp.StatusCode == 200 {
		log.Println("Operation succeeded.")
		log.Println(content)
	} else {
		log.Println(resp.StatusCode)
	}

}
