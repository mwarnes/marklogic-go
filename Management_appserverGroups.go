package marklogic

const (
	Groups = "LATEST/groups"
)

type GroupsList struct {
	GroupDefaultList struct {
		Meta         Meta         `json:"meta"`
		ListItems    ListItems    `json:"list-items"`
		RelatedViews RelatedViews `json:"related-views"`
	} `json:"group-default-list"`
}
