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

	c := marklogic.MarkLogicManageClient(conn)

	server := Structures.AppServerProperties{
		ServerName:      "myAppSrvr",
		GroupName:       "Default",
		ServerType:      "http",
		Root:            "/",
		Port:            8031,
		ContentDatabase: "Documents",
	}

	restartResponse, errorResponse, resp := c.AppServer.UpdateAppServer(server)

	if resp.StatusCode == 204 {
		log.Println("Server updated successfully.")
	} else if resp.StatusCode == 202 {
		log.Println("Server updated successfully, restart required.")
		log.Println(spew.Sdump(restartResponse))
	} else {
		log.Println(resp.Status)
		log.Println(spew.Sdump(errorResponse))
	}

}
