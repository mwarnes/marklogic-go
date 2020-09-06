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
		Password:           "admin",
		AuthenticationType: marklogic.DigestAuth,
	}

	c := marklogic.MarkLogicRestClient(conn)

	parms := Structures.ServerParameters{
		GroupId: "Default",
	}

	appSeverResponse, errorResponse, resp := c.RestService.GetAppServer("App-Services", parms)

	if resp.StatusCode == 200 {
		log.Println(spew.Sdump(appSeverResponse))
	} else {
		log.Println(resp.Status)
		log.Println(spew.Sdump(errorResponse))
	}

}
