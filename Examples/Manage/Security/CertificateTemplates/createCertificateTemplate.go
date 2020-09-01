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

	certSubject := marklogic.Subject{
		CountryName:            "UK",
		OrganizationName:       "MarkLogic",
		OrganizationalUnitName: "Support",
		EmailAddress:           "martin.warnes@marklogic.com",
	}

	certReq := marklogic.Req{
		Version: "0",
		Subject: certSubject,
	}

	keyOpts := marklogic.KeyOptions{
		KeyLength: "2048",
	}

	certTemplate := marklogic.CertificateTemplateProperties{
		TemplateName:        "ssl1",
		TemplateDescription: "GoLang created template.",
		KeyType:             "rsa",
		KeyOptions:          keyOpts,
		Req:                 certReq,
	}

	errorResponse, resp := c.Security.CreateCertificateTemplate(certTemplate)

	if resp.StatusCode == 201 {
		log.Println("Template created.")
	} else {
		log.Println(resp.Status)
		log.Println(spew.Sdump(errorResponse))
	}

}
