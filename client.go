package marklogic

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dpotapov/go-spnego"
	"github.com/mwarnes/digest"
)

// Client Connection
// This is a very simple lean interface which we can "decorate" with additional
// functionality only if and when it is needed
type Client interface {
	Do(*http.Request) (*http.Response, error)
}

// ClientFunc is a function type that implements the Client interface.
type ClientFunc func(*http.Request) (*http.Response, error)

func (f ClientFunc) Do(r *http.Request) (*http.Response, error) {
	return f(r)
}

// A Decorator wraps a Client with extra behaviour
type Decorator func(c Client) Client

// AddBasicAuthentication returns a Decorator that adds Basic Authentication to a MarkLogic Client
func AddBasicAuthentication(conn Connection) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			r.SetBasicAuth(conn.Username, conn.Password)
			return c.Do(r)
		})

	}
}

func AddHeader(name, value string) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			r.Header.Add(name, value)
			return c.Do(r)
		})

	}
}

// Logging returns a Decorator that logs a MarkLogic Client requests
func Logging(l *log.Logger) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			l.Printf("%s: %s %s", r.Header.Get("User-Agent"), r.Method, r.URL)
			return c.Do(r)
		})

	}
}

// Decorate decorates a Client c with all the given Decorators, in order.
func Decorate(c Client, ds ...Decorator) Client {
	decorated := c
	for _, decorate := range ds {
		decorated = decorate(decorated)
	}
	return decorated
}

type RestClient struct {
	Document *DocumentService
}

type AdminClient struct {
	Admin *AdminService
}

type ManageClient struct {
	AppServer *AppServerService
	Cluster   *ClusterService
	Security  *SecurityService
}

func MarkLogicAdminClient(connection Connection) AdminClient {

	client, base := getBasicClientAndBase(connection, AdminRest)

	cli := Decorate(client,
		AddHeader("User-Agent", MlGoHttpVersion),
		//Logging(log.New(os.Stdout, "client: ", log.LstdFlags)),
	)

	if connection.AuthenticationType == BasicAuth {
		cli = Decorate(client,
			AddBasicAuthentication(connection),
		)
	}
	adminService := NewAdminService(cli, base)

	return AdminClient{
		Admin: adminService,
	}
}

func MarkLogicManageClient(connection Connection) ManageClient {

	client, base := getBasicClientAndBase(connection, ManageRest)

	cli := Decorate(client,
		AddHeader("User-Agent", MlGoHttpVersion),
		//Logging(log.New(os.Stdout, "client: ", log.LstdFlags)),
	)

	if connection.AuthenticationType == BasicAuth {
		cli = Decorate(client,
			AddBasicAuthentication(connection),
		)
	}
	clusterService := NewClusterService(cli, base)
	securityService := NewSecurityService(cli, base)
	appServerService := NewAppServerService(cli, base)

	return ManageClient{
		AppServer: appServerService,
		Cluster:   clusterService,
		Security:  securityService,
	}
}

func MarkLogicRestClient(connection Connection) RestClient {

	client, base := getBasicClientAndBase(connection, ClientRest)

	cli := Decorate(client,
		AddHeader("User-Agent", MlGoHttpVersion),
		//Logging(log.New(os.Stdout, "client: ", log.LstdFlags)),
	)

	if connection.AuthenticationType == BasicAuth {
		cli = Decorate(client,
			AddBasicAuthentication(connection),
		)
	}
	documentService := NewDocumentService(cli, base)

	return RestClient{
		Document: documentService,
	}
}

func getBasicClientAndBase(connection Connection, clientType int) (Client, string) {

	var httpClient *http.Client

	switch connection.AuthenticationType {
	case DigestAuth:
		transport := digest.NewTransport(connection.Username, connection.Password, connection.TLSConfig)
		httpClient = &http.Client{Transport: transport}
	case KerberosAuth:
		transport := &spnego.Transport{}
		transport.TLSClientConfig = connection.TLSConfig
		httpClient = &http.Client{Transport: transport}
	default:
		transport := &http.Transport{TLSClientConfig: connection.TLSConfig}
		httpClient = &http.Client{Transport: transport}
	}

	var protocol string
	if connection.TLSConfig != nil {
		protocol = "https://"
	} else {
		protocol = "http://"
	}

	var base string
	if clientType == AdminRest {
		base = fmt.Sprintf("%s%s:%d/%s", protocol, connection.Host, connection.Port, Admin_Path)
	} else if clientType == ClientRest {
		base = fmt.Sprintf("%s%s:%d/%s", protocol, connection.Host, connection.Port, Client_Path)
	} else if clientType == ManageRest {
		base = fmt.Sprintf("%s%s:%d/%s", protocol, connection.Host, connection.Port, Manage_Path)
	}

	return httpClient, base
}
