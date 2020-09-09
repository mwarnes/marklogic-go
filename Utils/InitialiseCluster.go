package Utils

import (
	"github.com/mwarnes/marklogic-go"
	"github.com/mwarnes/marklogic-go/Structures"
	"log"
	"time"
)

type ClusterParameters struct {
	ClusterHosts []string
	Group        string
	LicenseKey   string
	Licensee     string
	Username     string
	Password     string
	WalletPW     string
	Realm        string
}

func InitialiseCluster(cp ClusterParameters) {

	var bootstrapHost = cp.ClusterHosts[0]
	log.Printf("Bootstrap host: %s\n", bootstrapHost)
	var joiningHosts = []string{}
	for i := 1; i < len(cp.ClusterHosts); i++ {
		joiningHosts = append(joiningHosts, cp.ClusterHosts[i])

	}
	log.Printf("Cluster hosts: %v\n", joiningHosts)

	// Check if all hosts are available
	log.Println("Check all MarkLogic servers are running.")
	time.Sleep(time.Duration(5) * time.Second)
	for i := 0; i < len(cp.ClusterHosts); i++ {

		conn := marklogic.Connection{
			Host:               cp.ClusterHosts[i],
			Port:               8001,
			AuthenticationType: marklogic.None,
		}
		client := marklogic.RestClient(conn)

		timestamp, resp := client.RestService.Timestamp()
		if resp.StatusCode != 200 {
			log.Fatal(resp.Status)
		}
		log.Printf("%v timestamp: %s\n", cp.ClusterHosts[i], timestamp)
	}

	// Initialise Bootstrap host
	log.Printf("Initialise bootstrap host %s license.\n", bootstrapHost)

	conn := marklogic.Connection{
		Host:               bootstrapHost,
		Port:               8001,
		AuthenticationType: marklogic.None,
	}

	client := marklogic.RestClient(conn)

	lic := Structures.LicenseProperties{
		LicenseKey: cp.LicenseKey,
		Licensee:   cp.Licensee,
	}

	// Initialize MarkLogic server
	timestamp, _ := client.RestService.Timestamp()
	restartResp, errorResp, resp := client.RestService.Init(lic)

	if resp.StatusCode == 204 {
		log.Println("License modified no restart required.")
	} else if resp.StatusCode == 202 {
		log.Println("License inserted restart required.")
		log.Println("Last Timestamp: ", restartResp.Restart.LastStartup[0].Value)
		result := checkServerRestart(client, timestamp, 15, 2)
		if result {
			log.Printf("Bootstrap Server [%s] restarted\n", bootstrapHost)
		} else {
			log.Printf("Bootstrap Server restart [%s] failed\n", bootstrapHost)
		}
	} else {
		log.Println(errorResp.ErrorResponse.Message)
	}

	conn = marklogic.Connection{
		Host:               bootstrapHost,
		Port:               8001,
		Username:           cp.Username,
		Password:           cp.Password,
		AuthenticationType: marklogic.DigestAuth,
	}
	client = marklogic.RestClient(conn)

	secProps := Structures.SecurityProperties{
		AdminUsername: cp.Username,
		AdminPassword: cp.Password,
		Realm:         cp.Realm,
	}

	// Initialize MarkLogic server Security database.
	timestamp, _ = client.RestService.Timestamp()
	restartResp, errorResp, resp = client.RestService.InstanceAdmin(secProps)

	if resp.StatusCode == 202 {
		log.Println("Security initialised restart required.")
		log.Println("Last Timestamp: ", restartResp.Restart.LastStartup[0].Value)
		result := checkServerRestart(client, timestamp, 15, 2)
		if result {
			log.Printf("Bootstrap Server [%s] restarted\n", bootstrapHost)
		} else {
			log.Printf("Bootstrap Server restart [%s] failed\n", bootstrapHost)
		}
	} else {
		log.Fatalln(resp.Status)
	}

	// Add each joining host to the cluster
	for i := 0; i < len(joiningHosts); i++ {

		// Create AdminClient for initialising the cluster hosts
		log.Printf("Initialise cluster host %s license.\n", joiningHosts[i])

		conn := marklogic.Connection{
			Host:               joiningHosts[i],
			Port:               8001,
			AuthenticationType: marklogic.None,
		}
		client := marklogic.RestClient(conn)

		lic := Structures.LicenseProperties{
			LicenseKey: cp.LicenseKey,
			Licensee:   cp.Licensee,
		}

		// Initialize MarkLogic server
		timestamp, _ = client.RestService.Timestamp()
		restartResp, errorResp, resp := client.RestService.Init(lic)

		if resp.StatusCode == 204 {
			log.Println("License modified no restart required.")
		} else if resp.StatusCode == 202 {
			log.Println("License inserted restart required.")
			log.Println("Last Timestamp: ", restartResp.Restart.LastStartup[0].Value)
			result := checkServerRestart(client, timestamp, 15, 2)
			if result {
				log.Printf("server [%s] restarted\n", joiningHosts[i])
			} else {
				log.Printf("server restart [%s] failed\n", joiningHosts[i])
			}
		} else {
			log.Println(errorResp.ErrorResponse.Message)
		}

		// Get joining host server configuration
		log.Printf("Get %s server configuration.\n", joiningHosts[i])

		conn = marklogic.Connection{
			Host:               joiningHosts[i],
			Port:               8001,
			Username:           cp.Username,
			Password:           cp.Password,
			AuthenticationType: marklogic.BasicAuth,
		}
		client = marklogic.RestClient(conn)

		// Get MarkLogic Server configuration.
		configuration, resp := client.RestService.GetServerConfig()

		if resp.StatusCode == 200 {
			log.Println("Successfully retrieved MarkLogic server configuration.")
		} else {
			log.Println("Error retrieving configuration.")
		}

		// Send joining cluster configuration to bootstrap host
		log.Printf("Sending %s configuration to %s.\n", joiningHosts[i], bootstrapHost)

		clusterConfiguration := Structures.ClusterConfigProperties{
			Group:        cp.Group,
			ServerConfig: configuration,
		}

		conn = marklogic.Connection{
			Host:               bootstrapHost,
			Port:               8001,
			Username:           cp.Username,
			Password:           cp.Password,
			AuthenticationType: marklogic.DigestAuth,
		}
		client = marklogic.RestClient(conn)

		zipConfig, resp := client.RestService.SendClusterConfigForm(clusterConfiguration)

		// Send  cluster configuration to Joining host
		log.Printf("Sending cluster configuration to %s.\n", joiningHosts[i])

		conn = marklogic.Connection{
			Host:               joiningHosts[i],
			Port:               8001,
			AuthenticationType: marklogic.None,
		}
		client = marklogic.RestClient(conn)

		restartResp, errorResp, resp = client.RestService.SendClusterConfigZip(zipConfig)

		if resp.StatusCode == 200 {
			log.Printf("%s added to cluster, no restart required.\n", joiningHosts[i])
		} else if resp.StatusCode == 202 {
			log.Printf("%s added to cluster, restart required.\n", joiningHosts[i])
		} else {
			log.Fatalln(errorResp.ErrorResponse.Message)
		}

	}

	log.Println("Cluster initialised successfully.")

}
