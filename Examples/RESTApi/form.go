package main

import (
	"crypto/tls"
	"github.com/google/go-querystring/query"
	"github.com/mwarnes/marklogic-go"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type FormOptions struct {
	Key1 string `url:"key1,omitempty"`
	Key2 string `url:"key2,omitempty"`
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
		marklogic.Logging(logger),
	)

	formOptions := FormOptions{
		Key1: "value1",
		Key2: "value2",
	}

	v, err := query.Values(formOptions)
	if err != nil {
		log.Fatalln(err)
	}
	options := v.Encode()

	req, _ := c.RestService.NewRequest("POST", "/echo/post/form"+"?"+options, nil)

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
