package main

import (
	"github.com/mwarnes/marklogic-go"
	"github.com/mwarnes/marklogic-go/Structures"
	"log"
)

func main() {

	conn := marklogic.Connection{
		Host:               "ml-node-1",
		Port:               8000,
		Username:           "admin",
		Password:           "admin",
		AuthenticationType: marklogic.DigestAuth,
	}

	c := marklogic.MarkLogicRestClient(conn)

	documentProperties := Structures.DocumentProperties{
		URI: "/recipe2.json",
	}

	document, httpResp := c.RestService.Read(documentProperties)

	if httpResp.StatusCode == 200 {
		log.Println("Document read.")
		log.Println(document)
	} else {
		log.Println(httpResp.Status)
	}

}
