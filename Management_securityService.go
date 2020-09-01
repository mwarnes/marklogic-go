package marklogic

import (
	"bytes"
	"encoding/json"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	CertificateTemplate    = "LATEST/certificate-templates"
	CertificateAuthorities = "LATEST/certificate-authorities"
	Certificates           = "LATEST/certificates"
	ExternalSecurity       = "LATEST/external-security"
)

type CertificateAuthoritiesResponse struct {
	CertificateAuthoritiesDefaultList CertificateAuthoritiesDefaultList `json:"certificate-authorities-default-list"`
}

type CertificateAuthoritiesDefaultList struct {
	Meta         Meta         `json:"meta"`
	ListItems    ListItems    `json:"list-items"`
	RelatedViews RelatedViews `json:"related-views"`
}

type CertificateAuthorityResponse struct {
	CertificateAuthorityDefault CertificateAuthorityDefault `json:"certificate-authority-default"`
}

type CertificateAuthorityDefault struct {
	ID           string       `json:"id"`
	Authority    string       `json:"authority"`
	Enabled      string       `json:"enabled"`
	HostName     string       `json:"host-name"`
	Cert         Cert         `json:"cert"`
	Meta         Meta         `json:"meta"`
	Relations    Relations    `json:"relations"`
	RelatedViews RelatedViews `json:"related-views"`
}

type Cert struct {
	Version       string `json:"version"`
	SerialNumber  string `json:"serialNumber"`
	SignatureType string `json:"signatureType"`
	Issuer        struct {
		CommonName             string `json:"commonName"`
		OrganizationName       string `json:"organizationName"`
		OrganizationalUnitName string `json:"organizationalUnitName"`
	} `json:"issuer"`
	Validity struct {
		NotBefore time.Time `json:"notBefore"`
		NotAfter  time.Time `json:"notAfter"`
	} `json:"validity"`
	Subject struct {
		CommonName             string `json:"commonName"`
		OrganizationName       string `json:"organizationName"`
		OrganizationalUnitName string `json:"organizationalUnitName"`
	} `json:"subject"`
	V3Ext struct {
		KeyUsage struct {
			Critical string `json:"critical"`
			Value    string `json:"value"`
		} `json:"keyUsage"`
		SubjectKeyIdentifier struct {
			Critical string `json:"critical"`
			Value    string `json:"value"`
		} `json:"subjectKeyIdentifier"`
		AuthorityKeyIdentifier struct {
			Critical string `json:"critical"`
			Value    string `json:"value"`
		} `json:"authorityKeyIdentifier"`
		BasicConstraints struct {
			Critical string `json:"critical"`
			Value    string `json:"value"`
		} `json:"basicConstraints"`
	} `json:"v3ext"`
}

type Certificate struct {
	CertificateDefault CertificateDefault `json:"certificate-default"`
}

type CertificateDefault struct {
	ID           string       `json:"id"`
	Authority    string       `json:"authority"`
	Enabled      string       `json:"enabled"`
	HostName     string       `json:"host-name"`
	Cert         Cert         `json:"cert"`
	Meta         Meta         `json:"meta"`
	Relations    Relations    `json:"relations"`
	RelatedViews RelatedViews `json:"related-views"`
}

type CertificateProperties struct {
	CertificateID string `json:"certificate-id"`
	Authority     string `json:"authority"`
	Enabled       string `json:"enabled"`
	HostName      string `json:"host-name"`
	Cert          Cert   `json:"cert"`
}

type CertificatesResponse struct {
	CertificateDefaultList CertificateDefaultList `json:"certificate-default-list"`
}

type CertificateDefaultList struct {
	Meta         Meta         `json:"meta"`
	ListItems    ListItems    `json:"list-items"`
	RelatedViews RelatedViews `json:"related-views"`
}

type CertificateAuthorityProperties struct {
	CertificateID string `json:"certificate-id"`
	Authority     string `json:"authority"`
	Enabled       string `json:"enabled"`
	HostName      string `json:"host-name"`
	Cert          Cert   `json:"cert"`
}

type TemplateCertificates struct {
	CertificateList CertificateList `json:"certificate-list"`
}

