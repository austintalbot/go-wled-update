package main

import (
	"fmt"
	"log"
	"os"

	"go-update-tool/internal/updater"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s <ip> <firmware>", os.Args[0])
	}

	ip := os.Args[1]
	firmware := os.Args[2]

	fmt.Printf("Starting update process for device at %s with firmware %s...\n", ip, firmware)

	err := updater.UpdateDevice(ip, firmware)
	if err != nil {
		log.Fatalf("Update failed: %v", err)
	}

	fmt.Println("Update process completed successfully.")
}
