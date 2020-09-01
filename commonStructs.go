package marklogic

import (
	"time"
)

type Meta struct {
	URI         string      `json:"uri"`
	CurrentTime time.Time   `json:"current-time"`
	ElapsedTime ElapsedTime `json:"elapsed-time"`
}

type ElapsedTime struct {
	Units string  `json:"units"`
	Value float64 `json:"value"`
}

type Relations struct {
	RelationGroup []RelationGroup `json:"relation-group"`
}

type RelationGroup struct {
	Uriref        string        `json:"uriref"`
	Typeref       string        `json:"typeref"`
	RelationCount RelationCount `json:"relation-count"`
	Relation      []Relation    `json:"relation"`
}

type Relation struct {
	Uriref  string `json:"uriref"`
	Roleref string `json:"roleref"`
	Idref   string `json:"idref"`
	Nameref string `json:"nameref"`
}

type RelationCount struct {
	Units string `json:"units"`
	Value int    `json:"value"`
}

type RelatedViews struct {
	RelatedView []RelatedView `json:"related-view"`
}

type RelatedView struct {
	ViewType string `json:"view-type"`
	ViewName string `json:"view-name"`
	ViewURI  string `json:"view-uri"`
}

type ListItems struct {
	ListCount ListCount  `json:"list-count"`
	ListItem  []ListItem `json:"list-item"`
}

type ListCount struct {
	Units string `json:"units"`
	Value int    `json:"value"`
}

type ListItem struct {
	Uriref  string `json:"uriref"`
	Idref   string `json:"idref"`
	Nameref string `json:"nameref"`
}

type ClusterRestart struct {
	Restart Restart `json:"restart"`
}

type Restart struct {
	LastStartup []LastStartup `json:"last-startup"`
	Link        Link          `json:"link"`
	Message     string        `json:"message"`
}

type LastStartup struct {
	HostID string    `json:"host-id"`
	Value  time.Time `json:"value"`
}

type Link struct {
	Kindref string `json:"kindref"`
	Uriref  string `json:"uriref"`
}

type RestartResponse struct {
	Restart struct {
		LastStartup []LastStartup `json:"last-startup"`
		Link        Link          `json:"link"`
		Message     string        `json:"message"`
	} `json:"restart"`
}

type RestErrorResponse struct {
	ErrorResponse ErrorResponse `json:"errorResponse"`
}

type ErrorResponse struct {
	StatusCode  string `json:"statusCode"`
	Status      string `json:"status"`
	MessageCode string `json:"messageCode"`
	Message     string `json:"message"`
}
