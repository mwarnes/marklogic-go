package Utils

import (
	"github.com/mwarnes/marklogic-go"
	"log"
	"time"
)

func checkServerRestart(client marklogic.MLRestClient, lastStartup string, nbretries int, retinterval int64) bool {
	time.Sleep(time.Duration(5) * time.Second)
	for i := 0; i <= nbretries; i++ {
		time.Sleep(time.Duration(retinterval) * time.Second)
		timestamp, _ := client.RestService.Timestamp()
		log.Println("Current startup timestamp=" + timestamp)
		if timestamp == "nil" || timestamp == "" {
			//return false
		} else {
			if timestamp != lastStartup {
				return true
			}
		}
	}
	return false
}
