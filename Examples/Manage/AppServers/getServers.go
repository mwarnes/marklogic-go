package main

import (
	"github.com/mwarnes/marklogic-go"
	"log"

	"marklogic-go"

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
		View:    "default",
	}

	appSeversResponse, errorResponse, resp := c.AppServer.GetAppServers(parms)

	if resp.StatusCode == 200 {
		log.Println(spew.Sdump(appSeversResponse))
	} else {
		log.Println(resp.Status)
		log.Println(spew.Sdump(errorResponse))
	}

}
