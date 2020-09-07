package marklogic

import (
	"bytes"
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/mwarnes/marklogic-go/Structures"
	"log"
	"net/http"
	"strings"
)

const (
	AppServers = "/manage/LATEST/servers"
)

//https://docs.marklogic.com/REST/GET/manage/v2/servers
func (s *RestService) GetAppServers(parms Structures.ServerParameters) (interface{}, Structures.RestErrorResponse, http.Response) {
	v, err := query.Values(parms)
	if err != nil {
		log.Fatalln(err)
	}
	urlParms := v.Encode()

	req, _ := http.NewRequest("GET", s.base+AppServers+"?"+urlParms, nil)

	if strings.EqualFold(parms.Format, "xml") {
		s.client = Decorate(s.client,
			AddHeader("Accept", "application/xml"),
		)
	} else {
		s.client = Decorate(s.client,
			AddHeader("Accept", "application/json"),
		)
	}

	var appServerResponse interface{}
	if strings.EqualFold(parms.View, "Metrics") {
		appServerResponse = new(Structures.AppServersMetricResponse)
	} else if strings.EqualFold(parms.View, "Status") {
		appServerResponse = new(Structures.AppServersStatusResponse)
	} else if strings.EqualFold(parms.View, "Package") {
		appServerResponse = new(Structures.AppServersPackageResponse)
	} else { // Default
		appServerResponse = new(Structures.AppServersDefaultResponse)
	}
	errorResponse := new(Structures.RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, appServerResponse, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return appServerResponse, *errorResponse, *resp
}

func (s *RestService) GetAppServer(appServer string, parms Structures.ServerParameters) (interface{}, Structures.RestErrorResponse, http.Response) {
	v, err := query.Values(parms)
	if err != nil {
		log.Fatalln(err)
	}
	urlParms := v.Encode()

	req, _ := http.NewRequest("GET", s.base+AppServers+"/"+appServer+"?"+urlParms, nil)

	if strings.EqualFold(parms.Format, "xml") {
		s.client = Decorate(s.client,
			AddHeader("Accept", "application/xml"),
		)
	} else {
		s.client = Decorate(s.client,
			AddHeader("Accept", "application/json"),
		)
	}

	var appServerResponse interface{}
	if strings.EqualFold(parms.View, "Status") {
		appServerResponse = new(Structures.AppServerStatusResponse)
	} else if strings.EqualFold(parms.View, "Package") {
		appServerResponse = new(Structures.AppServerPackageResponse)
	} else { // Default
		appServerResponse = new(Structures.AppServerDefaultResponse)
	}
	errorResponse := new(Structures.RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, appServerResponse, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return appServerResponse, *errorResponse, *resp

}

func (s *RestService) GetAppServerProperties(appServer string, parms Structures.ServerParameters) (interface{}, Structures.RestErrorResponse, http.Response) {
	v, err := query.Values(parms)
	if err != nil {
		log.Fatalln(err)
	}
	urlParms := v.Encode()

	req, _ := http.NewRequest("GET", s.base+AppServers+"/"+appServer+"/properties?"+urlParms, nil)

	if strings.EqualFold(parms.Format, "xml") {
		s.client = Decorate(s.client,
			AddHeader("Accept", "application/xml"),
		)
	} else {
		s.client = Decorate(s.client,
			AddHeader("Accept", "application/json"),
		)
	}

	appServerResponse := new(Structures.AppServerProperties)

	errorResponse := new(Structures.RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, appServerResponse, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return appServerResponse, *errorResponse, *resp

}

func (s *RestService) AddAppServer(appServer Structures.AppServerProperties) (Structures.RestErrorResponse, http.Response) {
	body, err := json.Marshal(appServer)
	if err != nil {
		log.Fatalln(err)
	}
	req, _ := http.NewRequest("POST", s.base+AppServers, bytes.NewBuffer(body))

	s.client = Decorate(s.client,
		AddHeader("Content-Type", "application/json"),
		AddHeader("Accept", "application/json"),
	)
	errorResponse := new(Structures.RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, nil, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return *errorResponse, *resp
}

func (s *RestService) UpdateAppServer(appServer Structures.AppServerProperties) (Structures.RestartResponse, Structures.RestErrorResponse, http.Response) {

	body, err := json.Marshal(appServer)
	if err != nil {
		log.Fatalln(err)
	}

	req, _ := http.NewRequest("PUT", s.base+AppServers+"/"+appServer.ServerName+"/properties?group-id="+appServer.GroupName, bytes.NewBuffer(body))

	s.client = Decorate(s.client,
		AddHeader("Content-Type", "application/json"),
		AddHeader("Accept", "application/json"),
	)
	errorResponse := new(Structures.RestErrorResponse)
	restartResponse := new(Structures.RestartResponse)
	resp, err := ExecuteRequest(s.client, req, restartResponse, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return *restartResponse, *errorResponse, *resp
}

func (s *RestService) DeleteAppServer(appServer string, parms Structures.ServerParameters) (Structures.RestartResponse, Structures.RestErrorResponse, http.Response) {
	v, err := query.Values(parms)
	if err != nil {
		log.Fatalln(err)
	}
	urlParms := v.Encode()

	req, _ := http.NewRequest("DELETE", s.base+AppServers+"/"+appServer+"?"+urlParms, nil)

	if strings.EqualFold(parms.Format, "xml") {
		s.client = Decorate(s.client,
			AddHeader("Accept", "application/xml"),
		)
	} else {
		s.client = Decorate(s.client,
			AddHeader("Accept", "application/json"),
		)
	}

	restartResponse := new(Structures.RestartResponse)
	errorResponse := new(Structures.RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, restartResponse, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return *restartResponse, *errorResponse, *resp

}
