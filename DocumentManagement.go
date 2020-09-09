package marklogic

import (
	"bytes"
	"github.com/google/go-querystring/query"
	"github.com/mwarnes/marklogic-go/Structures"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	Documents = "/LATEST/documents"
)

func (s *RestService) Write(properties Structures.DocumentProperties, content io.Reader) (Structures.RestErrorResponse, http.Response, error) {
	v, err := query.Values(properties)
	if err != nil {
		log.Fatalln(err)
	}
	options := v.Encode()

	req, _ := http.NewRequest("PUT", s.base+Documents+"?"+options, content)

	s.Client = Decorate(s.Client,
		AddHeader("Accept", "application/json"),
	)
	errorResponse := new(Structures.RestErrorResponse)
	response, err := ExecuteRequest(s.Client, req, nil, errorResponse)

	if err != nil {
		log.Fatalln(err)
	}

	return *errorResponse, *response, err
}

func (s *RestService) Read(properties Structures.DocumentProperties) (io.Reader, http.Response) {
	v, err := query.Values(properties)
	if err != nil {
		log.Fatalln(err)
	}
	options := v.Encode()

	req, _ := http.NewRequest("GET", s.base+Documents+"?"+options, nil)

	response, err := s.Client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return bytes.NewBuffer(contents), *response
}

func (s *RestService) Delete(properties Structures.DocumentProperties) http.Response {
	v, err := query.Values(properties)
	if err != nil {
		log.Fatalln(err)
	}
	options := v.Encode()

	req, _ := http.NewRequest("DELETE", s.base+Documents+"?"+options, nil)

	response, err := s.Client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}
	return *response
}
