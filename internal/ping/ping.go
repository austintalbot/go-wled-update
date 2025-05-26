package ping

import (
	"log"
	"time"

	probing "github.com/prometheus-community/pro-bing"
)

func PingDeviceOld(ip string) {
	for {
		pinger, err := probing.NewPinger(ip)
		if err != nil {
			log.Printf("Failed to initialize pinger for %s: %v", ip, err)
			time.Sleep(1 * time.Second)
			continue
		}
		pinger.Count = 1
		err = pinger.Run() // Blocks until finished.
		if err != nil {
			log.Printf("Host %s is unreachable: %v", ip, err)
		} else {
			stats := pinger.Statistics()
			if stats.PacketsRecv > 0 {
				log.Printf("Host %s is reachable", ip)
			} else {
				log.Printf("Host %s is unreachable", ip)
			}
		}
		time.Sleep(1 * time.Second) // Adjust the interval as needed
	}
}
