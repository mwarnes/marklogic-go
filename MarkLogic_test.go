package marklogic_test

import (
	"github.com/mwarnes/marklogic-go"
	"github.com/mwarnes/marklogic-go/Structures"
	"sync"
	"testing"
)

var seqMutex sync.Mutex

var TestingHost = "localhost"
var TestingAdminPort = 8001
var TestingAdminUserid = "admin"
var TestingAdminPassword = "admin"
var TestingAdminAuthType = marklogic.DigestAuth
var TestingManagementPort = 8002
var TestingManagementUserid = "admin"
var TestingManagementPassword = "admin"
var TestingManagementAuthType = marklogic.DigestAuth

func seq() func() {
	seqMutex.Lock()
	return func() {
		seqMutex.Unlock()
	}
}

func TestInitialAdminConnection(t *testing.T) {
	defer seq()()
	t.Log("Testing initial Admin connection")
	conn := marklogic.Connection{
		Host:               TestingHost,
		Port:               TestingAdminPort,
		Username:           TestingAdminUserid,
		Password:           TestingAdminPassword,
		AuthenticationType: TestingAdminAuthType,
	}
	// Create a new MarkLogic Admin REST API client
	c := marklogic.RestClient(conn)
	// Issue Timestamp request
	_, resp := c.RestService.Timestamp()
	if resp.StatusCode != 200 {
		t.Errorf("Response %s", resp.Status)
		t.Errorf("Response code, recevied: %d, expected: %d.", resp.StatusCode, 200)
	}
}

func TestInitialManagementConnection(t *testing.T) {
	defer seq()()
	t.Log("Testing initial Management connection")
	conn := marklogic.Connection{
		Host:               TestingHost,
		Port:               TestingManagementPort,
		Username:           TestingManagementUserid,
		Password:           TestingManagementPassword,
		AuthenticationType: TestingManagementAuthType,
	}
	c := marklogic.RestClient(conn)

	parms := Structures.ServerParameters{
		GroupId: "Default",
		View:    "default",
	}

	_, _, resp := c.RestService.GetAppServers(parms)

	if resp.StatusCode != 200 {
		t.Errorf("Response %s", resp.Status)
		t.Errorf("Response code, recevied: %d, expected: %d.", resp.StatusCode, 200)
	}
}

func TestCertificateTemplateCreation(t *testing.T) {
	defer seq()()
	t.Log("Testing Certificate Template Creation")
	conn := marklogic.Connection{
		Host:               TestingHost,
		Port:               TestingManagementPort,
		Username:           TestingManagementUserid,
		Password:           TestingManagementPassword,
		AuthenticationType: TestingManagementAuthType,
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
		TemplateName:        "GoLangTesting",
		TemplateDescription: "GoLang created template.",
		KeyType:             "rsa",
		KeyOptions:          keyOpts,
		Req:                 certReq,
	}

	_, resp := c.RestService.CreateCertificateTemplate(certTemplate)

	if resp.StatusCode != 201 {
		t.Errorf("Response %s", resp.Status)
		t.Errorf("Response code, recevied: %d, expected: %d.", resp.StatusCode, 201)
	}
}