type CertificateList struct {
	Certificate []struct {
		TemplateID      string `json:"template-id"`
		CertificateID   string `json:"certificate-id"`
		TemplateVersion string `json:"template-version"`
		Temporary       string `json:"temporary"`
		Authority       string `json:"authority"`
		HostName        string `json:"host-name"`
		Pem             string `json:"pem"`
		Cert            struct {
			Version       int    `json:"version"`
			SerialNumber  string `json:"serialNumber"`
			SignatureType string `json:"signatureType"`
			Issuer        struct {
				CountryName            string `json:"countryName"`
				OrganizationName       string `json:"organizationName"`
				OrganizationalUnitName string `json:"organizationalUnitName"`
				EmailAddress           string `json:"emailAddress"`
				CommonName             string `json:"commonName"`
			} `json:"issuer"`
			Validity struct {
				NotBefore time.Time `json:"notBefore"`
				NotAfter  time.Time `json:"notAfter"`
			} `json:"validity"`
			Subject struct {
				CommonName string `json:"commonName"`
			} `json:"subject"`
			PublicKey string `json:"publicKey"`
			V3Ext     struct {
				BasicConstraints struct {
					Critical string `json:"critical"`
					Value    string `json:"value"`
				} `json:"basicConstraints"`
				NsCertType struct {
					Critical string `json:"critical"`
					Value    string `json:"value"`
				} `json:"nsCertType"`
			} `json:"v3ext"`
		} `json:"cert"`
	} `json:"certificate"`
}

type CertificateTemplateResponse struct {
	CertificateTemplateDefault CertificateTemplateDefault `json:"certificate-template-default"`
}

type CertificateTemplateDefault struct {
	ID                  string       `json:"id"`
	TemplateName        string       `json:"template-name"`
	TemplateDescription string       `json:"template-description"`
	TemplateVersion     string       `json:"template-version"`
	KeyType             string       `json:"key-type"`
	KeyOptions          string       `json:"key-options"`
	Req                 Req          `json:"req"`
	Meta                Meta         `json:"meta"`
	Relations           Relations    `json:"relations"`
	RelatedViews        RelatedViews `json:"related-views"`
}

type CertificateTemplatesResponse struct {
	CertificateTemplatesDefaultList CertificateTemplatesDefaultList `json:"certificate-templates-default-list"`
}

type CertificateTemplatesDefaultList struct {
	Meta         Meta         `json:"meta"`
	ListItems    ListItems    `json:"list-items"`
	RelatedViews RelatedViews `json:"related-views"`
}

type CertificateTemplateProperties struct {
	TemplateName        string     `json:"template-name,omitempty"`
	TemplateDescription string     `json:"template-description,omitempty"`
	KeyType             string     `json:"key-type,omitempty"`
	KeyOptions          KeyOptions `json:"key-options,omitempty"`
	Req                 Req        `json:"req,omitempty"`
}

type KeyOptions struct {
	KeyLength string `json:"key-length,omitempty"`
}

type Req struct {
	Version string  `json:"version,omitempty"`
	Subject Subject `json:"subject,omitempty"`
}

type Subject struct {
	CountryName            string `json:"countryName,omitempty"`
	StateOrProvinceName    string `json:"stateOrProvinceName,omitempty"`
	LocalityName           string `json:"localityName,omitempty"`
	OrganizationName       string `json:"organizationName,omitempty"`
	OrganizationalUnitName string `json:"organizationalUnitName,omitempty"`
	CommonName             string `json:"commonName,omitempty"`
	EmailAddress           string `json:"emailAddress,omitempty"`
}

type CertificateTemplateOperation struct {
	Operation   string `json:"operation,omitempty"`
	ValidFor    int    `json:"valid-for,omitempty"`
	CommonName  string `json:"common-name,omitempty"`
	DNSName     string `json:"dns-name,omitempty"`
	IPAddr      string `json:"ip-addr,omitempty"`
	Certs       string `json:"certs,omitempty"`
	Pkey        string `json:"pkey,omitempty"`
	IfNecessary bool   `json:"if-necessary,omitempty"`
}

type ExternalSecurityParameters struct {
	Format string `url:"format"` //html, json, or xml
}

type ExternalSecurityResponse struct {
	ExternalSecurityDefaultList ExternalSecurityDefaultList `json:"external-security-default-list"`
}

