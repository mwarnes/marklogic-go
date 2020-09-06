package main

import (
	"github.com/mwarnes/marklogic-go"
	"log"
)

func main() {

	//tlsConfig := &tls.Config{
	//	InsecureSkipVerify: true,
	//}

	// MarkLogic REST API Connection parameters
	conn := marklogic.Connection{
		Host:               "mwca",
		Port:               8000,
		Username:           "admin",
		Password:           "admin",
		AuthenticationType: marklogic.DigestAuth,
		//TLSConfig:          tlsConfig,
	}

	// Create a new MarkLogic Admin REST API client
	c := marklogic.MarkLogicRestClient(conn)

	req, _ := c.RestService.NewRequest("GET", "/v1/config/indexes", nil)

	httpResp, err := c.RestService.ExecuteRequest(req)
	defer httpResp.Body.Close()
	log.Println(httpResp.Status, err)

}
