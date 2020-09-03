package marklogic

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	Documents = "LATEST/documents"
)

type DocumentProperties struct {
	URI                string            `url:"uri,omitempty"`
	Categories         []string          `url:"category,omitempty"`
	Collections        []string          `url:"collection,omitempty"`
	Permissions        map[string]string `url:"perm,omitempty"`
	Properties         map[string]string `url:"perm,omitempty"`
	Value              map[string]string `url:"value,omitempty"`
	Extract            string            `url:"extract,omitempty"`
	Repair             string            `url:"repair,omitempty"`
	Transform          string            `url:"transform,omitempty"`
	Trans              map[string]string `url:"trans,omitempty"`
	ForestName         string            `url:"forest-name,omitempty"`
	Txid               string            `url:"txid,omitempty"`
	Lang               string            `url:"lang,omitempty"`
	TemporalCollection string            `url:"temporal-collection,omitempty"`
	TemporalDocument   string            `url:"temporal-document,omitempty"`
	SystemTime         string            `url:"system-time,omitempty"`
}

type DocumentResponse struct {
	Collections    []string       `json:"collections"`
	Permissions    Permissions    `json:"permissions"`
	Properties     Properties     `json:"properties"`
	Quality        int            `json:"quality"`
	MetadataValues MetadataValues `json:"metadataValues"`
}

type Permissions []struct {
	RoleName     string   `json:"role-name"`
	Capabilities []string `json:"capabilities"`
}

type Properties struct {
	Playtype string `json:"playtype"`
}

type MetadataValues struct {
	Mk1 string `json:"mk1"`
	Mk2 string `json:"mk2"`
}

// The DocumentService structure holds:
type DocumentService struct {
	client Client
	base   string
}

type Service struct {
	client Client
	base   string
}

func NewService(client Client, base string) *Service {

	return &Service{
		base:   base,
		client: client,
	}
}

// NewService creates a new Admin service for processing MarkLogic Client REST API resquest.
// NewService takes a RestClient and builds a new sling HTTP Client configured with a Base URI and UserAgent header
// A new Service is returned
func NewDocumentService(client Client, base string) *DocumentService {

	return &DocumentService{
		client: client,
		base:   base,
	}
}

func (s *Service) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	url = strings.TrimLeft(url, "/")
	req, err := http.NewRequest(method, s.base+url, body)
	return req, err
}

func (s *Service) ExecuteRequest(req *http.Request) (*http.Response, error) {
	s.client = Decorate(s.client,
		AddHeader("Accept", "application/json"),
	)

	response, err := s.client.Do(req)
	if err != nil {
		return response, err
	}

	return response, err
}

func (s *DocumentService) Write(properties DocumentProperties, content io.Reader) (RestErrorResponse, http.Response, error) {
	v, err := query.Values(properties)
	if err != nil {
		log.Fatalln(err)
	}
	options := v.Encode()

	req, _ := http.NewRequest("PUT", s.base+Documents+"?"+options, content)

	s.client = Decorate(s.client,
		AddHeader("Accept", "application/json"),
	)
	errorResponse := new(RestErrorResponse)
	response, err := ExecuteRequest(s.client, req, nil, errorResponse)
	return *errorResponse, *response, err
}

func (s *DocumentService) Read(properties DocumentProperties) (io.Reader, http.Response) {
	v, err := query.Values(properties)
	if err != nil {
		log.Fatalln(err)
	}
	options := v.Encode()

	req, _ := http.NewRequest("GET", s.base+Documents+"?"+options, nil)

	response, err := s.client.Do(req)

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return bytes.NewBuffer(contents), *response
}

func (s *DocumentService) Delete(properties DocumentProperties) http.Response {
	v, err := query.Values(properties)
	if err != nil {
		log.Fatalln(err)
	}
	options := v.Encode()

	req, _ := http.NewRequest("DELETE", s.base+Documents+"?"+options, nil)

	response, err := s.client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}
	return *response
}
