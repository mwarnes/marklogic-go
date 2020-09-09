package main

import "github.com/mwarnes/marklogic-go/Utils"

func main() {

	var hosts []string
	hosts = append(
		hosts,
		"ml-node-1", "ml-node-2", "ml-node-3",
	)

	clusterParms := Utils.ClusterParameters{
		ClusterHosts: hosts,
		Group:        "Default",
		LicenseKey:   "",
		Licensee:     "",
		Username:     "admin",
		Password:     "admin",
		WalletPW:     "password",
		Realm:        "public",
	}

	Utils.InitialiseCluster(clusterParms)

}