type ExternalSecurityDefaultList struct {
	Meta      Meta      `json:"meta"`
	ListItems ListItems `json:"list-items"`
	Relations struct {
		RelationGroup struct {
			Typeref       string `json:"typeref"`
			RelationCount struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"relation-count"`
			Relation struct {
				Uriref  string `json:"uriref"`
				Roleref string `json:"roleref"`
				Idref   string `json:"idref"`
				Nameref string `json:"nameref"`
			} `json:"relation"`
		} `json:"relation-group"`
	} `json:"relations"`
	RelatedViews RelatedViews `json:"related-views"`
}

type ExternalSecurityConfigurationResponse struct {
	ExternalSecurityConfiguration ExternalSecurityConfiguration `json:"external-security-default"`
}

type ExternalSecurityConfiguration struct {
	ID                   string       `json:"id"`
	ExternalSecurityName string       `json:"external-security-name"`
	Description          string       `json:"description"`
	Authentication       string       `json:"authentication"`
	CacheTimeout         int          `json:"cache-timeout"`
	Authorization        string       `json:"authorization"`
	LdapServerURI        string       `json:"ldap-server-uri"`
	LdapBase             string       `json:"ldap-base"`
	LdapAttribute        string       `json:"ldap-attribute"`
	LdapDefaultUser      string       `json:"ldap-default-user"`
	LdapPassword         string       `json:"ldap-password"`
	LdapBindMethod       string       `json:"ldap-bind-method"`
	Meta                 Meta         `json:"meta,omitempty"`
	Relations            Relations    `json:"relations"`
	RelatedViews         RelatedViews `json:"related-views"`
}

type ExternalSecurityConfigurationParameters struct {
	ExternalSecurityName string `json:"external-security-name,omitempty"`
	Description          string `json:"description,omitempty"`
	Authentication       string `json:"authentication,omitempty"`
	CacheTimeout         int    `json:"cache-timeout,omitempty"`
	Authorization        string `json:"authorization,omitempty"`
	LdapServerURI        string `json:"ldap-server-uri,omitempty"`
	LdapBase             string `json:"ldap-base,omitempty"`
	LdapAttribute        string `json:"ldap-attribute,omitempty"`
	LdapDefaultUser      string `json:"ldap-default-user,omitempty"`
	LdapPassword         string `json:"ldap-password,omitempty"`
	LdapBindMethod       string `json:"ldap-bind-method,omitempty"`
}

type ExternalSecurityProperties struct {
	ExternalSecurityName string `json:"external-security-name"`
	Description          string `json:"description"`
	Authentication       string `json:"authentication"`
	CacheTimeout         int    `json:"cache-timeout"`
	Authorization        string `json:"authorization"`
	LdapServer           struct {
		LdapServerURI         string `json:"ldap-server-uri"`
		LdapBase              string `json:"ldap-base"`
		LdapAttribute         string `json:"ldap-attribute"`
		LdapDefaultUser       string `json:"ldap-default-user"`
		LdapPassword          string `json:"ldap-password"`
		LdapBindMethod        string `json:"ldap-bind-method"`
		LdapMemberofAttribute string `json:"ldap-memberof-attribute"`
		LdapMemberAttribute   string `json:"ldap-member-attribute"`
	} `json:"ldap-server"`
	SamlServer struct {
		SamlEntityID               string `json:"saml-entity-id"`
		SamlPrivilegeAttributeName string `json:"saml-privilege-attribute-name"`
		HTTPOptions                struct {
			CredentialID string `json:"credential-id"`
			Method       string `json:"method"`
			Username     string `json:"username"`
			Password     string `json:"password"`
			ClientCert   string `json:"client-cert"`
			ClientKey    string `json:"client-key"`
			PassPhrase   string `json:"pass-phrase"`
		} `json:"http-options"`
	} `json:"saml-server"`
	SslClientCertificateAuthorities interface{} `json:"ssl-client-certificate-authorities"`
	SslRequireClientCertificate     bool        `json:"ssl-require-client-certificate"`
}

type SecurityService struct {
	client Client
	base   string
}

func NewSecurityService(client Client, base string) *SecurityService {

	return &SecurityService{
		client: client,
		base:   base,
	}
}

func (s *SecurityService) GetCertificateTemplates() (CertificateTemplatesDefaultList, RestErrorResponse, http.Response) {
	req, _ := http.NewRequest("GET", s.base+CertificateTemplate, nil)

	s.client = Decorate(s.client,
		AddHeader("Accept", "application/json"),
	)
	certificateTemplateResponse := new(CertificateTemplatesResponse)
	errorResponse := new(RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, certificateTemplateResponse, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return certificateTemplateResponse.CertificateTemplatesDefaultList, *errorResponse, *resp
}

func (s *SecurityService) GetCertificateTemplate(template string) (CertificateTemplateDefault, RestErrorResponse, http.Response) {
	req, _ := http.NewRequest("GET", s.base+CertificateTemplate+"/"+template, nil)

	s.client = Decorate(s.client,
		AddHeader("Accept", "application/json"),
	)
	certificateTemplateSummary := new(CertificateTemplateResponse)
	errorResponse := new(RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, certificateTemplateSummary, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return certificateTemplateSummary.CertificateTemplateDefault, *errorResponse, *resp
}

func (s *SecurityService) CreateCertificateTemplate(templateProperties CertificateTemplateProperties) (RestErrorResponse, http.Response) {
	body, err := json.Marshal(templateProperties)
	if err != nil {
		log.Fatalln(err)
	}
	req, _ := http.NewRequest("POST", s.base+CertificateTemplate, bytes.NewBuffer(body))

	s.client = Decorate(s.client,
		AddHeader("Content-Type", "application/json"),
		AddHeader("Accept", "application/json"),
	)
	errorResponse := new(RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, nil, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return *errorResponse, *resp
}

func (s *SecurityService) PerformCertificateTemplateOperation(template string, operation CertificateTemplateOperation) ([]byte, http.Response) {
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

func (s *SecurityService) GetCertificateAuthorities() (CertificateAuthoritiesDefaultList, RestErrorResponse, http.Response) {
	req, _ := http.NewRequest("GET", s.base+CertificateAuthorities, nil)

	s.client = Decorate(s.client,
		AddHeader("Accept", "application/json"),
	)
	certificateAuthoritiesResponse := new(CertificateAuthoritiesResponse)
	errorResponse := new(RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, certificateAuthoritiesResponse, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return certificateAuthoritiesResponse.CertificateAuthoritiesDefaultList, *errorResponse, *resp
}

func (s *SecurityService) GetCertificateAuthority(id string) (CertificateAuthorityDefault, RestErrorResponse, http.Response) {
	req, _ := http.NewRequest("GET", s.base+CertificateAuthorities+"/"+id, nil)

	s.client = Decorate(s.client,
		AddHeader("Accept", "application/json"),
	)
	certificateAuthority := new(CertificateAuthorityResponse)
	errorResponse := new(RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, certificateAuthority, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return certificateAuthority.CertificateAuthorityDefault, *errorResponse, *resp
}

func (s *SecurityService) GetCertificateAuthorityProperties(id string) (CertificateAuthorityProperties, RestErrorResponse, http.Response) {
	req, _ := http.NewRequest("GET", s.base+CertificateAuthorities+"/"+id+"/properties", nil)

	s.client = Decorate(s.client,
		AddHeader("Accept", "application/json"),
	)
	certificateAuthorityProperties := new(CertificateAuthorityProperties)
	errorResponse := new(RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, certificateAuthorityProperties, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return *certificateAuthorityProperties, *errorResponse, *resp
}

func (s *SecurityService) GetCertificates() (CertificateDefaultList, RestErrorResponse, http.Response) {
	req, _ := http.NewRequest("GET", s.base+Certificates, nil)

	s.client = Decorate(s.client,
		AddHeader("Accept", "application/json"),
	)
	certificatesResponse := new(CertificatesResponse)
	errorResponse := new(RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, certificatesResponse, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return certificatesResponse.CertificateDefaultList, *errorResponse, *resp
}

func (s *SecurityService) AddCertificate(certificatePem string, trusted bool) (RestErrorResponse, http.Response) {
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
	errorResponse := new(RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, nil, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return *errorResponse, *resp
}

func (s *SecurityService) GetCertificate(id string) (Certificate, RestErrorResponse, http.Response) {
	idUri := url.QueryEscape(id)
	req, _ := http.NewRequest("GET", s.base+Certificates+"/"+idUri, nil)

	s.client = Decorate(s.client,
		AddHeader("Accept", "application/json"),
	)
	certificateResponse := new(Certificate)
	errorResponse := new(RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, certificateResponse, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return *certificateResponse, *errorResponse, *resp
}

func (s *SecurityService) GetCertificateProperties(id string) (CertificateProperties, RestErrorResponse, http.Response) {
	idUri := url.QueryEscape(id)
	req, _ := http.NewRequest("GET", s.base+Certificates+"/"+idUri+"/properties", nil)

	s.client = Decorate(s.client,
		AddHeader("Accept", "application/json"),
	)
	certificateProperties := new(CertificateProperties)
	errorResponse := new(RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, certificateProperties, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return *certificateProperties, *errorResponse, *resp
}

func (s *SecurityService) DeleteCertificate(id string) (RestErrorResponse, http.Response) {
	req, _ := http.NewRequest("DELETE", s.base+Certificates+"/"+id, nil)

	s.client = Decorate(s.client,
		AddHeader("Accept", "application/json"),
	)
	errorResponse := new(RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, nil, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return *errorResponse, *resp
}

func (s *SecurityService) GetExternalSecurity() (ExternalSecurityDefaultList, RestErrorResponse, http.Response) {
	req, _ := http.NewRequest("GET", s.base+ExternalSecurity, nil)

	s.client = Decorate(s.client,
		AddHeader("Accept", "application/json"),
	)
	externalSecurityResponse := new(ExternalSecurityResponse)
	errorResponse := new(RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, externalSecurityResponse, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return externalSecurityResponse.ExternalSecurityDefaultList, *errorResponse, *resp
}

func (s *SecurityService) GetExternalSecurityConfiguration(extSecurityConfig string) (ExternalSecurityConfiguration, RestErrorResponse, http.Response) {
	req, _ := http.NewRequest("GET", s.base+ExternalSecurity+"/"+extSecurityConfig, nil)

	s.client = Decorate(s.client,
		AddHeader("Accept", "application/json"),
	)
	externalSecurityConfiguration := new(ExternalSecurityConfigurationResponse)
	errorResponse := new(RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, externalSecurityConfiguration, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return externalSecurityConfiguration.ExternalSecurityConfiguration, *errorResponse, *resp
}

func (s *SecurityService) GetExternalSecurityProperties(extSecurityConfig string, parms ExternalSecurityParameters) (interface{}, RestErrorResponse, http.Response) {
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

	externalSecurityPropertiesResponse := new(ExternalSecurityProperties)

	errorResponse := new(RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, externalSecurityPropertiesResponse, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return externalSecurityPropertiesResponse, *errorResponse, *resp

}

func (s *SecurityService) AddExternalSecurityConfiguration(externalSecurityParameters ExternalSecurityConfigurationParameters) (RestErrorResponse, http.Response) {
	body, err := json.Marshal(externalSecurityParameters)
	if err != nil {
		log.Fatalln(err)
	}

	req, _ := http.NewRequest("POST", s.base+ExternalSecurity, bytes.NewBuffer(body))

	s.client = Decorate(s.client,
		AddHeader("Content-Type", "application/json"),
		AddHeader("Accept", "application/json"),
	)
	errorResponse := new(RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, nil, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return *errorResponse, *resp
}

func (s *SecurityService) UpdateExternalSecurityConfiguration(externalSecurityParameters ExternalSecurityConfigurationParameters) (RestartResponse, RestErrorResponse, http.Response) {

	body, err := json.Marshal(externalSecurityParameters)
	if err != nil {
		log.Fatalln(err)
	}

	req, _ := http.NewRequest("PUT", s.base+ExternalSecurity+"/"+externalSecurityParameters.ExternalSecurityName+"/properties", bytes.NewBuffer(body))

	s.client = Decorate(s.client,
		AddHeader("Content-Type", "application/json"),
		AddHeader("Accept", "application/json"),
	)
	errorResponse := new(RestErrorResponse)
	restartResponse := new(RestartResponse)
	resp, err := ExecuteRequest(s.client, req, restartResponse, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return *restartResponse, *errorResponse, *resp
}

func (s *SecurityService) DeleteExternalSecurityConfiguration(extSecurityConfig string) (RestartResponse, RestErrorResponse, http.Response) {

	req, _ := http.NewRequest("DELETE", s.base+ExternalSecurity+"/"+extSecurityConfig, nil)

	s.client = Decorate(s.client,
		AddHeader("Accept", "application/json"),
	)

	restartResponse := new(RestartResponse)
	errorResponse := new(RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, restartResponse, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return *restartResponse, *errorResponse, *resp

}
