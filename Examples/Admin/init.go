package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/mwarnes/marklogic-go"
	"github.com/mwarnes/marklogic-go/Structures"
	"log"
)

// Issue a init POST Request
// https://docs.marklogic.com/REST/POST/admin/v1/init
func main() {

	// MarkLogic REST API Connection parameters
	conn := marklogic.Connection{
		Host:               "ml-node-1",
		Port:               8001,
		Username:           "admin",
		Password:           "admin",
		AuthenticationType: marklogic.DigestAuth,
	}

	// Create a new MarkLogic Admin REST API client
	c := marklogic.RestClient(conn)

	lic := Structures.LicenseProperties{
		LicenseKey: "3189-DCE7-6A33-D93D-787D-0303-B9DD-C076-5433-AD2D-8A65-53BA-13DC-96F4-F5DD-B80",
		Licensee:   "MarkLogic - Martin Warnes",
	}

	//lic := services.LicenseProperties{}

	// Initialize MarkLogic server (With or without a license)
	restartResp, errorResp, resp := c.RestService.Init(lic)

	if resp.StatusCode == 204 {
		log.Println("License modified no restart required.")
	} else if resp.StatusCode == 202 {
		log.Println("License inserted restart required.")
		log.Println(spew.Sdump(restartResp))
	} else {
		log.Println(errorResp.ErrorResponse.Message)
	}

}
