package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/mwarnes/marklogic-go"
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

	c := marklogic.MarkLogicManageClient(conn)

	parms := marklogic.ServerParameters{
		GroupId: "Default",
	}

	appSeverPropertiesResponse, errorResponse, resp := c.AppServer.GetAppServerProperties("myAppSrvr", parms)

	if resp.StatusCode == 200 {
		log.Println(spew.Sdump(appSeverPropertiesResponse))
	} else {
		log.Println(resp.Status)
		log.Println(spew.Sdump(errorResponse))
	}

}
