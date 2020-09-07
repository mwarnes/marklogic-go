package main

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"github.com/mwarnes/marklogic-go"
	"log"
	"net/http"
	"os"
)

type XMLPOSTResponse struct {
	XMLName         xml.Name `xml:"Response"`
	Text            string   `xml:",chardata"`
	ResponseCode    string   `xml:"ResponseCode"`
	ResponseMessage string   `xml:"ResponseMessage"`
}

func main() {

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	logger := log.New(os.Stderr, "", log.LstdFlags)

	conn := marklogic.Connection{
		Host:               "reqbin.com",
		Port:               443,
		AuthenticationType: marklogic.None,
		TLSConfig:          tlsConfig,
	}

	c := marklogic.RestClient(conn)

	c.RestService.Client = marklogic.Decorate(c.RestService.Client,
		marklogic.AddHeader("Content-Type", "application/xml"),
		marklogic.AddHeader("Accept", "application/xml"),
		marklogic.Logging(logger),
	)

	body := `<?xml version="1.0" encoding="utf-8"?>
				<Request>
    				<Login>login</Login>
    				<Password>password</Password>
				</Request>`

	req, _ := c.RestService.NewRequest("POST", "/echo/post/xml", bytes.NewBufferString(body))

	httpResp, err := c.RestService.ExecuteRequest(req)

	response := new(XMLPOSTResponse)
	if err == nil {
		defer httpResp.Body.Close()
		if httpResp.StatusCode == http.StatusOK {
			marklogic.DecodeResponseBodyXML(httpResp, response)
			log.Println(response.ResponseCode)
		}
	} else {
		log.Println(err)
	}

}
