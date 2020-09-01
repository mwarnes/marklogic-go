package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/mwarnes/marklogic-go"
	"log"
)

func main() {

	// MarkLogic REST API Connection parameters
	conn := marklogic.Connection{
		Host:               "ml-node-1",
		Port:               8002,
		Username:           "admin",
		Password:           "admin",
		AuthenticationType: marklogic.DigestAuth,
	}

	c := marklogic.MarkLogicManageClient(conn)

	var pem = `-----BEGIN CERTIFICATE-----
MIIDejCCAmKgAwIBAgIBbzANBgkqhkiG9w0BAQUFADBOMSYwJAYDVQQDDB1Qcmlt
YXJ5IENlcnRpZmljYXRlIEF1dGhvcml0eTESMBAGA1UECgwJTWFya0xvZ2ljMRAw
DgYDVQQLDAdTdXBwb3J0MB4XDTE3MDYwODE0MjEyMloXDTM2MDYwODExMDAwMFow
TjEmMCQGA1UEAwwdUHJpbWFyeSBDZXJ0aWZpY2F0ZSBBdXRob3JpdHkxEjAQBgNV
BAoMCU1hcmtMb2dpYzEQMA4GA1UECwwHU3VwcG9ydDCCASIwDQYJKoZIhvcNAQEB
BQADggEPADCCAQoCggEBAKumF/sDqeybEUrN1YtWHDwCMut4hkVha6rqpaz9X+KR
x2alZDoLgGDvC1pdDjD2eaTPExj+hmsfuRYYcbZNPNR+XZ3euYt9fyIbdu4hmQs2
xZDJUZS+qgnOAD+lYb9MX5YSX09D+BfjQ8uek9Y/FdBFFsDwSkCd1KSGb8ZN1yRN
ibXyBu4dZqm21nBgy016rWatCNHkEcsVDYqf5bzg7JOAv79rXFE8MgTBVcZLJEIE
16K52QTPM4DbffB69ioEc5IwIZnhvVR7LUhIpHh6YBkChrCmYni4j2VKX7R6jn8X
e82yVGWR3lEzm4wyKuyI79gou7NngkFCPZPDpUiL8rkCAwEAAaNjMGEwDgYDVR0P
AQH/BAQDAgEGMB0GA1UdDgQWBBTYAjlkvesFGEq2RPgf+y1Yw5uM2zAfBgNVHSME
GDAWgBTYAjlkvesFGEq2RPgf+y1Yw5uM2zAPBgNVHRMBAf8EBTADAQH/MA0GCSqG
SIb3DQEBBQUAA4IBAQAwGyg5ngjrpdU6uULw4VtavOTQNVH9Kv6dmkicQSo3E6rE
juszMT3YdqDGTc4Tzj/YT+WJqnYQ0A+t2RNijKQzp3z0RtETqtPMlSlzwjoaF8F3
hwz+pqU2vEKdNXU3CkWHcycTJTCJXpOqt8emBPapbxjhG6j5hftfo8QHBLh1pIeQ
Umdn3uZe+uDqG9xh2ryo4c/w26uNtfa/neZBtV5a9Wp0zWXzPJMhCry+dl2/6u6Z
4a3Y+X1Mk1i2pfWeP1Ih2CNQl7Lk2x8NF/HAW0iVXVRd1UXeGumV4wjCjmaLfRJT
kQLil0KjORbt+X0vP3Jzy1XYgd2ru4DMFX7VkVj1
-----END CERTIFICATE-----`

	restError, resp := c.Security.AddCertificate(pem, true)

	if resp.StatusCode == 201 {
		log.Println("Certificate added.")
	} else {
		log.Println(spew.Sdump(restError))
	}

}
