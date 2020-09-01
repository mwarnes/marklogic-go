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

	parms := marklogic.ServerParameters{
		GroupId: "Default",
		View:    "package",
	}

	appSeversResponse, errorResponse, resp := c.AppServer.GetAppServers(parms)

	if resp.StatusCode == 200 {
		log.Println(spew.Sdump(appSeversResponse))
	} else {
		log.Println(resp.Status)
		log.Println(spew.Sdump(errorResponse))
	}

}
