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
	}

	restartResponse, errorResponse, resp := c.AppServer.DeleteAppServer("myAppSrvr", parms)

	if resp.StatusCode == 202 {
		log.Println("Server deleted successfully, restart required.")
		log.Println(spew.Sdump(restartResponse))
	} else if resp.StatusCode == 204 {
		log.Println("Server deleted successfully.")
	} else {
		log.Println(resp.Status)
		log.Println(spew.Sdump(errorResponse))
	}

}
