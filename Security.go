package marklogic

import (
	"bytes"
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/mwarnes/marklogic-go/Structures"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	CertificateTemplate    = "/manage/LATEST/certificate-templates"
	CertificateAuthorities = "/manage/LATEST/certificate-authorities"
	Certificates           = "/manage/LATEST/certificates"
	ExternalSecurity       = "/manage/LATEST/external-security"
)

func (s *RestService) GetCertificateTemplates() (Structures.CertificateTemplatesDefaultList, Structures.RestErrorResponse, http.Response) {
	req, _ := http.NewRequest("GET", s.base+CertificateTemplate, nil)

	s.client = Decorate(s.client,
		AddHeader("Accept", "application/json"),
	)
	certificateTemplateResponse := new(Structures.CertificateTemplatesResponse)
	errorResponse := new(Structures.RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, certificateTemplateResponse, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return certificateTemplateResponse.CertificateTemplatesDefaultList, *errorResponse, *resp
}

func (s *RestService) GetCertificateTemplate(template string) (Structures.CertificateTemplateDefault, Structures.RestErrorResponse, http.Response) {
	req, _ := http.NewRequest("GET", s.base+CertificateTemplate+"/"+template, nil)

	s.client = Decorate(s.client,
		AddHeader("Accept", "application/json"),
	)
	certificateTemplateSummary := new(Structures.CertificateTemplateResponse)
	errorResponse := new(Structures.RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, certificateTemplateSummary, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return certificateTemplateSummary.CertificateTemplateDefault, *errorResponse, *resp
}

func (s *RestService) CreateCertificateTemplate(templateProperties Structures.CertificateTemplateProperties) (Structures.RestErrorResponse, http.Response) {
	body, err := json.Marshal(templateProperties)
	if err != nil {
		log.Fatalln(err)
	}
	req, _ := http.NewRequest("POST", s.base+CertificateTemplate, bytes.NewBuffer(body))

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

func (s *RestService) PerformCertificateTemplateOperation(template string, operation Structures.CertificateTemplateOperation) ([]byte, http.Response) {
	body, err := json.Marshal(operation)
	if err != nil {
		log.Fatalln(err)
	}
	req, _ := http.NewRequest("POST", s.base+CertificateTemplate+"/"+template, bytes.NewBuffer(body))

	s.client = Decorate(s.client,
		AddHeader("Content-Type", "application/json"),
		AddHeader("Accept", "application/json"),
	)

	resp, err := s.client.Do(req)
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return contents, *resp
}

func (s *RestService) GetCertificateAuthorities() (Structures.CertificateAuthoritiesDefaultList, Structures.RestErrorResponse, http.Response) {
	req, _ := http.NewRequest("GET", s.base+CertificateAuthorities, nil)

	s.client = Decorate(s.client,
		AddHeader("Accept", "application/json"),
	)
	certificateAuthoritiesResponse := new(Structures.CertificateAuthoritiesResponse)
	errorResponse := new(Structures.RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, certificateAuthoritiesResponse, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return certificateAuthoritiesResponse.CertificateAuthoritiesDefaultList, *errorResponse, *resp
}

func (s *RestService) GetCertificateAuthority(id string) (Structures.CertificateAuthorityDefault, Structures.RestErrorResponse, http.Response) {
	req, _ := http.NewRequest("GET", s.base+CertificateAuthorities+"/"+id, nil)

	s.client = Decorate(s.client,
		AddHeader("Accept", "application/json"),
	)
	certificateAuthority := new(Structures.CertificateAuthorityResponse)
	errorResponse := new(Structures.RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, certificateAuthority, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return certificateAuthority.CertificateAuthorityDefault, *errorResponse, *resp
}

func (s *RestService) GetCertificateAuthorityProperties(id string) (Structures.CertificateAuthorityProperties, Structures.RestErrorResponse, http.Response) {
	req, _ := http.NewRequest("GET", s.base+CertificateAuthorities+"/"+id+"/properties", nil)

	s.client = Decorate(s.client,
		AddHeader("Accept", "application/json"),
	)
	certificateAuthorityProperties := new(Structures.CertificateAuthorityProperties)
	errorResponse := new(Structures.RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, certificateAuthorityProperties, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return *certificateAuthorityProperties, *errorResponse, *resp
}

func (s *RestService) GetCertificates() (Structures.CertificateDefaultList, Structures.RestErrorResponse, http.Response) {
	req, _ := http.NewRequest("GET", s.base+Certificates, nil)

	s.client = Decorate(s.client,
		AddHeader("Accept", "application/json"),
	)
	certificatesResponse := new(Structures.CertificatesResponse)
	errorResponse := new(Structures.RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, certificatesResponse, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return certificatesResponse.CertificateDefaultList, *errorResponse, *resp
}

func (s *RestService) AddCertificate(certificatePem string, trusted bool) (Structures.RestErrorResponse, http.Response) {
	var req *http.Request
	if trusted == true {
		req, _ = http.NewRequest("POST", s.base+Certificates+"?trusted=true", bytes.NewBufferString(certificatePem))
	} else {
		req, _ = http.NewRequest("POST", s.base+Certificates+"?trusted=false", bytes.NewBufferString(certificatePem))
	}

	s.client = Decorate(s.client,
		AddHeader("Content-Type", "text/plain"),
		AddHeader("Accept", "application/json"),
	)
	errorResponse := new(Structures.RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, nil, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return *errorResponse, *resp
}

func (s *RestService) GetCertificate(id string) (Structures.Certificate, Structures.RestErrorResponse, http.Response) {
	idUri := url.QueryEscape(id)
	req, _ := http.NewRequest("GET", s.base+Certificates+"/"+idUri, nil)

	s.client = Decorate(s.client,
		AddHeader("Accept", "application/json"),
	)
	certificateResponse := new(Structures.Certificate)
	errorResponse := new(Structures.RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, certificateResponse, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return *certificateResponse, *errorResponse, *resp
}

func (s *RestService) GetCertificateProperties(id string) (Structures.CertificateProperties, Structures.RestErrorResponse, http.Response) {
	idUri := url.QueryEscape(id)
	req, _ := http.NewRequest("GET", s.base+Certificates+"/"+idUri+"/properties", nil)

	s.client = Decorate(s.client,
		AddHeader("Accept", "application/json"),
	)
	certificateProperties := new(Structures.CertificateProperties)
	errorResponse := new(Structures.RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, certificateProperties, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return *certificateProperties, *errorResponse, *resp
}

func (s *RestService) DeleteCertificate(id string) (Structures.RestErrorResponse, http.Response) {
	req, _ := http.NewRequest("DELETE", s.base+Certificates+"/"+id, nil)

	s.client = Decorate(s.client,
		AddHeader("Accept", "application/json"),
	)
	errorResponse := new(Structures.RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, nil, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return *errorResponse, *resp
}

func (s *RestService) GetExternalSecurity() (Structures.ExternalSecurityDefaultList, Structures.RestErrorResponse, http.Response) {
	req, _ := http.NewRequest("GET", s.base+ExternalSecurity, nil)

	s.client = Decorate(s.client,
		AddHeader("Accept", "application/json"),
	)
	externalSecurityResponse := new(Structures.ExternalSecurityResponse)
	errorResponse := new(Structures.RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, externalSecurityResponse, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return externalSecurityResponse.ExternalSecurityDefaultList, *errorResponse, *resp
}

func (s *RestService) GetExternalSecurityConfiguration(extSecurityConfig string) (Structures.ExternalSecurityConfiguration, Structures.RestErrorResponse, http.Response) {
	req, _ := http.NewRequest("GET", s.base+ExternalSecurity+"/"+extSecurityConfig, nil)

	s.client = Decorate(s.client,
		AddHeader("Accept", "application/json"),
	)
	externalSecurityConfiguration := new(Structures.ExternalSecurityConfigurationResponse)
	errorResponse := new(Structures.RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, externalSecurityConfiguration, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return externalSecurityConfiguration.ExternalSecurityConfiguration, *errorResponse, *resp
}

func (s *RestService) GetExternalSecurityProperties(extSecurityConfig string, parms Structures.ExternalSecurityParameters) (interface{}, Structures.RestErrorResponse, http.Response) {
	v, err := query.Values(parms)
	if err != nil {
		log.Fatalln(err)
	}
	urlParms := v.Encode()

	req, _ := http.NewRequest("GET", s.base+ExternalSecurity+"/"+extSecurityConfig+"/properties?"+urlParms, nil)

	if strings.EqualFold(parms.Format, "xml") {
		s.client = Decorate(s.client,
			AddHeader("Accept", "application/xml"),
		)
	} else {
		s.client = Decorate(s.client,
			AddHeader("Accept", "application/json"),
		)
	}

	externalSecurityPropertiesResponse := new(Structures.ExternalSecurityProperties)

	errorResponse := new(Structures.RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, externalSecurityPropertiesResponse, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return externalSecurityPropertiesResponse, *errorResponse, *resp

}

func (s *RestService) AddExternalSecurityConfiguration(externalSecurityParameters Structures.ExternalSecurityConfigurationParameters) (Structures.RestErrorResponse, http.Response) {
	body, err := json.Marshal(externalSecurityParameters)
	if err != nil {
		log.Fatalln(err)
	}

	req, _ := http.NewRequest("POST", s.base+ExternalSecurity, bytes.NewBuffer(body))

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

func (s *RestService) UpdateExternalSecurityConfiguration(externalSecurityParameters Structures.ExternalSecurityConfigurationParameters) (Structures.RestartResponse, Structures.RestErrorResponse, http.Response) {

	body, err := json.Marshal(externalSecurityParameters)
	if err != nil {
		log.Fatalln(err)
	}

	req, _ := http.NewRequest("PUT", s.base+ExternalSecurity+"/"+externalSecurityParameters.ExternalSecurityName+"/properties", bytes.NewBuffer(body))

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

func (s *RestService) DeleteExternalSecurityConfiguration(extSecurityConfig string) (Structures.RestartResponse, Structures.RestErrorResponse, http.Response) {

	req, _ := http.NewRequest("DELETE", s.base+ExternalSecurity+"/"+extSecurityConfig, nil)

	s.client = Decorate(s.client,
		AddHeader("Accept", "application/json"),
	)

	restartResponse := new(Structures.RestartResponse)
	errorResponse := new(Structures.RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, restartResponse, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return *restartResponse, *errorResponse, *resp

}
