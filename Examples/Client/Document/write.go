package main

import (
	"bytes"
	"github.com/mwarnes/marklogic-go"
	"log"
)

// Issue a init POST Request
// https://docs.marklogic.com/REST/POST/admin/v1/init
func main() {

	// MarkLogic REST API Connection parameters
	conn := marklogic.Connection{
		Host:               "ml-node-1",
		Port:               8000,
		Username:           "admin",
		Password:           "admin",
		AuthenticationType: marklogic.DigestAuth,
	}

	// Create a new MarkLogic Admin REST API client
	c := marklogic.MarkLogicRestClient(conn)

	s := `{"recipe":"Apple pie", "fromScratch":true, "ingredients":"apples"}`

	documentProperties := marklogic.DocumentProperties{
		URI: "/recipe2.json",
	}

	// Initialize MarkLogic server (With or without a license)
	restApiResp, httpResp, err := c.Document.Write(documentProperties, bytes.NewBufferString(s))

	if httpResp.StatusCode == 201 {
		log.Println("Document created.")
	} else if httpResp.StatusCode == 204 {
		log.Println("Document updated.")
	} else {
		log.Println(httpResp.Request.RequestURI, httpResp.Status, err, restApiResp)
	}

}
