package main

import (
	"github.com/mwarnes/marklogic-go"
	"log"
	"net/http"
	"os"
	"time"
)

type Servers struct {
	ServerDefaultList struct {
		Meta struct {
			URI         string    `json:"uri"`
			CurrentTime time.Time `json:"current-time"`
			ElapsedTime struct {
				Units string  `json:"units"`
				Value float64 `json:"value"`
			} `json:"elapsed-time"`
		} `json:"meta"`
		Relations struct {
			RelationGroup []struct {
				Typeref       string `json:"typeref"`
				RelationCount struct {
					Units string `json:"units"`
					Value int    `json:"value"`
				} `json:"relation-count"`
				Relation []struct {
					Uriref  string `json:"uriref"`
					Idref   string `json:"idref"`
					Nameref string `json:"nameref"`
				} `json:"relation"`
			} `json:"relation-group"`
		} `json:"relations"`
		ListItems struct {
			ListCount struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"list-count"`
			ListItem []struct {
				RelationID   []string `json:"relation-id"`
				Groupnameref string   `json:"groupnameref"`
				Uriref       string   `json:"uriref"`
				Kindref      string   `json:"kindref"`
				ContentDb    string   `json:"content-db"`
				Idref        string   `json:"idref"`
				Nameref      string   `json:"nameref"`
				ModulesDb    string   `json:"modules-db,omitempty"`
			} `json:"list-item"`
		} `json:"list-items"`
		RelatedViews struct {
			RelatedView []struct {
				ViewType string `json:"view-type"`
				ViewName string `json:"view-name"`
				ViewURI  string `json:"view-uri"`
			} `json:"related-view"`
		} `json:"related-views"`
	} `json:"server-default-list"`
}

func main() {

	logger := log.New(os.Stderr, "", log.LstdFlags)

	conn := marklogic.Connection{
		Host:               "mwca",
		Port:               8002,
		Username:           "admin",
		Password:           "admin",
		AuthenticationType: marklogic.DigestAuth,
	}

	c := marklogic.RestClient(conn)

	c.RestService.Client = marklogic.Decorate(c.RestService.Client,
		marklogic.AddHeader("SomeHeader", "SomeValue"),
		marklogic.Logging(logger),
	)

	req, _ := c.RestService.NewRequest("GET", "/manage/v2/servers", nil)

	req.Header.Del("User-Agent")

	httpResp, err := c.RestService.ExecuteRequest(req)

	servers := new(Servers)

	if err == nil {
		defer httpResp.Body.Close()
		if httpResp.StatusCode == http.StatusOK {
			marklogic.DecodeResponseBodyJSON(httpResp, servers)
			log.Println(servers.ServerDefaultList)
		}
	} else {
		log.Println(err)
	}

}
