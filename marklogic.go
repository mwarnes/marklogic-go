package marklogic

import (
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"net/http"
)

const (
	MlGoHttpVersion = "MarkLogic-Go-Client/1.0.0"
)

// Connection structure contains the information needed for a MarkLogic REST connection
type Connection struct {
	Host               string
	Port               int
	Username           string
	Password           string
	AuthenticationType int
	TLSConfig          *tls.Config
}

const (
	BasicAuth = iota
	DigestAuth
	KerberosAuth
	CertificateAuth
	None
)

// REST API
const (
	ClientRest = iota
	AdminRest
	ManageRest
)

const (
	Client_Path = ""
	Admin_Path  = "admin/"
	Manage_Path = "manage/"
)

func ExecuteRequest(client Client, req *http.Request, successV, failureV interface{}) (*http.Response, error) {
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Close body when finished
	defer response.Body.Close()

	// 204 No Content; so no need to decode.
	if response.StatusCode == 204 {
		return response, nil
	}

	// TODO At some point handle decoding JSON and XML
	//log.Println("Content Type:",response.Header.Get("Content-Type"))

	// Decode from json
	if successV != nil || failureV != nil {
		err = decodeResponseJSON(response, successV, failureV)
	}

	return response, err
}

func decodeResponseJSON(resp *http.Response, successV, failureV interface{}) error {
	if code := resp.StatusCode; 200 <= code && code <= 299 {
		if successV != nil {
			return decodeResponseBodyJSON(resp, successV)
		}
	} else {
		if failureV != nil {
			return decodeResponseBodyJSON(resp, failureV)
		}
	}
	return nil
}

func decodeResponseXML(resp *http.Response, successV, failureV interface{}) error {
	if code := resp.StatusCode; 200 <= code && code <= 299 {
		if successV != nil {
			return decodeResponseBodyXML(resp, successV)
		}
	} else {
		if failureV != nil {
			return decodeResponseBodyXML(resp, failureV)
		}
	}
	return nil
}

func decodeResponseBodyJSON(resp *http.Response, v interface{}) error {
	return json.NewDecoder(resp.Body).Decode(v)
}

func decodeResponseBodyXML(resp *http.Response, v interface{}) error {
	return xml.NewDecoder(resp.Body).Decode(v)
}
