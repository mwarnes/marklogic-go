package main

import (
	"encoding/json"
	"github.com/mwarnes/marklogic-go"
	"log"
	"net/http"
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

func decodeResponseBodyJSON(resp *http.Response, v interface{}) error {
	return json.NewDecoder(resp.Body).Decode(v)
}

func main() {

	conn := marklogic.Connection{
		Host:               "mwca",
		Port:               8002,
		Username:           "admin",
		Password:           "admin",
		AuthenticationType: marklogic.DigestAuth,
	}

	c := marklogic.MarkLogicManageClient(conn)

	req, _ := c.Adhoc.NewRequest("GET", "/v2/servers", nil)

	httpResp, err := c.Adhoc.ExecuteRequest(req)

	servers := new(Servers)

	if err == nil {
		defer httpResp.Body.Close()
		if httpResp.StatusCode == http.StatusOK {
			decodeResponseBodyJSON(httpResp, servers)
			log.Println(servers.ServerDefaultList.ListItems.ListItem[0].Uriref)
		}
	} else {
		log.Println(err)
	}

}