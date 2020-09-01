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

	server := marklogic.AppServerProperties{
		ServerName:      "myAppSrvr",
		GroupName:       "Default",
		ServerType:      "http",
		Root:            "/",
		Port:            8030,
		ContentDatabase: "Documents",
	}

	errorResponse, resp := c.AppServer.AddAppServer(server)

	if resp.StatusCode == 200 {
		log.Println("Server added successfully.")
	} else if resp.StatusCode == 201 {
		log.Println("Server added successfully, restart required.")
	} else {
		log.Println(resp.Status)
		log.Println(spew.Sdump(errorResponse))
	}

}
