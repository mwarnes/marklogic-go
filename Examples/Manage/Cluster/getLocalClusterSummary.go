package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/mwarnes/marklogic-go"
	"log"
)

func main() {

	conn := marklogic.Connection{
		Host:               "mwca",
		Port:               8002,
		Username:           "admin",
		Password:           "admin",
		AuthenticationType: marklogic.BasicAuth,
	}

	c := marklogic.RestClient(conn)

	clusterSummary, errorResp, httpResp := c.RestService.GetLocalClusterSummary()

	if httpResp.StatusCode == 200 {
		log.Println(spew.Sdump(clusterSummary))
	} else {
		log.Println(httpResp.Status)
		log.Println(spew.Sdump(errorResp))
	}

}
