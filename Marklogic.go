package marklogic

import (
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
)

const (
	MlGoHttpVersion = "MarkLogic-Go-Client/2.0.0"
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
)

func ExecuteRequest(client Client, req *http.Request, successV, failureV interface{}) (*http.Response, error) {
	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	// Close body when finished
	defer response.Body.Close()

	// 204 No Content or 4xx; so no need to decode.
	if response.StatusCode == 204 || response.StatusCode >= 400 {
		return response, nil
	}

	// TODO At some point handle decoding JSON and XML
	//log.Println("Content Type:",response.Header.Get("Content-Type"))

	// Decode from json
	if successV != nil || failureV != nil {
		err = DecodeResponseJSON(response, successV, failureV)
	}

	return response, err
}

func DecodeResponseJSON(resp *http.Response, successV, failureV interface{}) error {
	if code := resp.StatusCode; 200 <= code && code <= 299 {
		if successV != nil {
			return DecodeResponseBodyJSON(resp, successV)
		}
	} else {
		if failureV != nil {
			return DecodeResponseBodyJSON(resp, failureV)
		}
	}
	return nil
}

func DecodeResponseBodyJSON(resp *http.Response, v interface{}) error {
	return json.NewDecoder(resp.Body).Decode(v)
}

func DecodeResponseBodyXML(resp *http.Response, v interface{}) error {
	return xml.NewDecoder(resp.Body).Decode(v)
}
