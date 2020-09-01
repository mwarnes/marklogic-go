package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/mwarnes/marklogic-go"
	"log"
)

func main() {

	conn := marklogic.Connection{
		Host:               "localhost",
		Port:               8002,
		Username:           "admin",
		Password:           "admin",
		AuthenticationType: marklogic.DigestAuth,
	}

	c := marklogic.MarkLogicManageClient(conn)

	certificate, errorResponse, resp := c.Security.GetCertificate("Entrust.net")

	if resp.StatusCode == 200 {
		log.Println(spew.Sdump(certificate))
	} else {
		log.Println(resp.Status)
		log.Println(spew.Sdump(errorResponse))
	}

}
