package Structures

// Clusters
type LocalClusterResponse struct {
	LocalClusterDefault LocalClusterDefault `json:"local-cluster-default"`
}

type LocalClusterDefault struct {
	ID               string       `json:"id"`
	Name             string       `json:"name"`
	Version          string       `json:"version"`
	EffectiveVersion int          `json:"effective-version"`
	Role             string       `json:"role"`
	Meta             Meta         `json:"meta"`
	Relations        Relations    `json:"relations"`
	RelatedViews     RelatedViews `json:"related-views"`
}

type ClusterProperties struct {
	ClusterID                  string          `json:"cluster-id,omitempty"`
	ClusterName                string          `json:"cluster-name,omitempty"`
	Role                       string          `json:"role,omitempty"`
	Version                    string          `json:"version,omitempty"`
	EffectiveVersion           int             `json:"effective-version,omitempty"`
	SecurityVersion            int             `json:"security-version,omitempty"`
	SslFipsEnabled             bool            `json:"ssl-fips-enabled,omitempty"`
	XdqpSslCertificate         string          `json:"xdqp-ssl-certificate,omitempty"`
	BootstrapHost              []BootstrapHost `json:"bootstrap-host,omitempty"`
	DataDirectory              string          `json:"data-directory,omitempty"`
	FilesystemDirectory        string          `json:"filesystem-directory,omitempty"`
	OpsdirectorLogLevel        string          `json:"opsdirector-log-level,omitempty"`
	OpsdirectorMetering        string          `json:"opsdirector-metering,omitempty"`
	OpsdirectorConfig          string          `json:"opsdirector-config,omitempty"`
	OpsdirectorSessionEndpoint interface{}     `json:"opsdirector-session-endpoint,omitempty"`
}

type BootstrapHost struct {
	BootstrapHostID      string `json:"bootstrap-host-id"`
	BootstrapHostName    string `json:"bootstrap-host-name"`
	BootstrapConnectPort int    `json:"bootstrap-connect-port"`
}

// Perform a local cluster operation
// Valid operations are:
// . "restart-local-cluster"
// . "commit-upgrade-local-cluster"
// . "security-database-upgrade-local-cluster"
type ClusterOperation struct {
	Operation string `json:"operation,omitempty"`
}
