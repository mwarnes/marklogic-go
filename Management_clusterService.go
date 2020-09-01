package marklogic

import (
	"log"
	"net/http"
)

const (
	ClustersV2             = "LATEST"
	LocalClusterProperties = "LATEST/properties"
)

// Clusters
type LocalClusterResponse struct {
	LocalClusterDefault LocalClusterDefault `json:"local-cluster-default"`
}

type LocalClusterDefault struct {
	ID               string       `json:"id"`
	Name             string       `json:"name"`
	Version          string       `json:"version"`
	EffectiveVersion int          `json:"effective-version"`
	Role             string       `json:"role"`
	Meta             Meta         `json:"meta"`
	Relations        Relations    `json:"relations"`
	RelatedViews     RelatedViews `json:"related-views"`
}

type ClusterProperties struct {
	ClusterID                  string          `json:"cluster-id,omitempty"`
	ClusterName                string          `json:"cluster-name,omitempty"`
	Role                       string          `json:"role,omitempty"`
	Version                    string          `json:"version,omitempty"`
	EffectiveVersion           int             `json:"effective-version,omitempty"`
	SecurityVersion            int             `json:"security-version,omitempty"`
	SslFipsEnabled             bool            `json:"ssl-fips-enabled,omitempty"`
	XdqpSslCertificate         string          `json:"xdqp-ssl-certificate,omitempty"`
	BootstrapHost              []BootstrapHost `json:"bootstrap-host,omitempty"`
	DataDirectory              string          `json:"data-directory,omitempty"`
	FilesystemDirectory        string          `json:"filesystem-directory,omitempty"`
	OpsdirectorLogLevel        string          `json:"opsdirector-log-level,omitempty"`
	OpsdirectorMetering        string          `json:"opsdirector-metering,omitempty"`
	OpsdirectorConfig          string          `json:"opsdirector-config,omitempty"`
	OpsdirectorSessionEndpoint interface{}     `json:"opsdirector-session-endpoint,omitempty"`
}

type BootstrapHost struct {
	BootstrapHostID      string `json:"bootstrap-host-id"`
	BootstrapHostName    string `json:"bootstrap-host-name"`
	BootstrapConnectPort int    `json:"bootstrap-connect-port"`
}

// Perform a local cluster operation
// Valid operations are:
// . "restart-local-cluster"
// . "commit-upgrade-local-cluster"
// . "security-database-upgrade-local-cluster"
type ClusterOperation struct {
	Operation string `json:"operation,omitempty"`
}

type ClusterService struct {
	client Client
	base   string
}

// NewService creates a new Admin service for processing MarkLogic Client REST API resquest.
// NewService takes a RestClient and builds a new sling HTTP Client configured with a Base URI and UserAgent header
// A new Service is returned
func NewClusterService(client Client, base string) *ClusterService {

	return &ClusterService{
		client: client,
		base:   base,
	}
}

// SetServerProperties sets the database properties
//func (s *ManageService) PerformClusterOperation(operation ClusterOperation) (LocalClusterSummary, RestErrorResponse, http.Response) {
//	req, _ := s.sling.New().Get(ClustersV2).
//		Add("Accept", "application/json").BodyJSON(operation).Request()
//	clustersResponse := new(LocalClusterSummary)
//	errorResponse := new(RestErrorResponse)
//	resp, err := s.sling.New().Do(req, clustersResponse, errorResponse)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	return *clustersResponse, *errorResponse, *resp
//}

func (s *ClusterService) GetLocalClusterSummary() (LocalClusterDefault, RestErrorResponse, http.Response) {
	req, _ := http.NewRequest("GET", s.base+ClustersV2, nil)

	s.client = Decorate(s.client,
		AddHeader("Accept", "application/json"),
	)
	clustersResponse := new(LocalClusterResponse)
	errorResponse := new(RestErrorResponse)
	response, err := ExecuteRequest(s.client, req, clustersResponse, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return clustersResponse.LocalClusterDefault, *errorResponse, *response
}

//func (s *ManageService) GetLocalClusterProperties() (ClusterProperties, RestErrorResponse, http.Response) {
//	req, _ := s.sling.New().Get(LocalClusterProperties).
//		Add("Accept", "application/json").Request()
//	clustersResponse := new(ClusterProperties)
//	errorResponse := new(RestErrorResponse)
//	resp, err := s.sling.New().Do(req, clustersResponse, errorResponse)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	return *clustersResponse, *errorResponse, *resp
//}
