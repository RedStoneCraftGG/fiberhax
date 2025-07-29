package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func generatePassword(ssid string) (string, error) {
	if !strings.HasPrefix(ssid, "fh_") {
		return "", fmt.Errorf("Invalid SSID, must start with 'fh_'")
	}

	cleanSSID := ssid
	if strings.HasSuffix(ssid, "_5G") || strings.HasSuffix(ssid, "_5g") {
		cleanSSID = ssid[:len(ssid)-3]
	}

	parts := strings.Split(cleanSSID, "_")
	if len(parts) < 2 || len(parts[1]) != 6 {
		return "", fmt.Errorf("SSID format is incorrect: %s", ssid)
	}
	suffix := parts[1]

	num, err := strconv.ParseInt(suffix, 16, 64)
	if err != nil {
		return "", fmt.Errorf("Failed to parse hex from SSID: %v", err)
	}

	complement := 0xFFFFFF - num
	password := fmt.Sprintf("wlan%06x", complement)
	return password, nil
}

func main() {
	ssid := flag.String("ssid", "", "Default FiberHome SSID (example: fh_e12540 or fh_e12540_5G)")
	flag.Parse()

	if *ssid == "" {
		fmt.Println("Usage: fiberhax -ssid fh_xxxxxx[_5G]")
		os.Exit(1)
	}

	pass, err := generatePassword(*ssid)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("SSID     :", *ssid)
	fmt.Println("Password :", pass)
}
