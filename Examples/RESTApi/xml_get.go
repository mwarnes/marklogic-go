package main

import (
	"crypto/tls"
	"encoding/xml"
	"github.com/mwarnes/marklogic-go"
	"log"
	"net/http"
	"os"
)

type XMLResponse struct {
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

	req, _ := c.RestService.NewRequest("GET", "/echo/get/xml", nil)

	httpResp, err := c.RestService.ExecuteRequest(req)

	response := new(XMLResponse)
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
