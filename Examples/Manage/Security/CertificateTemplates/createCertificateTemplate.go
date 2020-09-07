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

	certSubject := Structures.Subject{
		CountryName:            "UK",
		OrganizationName:       "MarkLogic",
		OrganizationalUnitName: "Support",
		EmailAddress:           "martin.warnes@marklogic.com",
	}

	certReq := Structures.Req{
		Version: "0",
		Subject: certSubject,
	}

	keyOpts := Structures.KeyOptions{
		KeyLength: "2048",
	}

	certTemplate := Structures.CertificateTemplateProperties{
		TemplateName:        "ssl1",
		TemplateDescription: "GoLang created template.",
		KeyType:             "rsa",
		KeyOptions:          keyOpts,
		Req:                 certReq,
	}

	errorResponse, resp := c.RestService.CreateCertificateTemplate(certTemplate)

	if resp.StatusCode == 201 {
		log.Println("Template created.")
	} else {
		log.Println(resp.Status)
		log.Println(spew.Sdump(errorResponse))
	}
}
