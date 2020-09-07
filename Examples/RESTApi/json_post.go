package main

import (
	"bytes"
	"crypto/tls"
	"github.com/mwarnes/marklogic-go"
	"log"
	"net/http"
	"os"
)

type JSONPOSTResponse struct {
	Success string `json:"success"`
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
		marklogic.AddHeader("Content-Type", "application/json"),
		marklogic.AddHeader("Accept", "application/json"),
		marklogic.Logging(logger),
	)

	body := `{"Id": 78912,"Customer": "Jason Sweet","Quantity": 1,"Price": 18.00}`

	req, _ := c.RestService.NewRequest("POST", "/echo/post/json", bytes.NewBufferString(body))

	httpResp, err := c.RestService.ExecuteRequest(req)

	response := new(JSONPOSTResponse)
	if err == nil {
		defer httpResp.Body.Close()
		if httpResp.StatusCode == http.StatusOK {
			marklogic.DecodeResponseBodyJSON(httpResp, response)
			log.Println(response.Success)
		}
	} else {
		log.Println(err)
	}

}
