package main

import (
	"github.com/mwarnes/marklogic-go"
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

	documentProperties := marklogic.DocumentProperties{
		URI: "/recipe2.json",
	}

	httpResp := c.Document.Delete(documentProperties)

	if httpResp.StatusCode == 204 {
		log.Println("Document deleted.")
	} else {
		log.Println(httpResp.Status)
	}

}
