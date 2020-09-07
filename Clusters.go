package marklogic

import (
	"github.com/mwarnes/marklogic-go/Structures"
	"log"
	"net/http"
)

const (
	Clusters               = "/manage/LATEST"
	LocalClusterProperties = "/manage/LATEST/properties"
)

func (s *RestService) GetLocalClusterSummary() (Structures.LocalClusterDefault, Structures.RestErrorResponse, http.Response) {
	req, _ := http.NewRequest("GET", s.base+Clusters, nil)

	s.Client = Decorate(s.Client,
		AddHeader("Accept", "application/json"),
	)
	clustersResponse := new(Structures.LocalClusterResponse)
	errorResponse := new(Structures.RestErrorResponse)
	response, err := ExecuteRequest(s.Client, req, clustersResponse, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return clustersResponse.LocalClusterDefault, *errorResponse, *response
}
