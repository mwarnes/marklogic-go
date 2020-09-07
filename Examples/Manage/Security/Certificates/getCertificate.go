package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/mwarnes/marklogic-go"
	"log"
)

func main() {

	conn := marklogic.Connection{
		Host:               "mwca",
		Port:               8002,
		Username:           "admin",
		Password:           "admin",
		AuthenticationType: marklogic.DigestAuth,
	}

	c := marklogic.RestClient(conn)

	certificate, errorResponse, resp := c.RestService.GetCertificate("Entrust.net")

	if resp.StatusCode == 200 {
		log.Println(spew.Sdump(certificate))
	} else {
		log.Println(resp.Status)
		log.Println(spew.Sdump(errorResponse))
	}

}
