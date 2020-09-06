package main

import (
	"bytes"
	"fmt"
	"github.com/mwarnes/marklogic-go"
	"github.com/mwarnes/marklogic-go/Structures"
	"log"
)

var (
	ucXML = `<?xml  version="1.0" encoding="UTF-8"?>
<article>
<documentSource>Nexus</documentSource>
<isActive>true</isActive>
<description>document</description>`
)

// Issue a init POST Request
// https://docs.marklogic.com/REST/POST/admin/v1/init
func main() {

	//MarkLogic REST API Connection parameters
	conn := marklogic.Connection{
		Host:               "ml-node-1",
		Port:               8000,
		Username:           "admin",
		Password:           "admin",
		AuthenticationType: marklogic.DigestAuth,
	}

	// Create a new MarkLogic Admin REST API client
	c := marklogic.MarkLogicRestClient(conn)

	for i := 29000; i < 100000; i++ {

		documentProperties := Structures.DocumentProperties{
			URI: fmt.Sprintf("%s%d.%s", "/xml/document-", i, ".xml"),
		}

		var buffer bytes.Buffer
		for b := 0; b < 100; b++ {
			buffer.WriteString(randomdata.Paragraph())
		}
		xml := fmt.Sprintf("%s%s%s", ucXML, fmt.Sprintf("%s%s%s", "<para>", buffer.String(), "</para>"), "</article>")

		restApiResp, httpResp, err := c.RestService.Write(documentProperties, bytes.NewBufferString(xml))

		if httpResp.StatusCode == 201 {
			log.Println("Document created. ", i)
		} else if httpResp.StatusCode == 204 {
			log.Println("Document updated.")
		} else {
			log.Println(httpResp.Request.RequestURI, httpResp.Status, err, restApiResp)
		}
	}

}
