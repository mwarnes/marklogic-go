package Structures

import (
	"time"
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
