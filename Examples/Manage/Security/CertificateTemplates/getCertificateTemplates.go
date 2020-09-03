package main

import (
	"log"

	"github.com/mwarnes/marklogic-go"

	"github.com/davecgh/go-spew/spew"
)

func main() {

	conn := marklogic.Connection{
		Host:               "mwca",
		Port:               8002,
		Username:           "admin",
		Password:           "admin",
		AuthenticationType: marklogic.DigestAuth,
	}

	c := marklogic.MarkLogicManageClient(conn)

	certificateTemplateList, errorResponse, resp := c.Security.GetCertificateTemplates()

	if resp.StatusCode == 200 {
		// log.Println(spew.Sdump(certificateTemplateList))
		log.Println(spew.Sdump(certificateTemplateList.ListItems))
	} else {
		log.Println(resp.Status)
		log.Println(spew.Sdump(errorResponse))
	}

}
