package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/mwarnes/marklogic-go"
	"github.com/mwarnes/marklogic-go/Structures"
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

	externalSecurity := Structures.ExternalSecurityConfigurationParameters{
		ExternalSecurityName: "Auth1",
		Description:          "Updated by go",
		CacheTimeout:         600,
	}

	restartResponse, errorResponse, resp := c.Security.UpdateExternalSecurityConfiguration(externalSecurity)

	if resp.StatusCode == 204 {
		log.Println("External Security configuration updated successfully.")
	} else if resp.StatusCode == 202 {
		log.Println("External Security configuration updated successfully, restart required.")
		log.Println(spew.Sdump(restartResponse))
	} else {
		log.Println(resp.Status)
		log.Println(spew.Sdump(errorResponse))
	}

}
