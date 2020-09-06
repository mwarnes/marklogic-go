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
		Description:          "This is an external auth created by go",
		CacheTimeout:         300,
		Authentication:       "kerberos",
		Authorization:        "internal",
	}

	errorResponse, resp := c.Security.AddExternalSecurityConfiguration(externalSecurity)

	if resp.StatusCode == 200 {
		log.Println("External Security configuration added successfully.")
	} else if resp.StatusCode == 201 {
		log.Println("External Security configuration added successfully, restart required.")
	} else {
		log.Println(resp.Status)
		log.Println(spew.Sdump(errorResponse))
	}

}
