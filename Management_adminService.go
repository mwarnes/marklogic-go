package marklogic

import (
	"bytes"
	"encoding/json"
	"github.com/google/go-querystring/query"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	Timestamp     = "v1/timestamp"
	Initialize    = "v1/init"
	InstanceAdmin = "v1/instance-admin"
	ServerConfig  = "v1/server-config"
	ClusterConfig = "v1/cluster-config"
)

type TimestampHead struct {
	StatusCode    int    `json:"status-code"`
	ServerDetails string `json:"server-details"`
	Connection    string `json:"connection"`
	KeepAlive     string `json:"keep-alive"`
}

type LicenseProperties struct {
	LicenseKey string `json:"license-key"`
	Licensee   string `json:"licensee"`
}

type SecurityProperties struct {
	AdminUsername string `url:"admin-username"`
	AdminPassword string `url:"admin-password"`
	Realm         string `url:"realm"`
}

type ClusterConfigProperties struct {
	Group        string `url:"group,omitempty"`
	ServerConfig string `url:"server-config,omitempty"`
	Zone         string `url:"server-config,omitempty"`
}

// The AdminService structure holds:
type AdminService struct {
	client Client
	base   string
}

// NewService creates a new Admin service for processing MarkLogic Admin REST API resquest.
// NewService takes a RestClient and builds a new sling HTTP Client configured with a Base URI and UserAgent header
// A new Service is returned
func NewAdminService(client Client, base string) *AdminService {

	return &AdminService{
		client: client,
		base:   base,
	}
}

// Timestamp returns the current MarkLogic server timestamp and an error if a problem was encountered.
func (s *AdminService) Timestamp() (string, http.Response) {
	req, _ := http.NewRequest("GET", s.base+Timestamp, nil)
	response, err := s.client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(contents), *response
}

//TimestampHead returns the current MarkLogic status and an error if a problem was encountered.
//Note: Despite it's name the server timestamp is not returned only the server status, if you need the actual server
//timestamp use Timestamp() instead.
func (s *AdminService) TimestampHead() (TimestampHead, http.Response) {
	req, _ := http.NewRequest("HEAD", s.base+Timestamp, nil)
	response, err := s.client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
	returnResp := TimestampHead{
		StatusCode:    response.StatusCode,
		ServerDetails: response.Header.Get("Server"),
		Connection:    response.Header.Get("Connection"),
		KeepAlive:     response.Header.Get("Keep-Alive"),
	}
	return returnResp, *response
}

//...
func (s *AdminService) Init(license LicenseProperties) (RestartResponse, RestErrorResponse, http.Response) {
	body, err := json.Marshal(license)
	if err != nil {
		log.Fatalln(err)
	}
	req, _ := http.NewRequest("POST", s.base+Initialize, bytes.NewBuffer(body))

	// TODO For the time being only JSON formatting accpeted.
	s.client = Decorate(s.client,
		AddHeader("Accept", "application/json"),
		AddHeader("Content-Type", "application/json"),
	)
	restartResponse := new(RestartResponse)
	errorResponse := new(RestErrorResponse)
	response, err := ExecuteRequest(s.client, req, restartResponse, errorResponse)
	return *restartResponse, *errorResponse, *response
}

func (s *AdminService) InstanceAdmin(properties SecurityProperties) (RestartResponse, RestErrorResponse, http.Response) {
	v, err := query.Values(properties)
	if err != nil {
		log.Fatalln(err)
	}
	body := v.Encode()

	req, _ := http.NewRequest("POST", s.base+InstanceAdmin, bytes.NewBufferString(body))

	// TODO For the time being only JSON formatting accpeted.
	s.client = Decorate(s.client,
		AddHeader("Accept", "application/json"),
		AddHeader("Content-Type", "application/x-www-form-urlencoded"),
	)
	restartResponse := new(RestartResponse)
	errorResponse := new(RestErrorResponse)
	response, err := ExecuteRequest(s.client, req, restartResponse, errorResponse)
	return *restartResponse, *errorResponse, *response
}

func (s *AdminService) GetServerConfig() (string, http.Response) {
	req, _ := http.NewRequest("GET", s.base+ServerConfig, nil)
	s.client = Decorate(s.client,
		AddHeader("Accept", "application/xml"),
	)
	response, err := s.client.Do(req)
	if err != nil {
		return "", *response
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(contents), *response
}

func (s *AdminService) SendClusterConfigForm(properties ClusterConfigProperties) (io.Reader, http.Response) {
	v, err := query.Values(properties)
	if err != nil {
		log.Fatalln(err)
	}
	body := v.Encode()

	req, _ := http.NewRequest("POST", s.base+ClusterConfig, bytes.NewBufferString(body))
	s.client = Decorate(s.client,
		AddHeader("Accept", "application/zip"),
		AddHeader("Content-Type", "application/x-www-form-urlencoded"),
	)
	response, err := s.client.Do(req)
	if err != nil {
		return nil, *response
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return bytes.NewBuffer(contents), *response
}

func (s *AdminService) SendClusterConfigZip(config io.Reader) (RestartResponse, RestErrorResponse, http.Response) {
	req, _ := http.NewRequest("POST", s.base+ClusterConfig, config)
	s.client = Decorate(s.client,
		AddHeader("Accept", "application/json"),
		AddHeader("Content-Type", "application/zip"),
	)
	restartResponse := new(RestartResponse)
	errorResponse := new(RestErrorResponse)
	response, _ := ExecuteRequest(s.client, req, restartResponse, errorResponse)
	return *restartResponse, *errorResponse, *response
}
