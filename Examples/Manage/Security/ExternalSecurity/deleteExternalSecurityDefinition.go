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

	restartResponse, errorResponse, resp := c.Security.DeleteExternalSecurityConfiguration("Auth1")

	if resp.StatusCode == 202 {
		log.Println("External Security configuration deleted successfully, restart required.")
		log.Println(spew.Sdump(restartResponse))
	} else if resp.StatusCode == 204 {
		log.Println("External Security configuration deleted successfully.")
	} else {
		log.Println(resp.Status)
		log.Println(spew.Sdump(errorResponse))
	}

}
