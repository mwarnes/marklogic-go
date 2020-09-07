package main

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/mwarnes/marklogic-go"
	"github.com/mwarnes/marklogic-go/Structures"
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

	operation := Structures.CertificateTemplateOperation{
		Operation: "get-certificates-for-template",
	}

	content, resp := c.RestService.PerformCertificateTemplateOperation("ssl1", operation)

	if resp.StatusCode == 200 {
		var templateCertificates Structures.TemplateCertificates
		err := json.Unmarshal(content, &templateCertificates)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(spew.Sdump(templateCertificates))
	} else {
		log.Println(resp.StatusCode)
	}

}
