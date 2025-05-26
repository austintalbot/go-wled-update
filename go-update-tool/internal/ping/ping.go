package ping

import (
	"log"
	"time"

	"github.com/charmbracelet/huh"
)

func PingDevice(ip string) {
	for {
		_, err := huh.Ping(ip)
		if err != nil {
			log.Printf("Host %s is unreachable: %v", ip, err)
		} else {
			log.Printf("Host %s is reachable", ip)
		}
		time.Sleep(1 * time.Second) // Adjust the interval as needed
	}
}
