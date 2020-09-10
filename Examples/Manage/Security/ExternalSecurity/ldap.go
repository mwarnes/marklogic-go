package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/mwarnes/marklogic-go"
	"github.com/mwarnes/marklogic-go/Structures"
	"log"
)

func main() {

	conn := marklogic.Connection{
		Host:               "ml-node-1",
		Port:               8002,
		Username:           "admin",
		Password:           "admin",
		AuthenticationType: marklogic.DigestAuth,
	}

	c := marklogic.RestClient(conn)

	externalSecurity := Structures.ExternalSecurityConfigurationParameters{
		ExternalSecurityName: "LDAP1",
		Description:          "This is an external auth created by go",
		CacheTimeout:         300,
		Authentication:       "ldap",
		Authorization:        "internal",
		LdapServerURI:        "http://localhost:389",
		LdapBase:             "cn=base",
		LdapBindMethod:       "simple",
		LdapAttribute:        "uid",
		LdapDefaultUser:      "admin",
		LdapPassword:         "password",
	}

	errorResponse, resp := c.RestService.AddExternalSecurityConfiguration(externalSecurity)

	if resp.StatusCode == 200 {
		log.Println("External Security configuration added successfully.")
	} else if resp.StatusCode == 201 {
		log.Println("External Security configuration added successfully, restart required.")
	} else {
		log.Println(resp.Status)
		log.Println(spew.Sdump(errorResponse))
	}

}
