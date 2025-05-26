package updater

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func UpdateDevice(ip string, fw string) error {
	fwfile := fmt.Sprintf("./%s", fw)

	if _, err := exec.LookPath(fwfile); err != nil {
		log.Printf("Firmware file not found: %s", fwfile)
		return fmt.Errorf("firmware file not found: %w", err)
	}

	log.Printf("Pinging %s...", ip)
	if err := PingDevice(ip); err != nil {
		return err
	}

	log.Printf("Updating %s with %s...", ip, fw)
	// Simulate the update process
	if err := exec.Command("curl", "-s", "-F", fmt.Sprintf("update=@%s", fwfile), fmt.Sprintf("http://%s/update", ip)).Run(); err != nil {
		log.Printf("Update failed for %s: %v", ip, err)
		return fmt.Errorf("update failed for %s: %w", ip, err)
	}

	log.Printf("Update successful for %s", ip)
	return nil
}

func PingDevice(ip string) error {
	for {
		cmd := exec.Command("ping", "-c", "1", ip)
		if err := cmd.Run(); err != nil {
			log.Printf("Host %s is unreachable.", ip)
		} else {
			log.Printf("Host %s is reachable.", ip)
		}
		time.Sleep(5 * time.Second) // Adjust the interval as needed
	}
}
