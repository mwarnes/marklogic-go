package Structures

type DocumentProperties struct {
	URI                string            `url:"uri,omitempty"`
	Categories         []string          `url:"category,omitempty"`
	Collections        []string          `url:"collection,omitempty"`
	Permissions        map[string]string `url:"perm,omitempty"`
	Properties         map[string]string `url:"perm,omitempty"`
	Value              map[string]string `url:"value,omitempty"`
	Extract            string            `url:"extract,omitempty"`
	Repair             string            `url:"repair,omitempty"`
	Transform          string            `url:"transform,omitempty"`
	Trans              map[string]string `url:"trans,omitempty"`
	ForestName         string            `url:"forest-name,omitempty"`
	Txid               string            `url:"txid,omitempty"`
	Lang               string            `url:"lang,omitempty"`
	TemporalCollection string            `url:"temporal-collection,omitempty"`
	TemporalDocument   string            `url:"temporal-document,omitempty"`
	SystemTime         string            `url:"system-time,omitempty"`
}

type DocumentResponse struct {
	Collections    []string       `json:"collections"`
	Permissions    Permissions    `json:"permissions"`
	Properties     Properties     `json:"properties"`
	Quality        int            `json:"quality"`
	MetadataValues MetadataValues `json:"metadataValues"`
}

type Permissions []struct {
	RoleName     string   `json:"role-name"`
	Capabilities []string `json:"capabilities"`
}

type Properties struct {
	Playtype string `json:"playtype"`
}

type MetadataValues struct {
	Mk1 string `json:"mk1"`
	Mk2 string `json:"mk2"`
}
