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

	c := marklogic.RestClient(conn)

	server := Structures.AppServerProperties{
		ServerName:      "myAppSrvr",
		GroupName:       "Default",
		ServerType:      "http",
		Root:            "/",
		Port:            8030,
		ContentDatabase: "Documents",
	}

	errorResponse, resp := c.RestService.AddAppServer(server)

	if resp.StatusCode == 200 {
		log.Println("Server added successfully.")
	} else if resp.StatusCode == 201 {
		log.Println("Server added successfully, restart required.")
	} else {
		log.Println(resp.Status)
		log.Println(spew.Sdump(errorResponse))
	}

}
