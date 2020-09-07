package main

import (
	"bytes"
	"github.com/mwarnes/marklogic-go/Structures"
	"log"

	"github.com/mwarnes/marklogic-go"
)

func main() {

	//tlsConfig := &tls.Config{
	//	InsecureSkipVerify: true,
	//}

	// MarkLogic REST API Connection parameters
	conn := marklogic.Connection{
		Host:               "mwca",
		Port:               8000,
		Username:           "admin",
		Password:           "admin",
		AuthenticationType: marklogic.DigestAuth,
		//TLSConfig:          tlsConfig,
	}

	// Create a new MarkLogic Admin REST API client
	c := marklogic.RestClient(conn)

	s := `{"recipe":"Apple pie", "fromScratch":true, "ingredients":"apples"}`

	documentProperties := Structures.DocumentProperties{
		URI: "/recipe5v2.json",
	}

	// Initialize MarkLogic server (With or without a license)
	restAPIResp, httpResp, err := c.RestService.Write(documentProperties, bytes.NewBufferString(s))

	if httpResp.StatusCode == 201 {
		log.Println("Document created.")
	} else if httpResp.StatusCode == 204 {
		log.Println("Document updated.")
	} else {
		log.Println(httpResp.Request.RequestURI, httpResp.Status, err, restAPIResp)
	}

}
