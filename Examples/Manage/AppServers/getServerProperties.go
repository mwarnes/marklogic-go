package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/mwarnes/marklogic-go"
	"github.com/mwarnes/marklogic-go/Structures"
	"log"
)

func main() {

	conn := marklogic.Connection{
		Host:               "mwca",
		Port:               8002,
		Username:           "admin",
		Password:           "adminx",
		AuthenticationType: marklogic.DigestAuth,
	}

	c := marklogic.RestClient(conn)

	parms := Structures.ServerParameters{
		GroupId: "Default",
	}

	appSeverPropertiesResponse, _, resp := c.RestService.GetAppServerProperties("xdbc1", parms)

	if resp.StatusCode == 200 {
		log.Println(spew.Sdump(appSeverPropertiesResponse))
	} else {
		log.Println(resp.Status)
		//log.Println(spew.Sdump(errorResponse))
	}

}
