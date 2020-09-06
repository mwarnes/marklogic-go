package main

import (
	"github.com/mwarnes/marklogic-go"
	"io/ioutil"
	"log"
	"net/http"
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
	c := marklogic.MarkLogicRestClient(conn)

	req, _ := c.RestService.NewRequest("GET", marklogic.Timestamp, nil)

	httpResp, err := c.RestService.ExecuteRequest(req)

	if err == nil {
		defer httpResp.Body.Close()
		if httpResp.StatusCode == http.StatusOK {
			contents, err := ioutil.ReadAll(httpResp.Body)
			if err != nil {
				log.Fatalln(err)
			}
			log.Println(string(contents))
		}
	} else {
		log.Println(err)
	}

}
