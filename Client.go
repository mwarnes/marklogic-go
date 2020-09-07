package marklogic

import (
	"fmt"
	"github.com/dpotapov/go-spnego"
	"github.com/mwarnes/digest"
	"io"
	"log"
	"net/http"
)

type RestService struct {
	Client Client
	base   string
}

func NewRestService(client Client, base string) *RestService {

	return &RestService{
		base:   base,
		Client: client,
	}
}

func (s *RestService) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, s.base+url, body)
	return req, err
}

func (s *RestService) ExecuteRequest(req *http.Request) (*http.Response, error) {
	s.Client = Decorate(s.Client,
		AddHeader("Accept", "application/json"),
	)

	response, err := s.Client.Do(req)
	if err != nil {
		return response, err
	}

	return response, err
}

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

type MLRestClient struct {
	RestService *RestService
}

func RestClient(connection Connection) MLRestClient {

	client, base := getBasicClientAndBase(connection, ClientRest)

	cli := Decorate(client,
		AddHeader("User-Agent", MlGoHttpVersion),
	)

	if connection.AuthenticationType == BasicAuth {
		cli = Decorate(client,
			AddBasicAuthentication(connection),
		)
	}

	restService := NewRestService(cli, base)

	return MLRestClient{
		RestService: restService,
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
	base = fmt.Sprintf("%s%s:%d", protocol, connection.Host, connection.Port)

	return httpClient, base
}
