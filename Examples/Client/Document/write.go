package main

import (
	"bytes"
	"crypto/tls"
	"log"

	"marklogic-go"
)

func main() {

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	// MarkLogic REST API Connection parameters
	conn := marklogic.Connection{
		Host: "mwca",
		Port: 8000,
		// Username:           "admin",
		// Password:           "admin",
		AuthenticationType: marklogic.KerberosAuth,
		TLSConfig:          tlsConfig,
	}

	// Create a new MarkLogic Admin REST API client
	c := marklogic.MarkLogicRestClient(conn)

	s := `{"recipe":"Apple pie", "fromScratch":true, "ingredients":"apples"}`

	documentProperties := marklogic.DocumentProperties{
		URI: "/recipe5.json",
	}

	// Initialize MarkLogic server (With or without a license)
	restAPIResp, httpResp, err := c.Document.Write(documentProperties, bytes.NewBufferString(s))

	if httpResp.StatusCode == 201 {
		log.Println("Document created.")
	} else if httpResp.StatusCode == 204 {
		log.Println("Document updated.")
	} else {
		log.Println(httpResp.Request.RequestURI, httpResp.Status, err, restAPIResp)
	}

}
