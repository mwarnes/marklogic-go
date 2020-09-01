package main

import (
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
	}

	appSeverResponse, errorResponse, resp := c.AppServer.GetAppServer("App-Services", parms)

	if resp.StatusCode == 200 {
		log.Println(spew.Sdump(appSeverResponse))
	} else {
		log.Println(resp.Status)
		log.Println(spew.Sdump(errorResponse))
	}

}
