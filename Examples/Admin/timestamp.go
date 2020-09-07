package main

import (
	"github.com/mwarnes/marklogic-go"
	"log"
)

// Issue a Timestamp GET Request
// https://docs.marklogic.com/REST/GET/admin/v1/timestamp
func main() {

	//tlsConfig := &tls.Config{
	//	InsecureSkipVerify:true,
	//}

	// MarkLogic REST API Connection parameters
	conn := marklogic.Connection{
		Host:               "mwca",
		Port:               8001,
		Username:           "admin",
		Password:           "admin",
		AuthenticationType: marklogic.DigestAuth,
		//TLSConfig: tlsConfig,
	}

	// Create a new MarkLogic Admin REST API client
	c := marklogic.RestClient(conn)

	// Issue Timestamp request
	timestamp, _ := c.RestService.Timestamp()

	log.Println("Current timestamp:", timestamp)
}
