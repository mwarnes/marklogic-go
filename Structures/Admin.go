package Structures

type TimestampHead struct {
	StatusCode    int    `json:"status-code"`
	ServerDetails string `json:"server-details"`
	Connection    string `json:"connection"`
	KeepAlive     string `json:"keep-alive"`
}

type LicenseProperties struct {
	LicenseKey string `json:"license-key"`
	Licensee   string `json:"licensee"`
}

type SecurityProperties struct {
	AdminUsername string `url:"admin-username"`
	AdminPassword string `url:"admin-password"`
	Realm         string `url:"realm"`
}

type ClusterConfigProperties struct {
	Group        string `url:"group,omitempty"`
	ServerConfig string `url:"server-config,omitempty"`
	Zone         string `url:"server-config,omitempty"`
}
