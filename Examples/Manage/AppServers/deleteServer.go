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

	parms := Structures.ServerParameters{
		GroupId: "Default",
	}

	restartResponse, errorResponse, resp := c.RestService.DeleteAppServer("myAppSrvr", parms)

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
