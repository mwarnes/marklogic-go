package main

import (
	"github.com/mwarnes/marklogic-go"
	"github.com/mwarnes/marklogic-go/Structures"
	"log"
)

func main() {

	conn := marklogic.Connection{
		Host:               "mwca",
		Port:               8000,
		Username:           "admin",
		Password:           "admin",
		AuthenticationType: marklogic.DigestAuth,
	}

	c := marklogic.RestClient(conn)

	documentProperties := Structures.DocumentProperties{
		URI: "/recipe5v2.json",
	}

	httpResp := c.RestService.Delete(documentProperties)

	if httpResp.StatusCode == 204 {
		log.Println("Document deleted.")
	} else {
		log.Println(httpResp.Status)
	}

}
