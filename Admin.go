package marklogic

import (
	"bytes"
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/mwarnes/marklogic-go/Structures"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	Timestamp     = "/admin/v1/timestamp"
	Initialize    = "/admin/v1/init"
	InstanceAdmin = "/admin/v1/instance-admin"
	ServerConfig  = "/admin/v1/server-config"
	ClusterConfig = "/admin/v1/cluster-config"
)

// Timestamp returns the current MarkLogic server timestamp and an error if a problem was encountered.
func (s *RestService) Timestamp() (string, http.Response) {
	req, _ := http.NewRequest("GET", s.base+Timestamp, nil)
	response, err := s.Client.Do(req)
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
func (s *RestService) TimestampHead() (Structures.TimestampHead, http.Response) {
	req, _ := http.NewRequest("HEAD", s.base+Timestamp, nil)
	response, err := s.Client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
	returnResp := Structures.TimestampHead{
		StatusCode:    response.StatusCode,
		ServerDetails: response.Header.Get("Server"),
		Connection:    response.Header.Get("Connection"),
		KeepAlive:     response.Header.Get("Keep-Alive"),
	}
	return returnResp, *response
}

//...
func (s *RestService) Init(license Structures.LicenseProperties) (Structures.RestartResponse, Structures.RestErrorResponse, http.Response) {
	body, err := json.Marshal(license)
	if err != nil {
		log.Fatalln(err)
	}
	req, _ := http.NewRequest("POST", s.base+Initialize, bytes.NewBuffer(body))

	// TODO For the time being only JSON formatting accpeted.
	s.Client = Decorate(s.Client,
		AddHeader("Accept", "application/json"),
		AddHeader("Content-Type", "application/json"),
	)
	restartResponse := new(Structures.RestartResponse)
	errorResponse := new(Structures.RestErrorResponse)
	response, err := ExecuteRequest(s.Client, req, restartResponse, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return *restartResponse, *errorResponse, *response
}

func (s *RestService) InstanceAdmin(properties Structures.SecurityProperties) (Structures.RestartResponse, Structures.RestErrorResponse, http.Response) {
	v, err := query.Values(properties)
	if err != nil {
		log.Fatalln(err)
	}
	body := v.Encode()

	req, _ := http.NewRequest("POST", s.base+InstanceAdmin, bytes.NewBufferString(body))

	// TODO For the time being only JSON formatting accpeted.
	s.Client = Decorate(s.Client,
		AddHeader("Accept", "application/json"),
		AddHeader("Content-Type", "application/x-www-form-urlencoded"),
	)
	restartResponse := new(Structures.RestartResponse)
	errorResponse := new(Structures.RestErrorResponse)
	response, err := ExecuteRequest(s.Client, req, restartResponse, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return *restartResponse, *errorResponse, *response
}

func (s *RestService) GetServerConfig() (string, http.Response) {
	req, _ := http.NewRequest("GET", s.base+ServerConfig, nil)
	s.Client = Decorate(s.Client,
		AddHeader("Accept", "application/xml"),
	)
	response, err := s.Client.Do(req)
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

func (s *RestService) SendClusterConfigForm(properties Structures.ClusterConfigProperties) (io.Reader, http.Response) {
	v, err := query.Values(properties)
	if err != nil {
		log.Fatalln(err)
	}
	body := v.Encode()

	req, _ := http.NewRequest("POST", s.base+ClusterConfig, bytes.NewBufferString(body))
	s.Client = Decorate(s.Client,
		AddHeader("Accept", "application/zip"),
		AddHeader("Content-Type", "application/x-www-form-urlencoded"),
	)
	response, err := s.Client.Do(req)
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

func (s *RestService) SendClusterConfigZip(config io.Reader) (Structures.RestartResponse, Structures.RestErrorResponse, http.Response) {
	req, _ := http.NewRequest("POST", s.base+ClusterConfig, config)
	s.Client = Decorate(s.Client,
		AddHeader("Accept", "application/json"),
		AddHeader("Content-Type", "application/zip"),
	)
	restartResponse := new(Structures.RestartResponse)
	errorResponse := new(Structures.RestErrorResponse)
	response, _ := ExecuteRequest(s.Client, req, restartResponse, errorResponse)

	return *restartResponse, *errorResponse, *response
}
