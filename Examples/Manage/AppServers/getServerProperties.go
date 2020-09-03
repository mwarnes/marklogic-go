package main

import (
	"github.com/mwarnes/marklogic-go"
	"log"

	"github.com/davecgh/go-spew/spew"
)

func main() {

	conn := marklogic.Connection{
		Host:               "mwca",
		Port:               8002,
		Username:           "admin",
		Password:           "admin",
		AuthenticationType: marklogic.DigestAuth,
	}

	c := marklogic.MarkLogicManageClient(conn)

	parms := marklogic.ServerParameters{
		GroupId: "Default",
	}

	appSeverPropertiesResponse, errorResponse, resp := c.AppServer.GetAppServerProperties("xdbc1", parms)

	if resp.StatusCode == 200 {
		log.Println(spew.Sdump(appSeverPropertiesResponse))
	} else {
		log.Println(resp.Status)
		log.Println(spew.Sdump(errorResponse))
	}

}
