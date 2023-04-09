package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	var ipAddress string
	fmt.Print("Enter an IP address to scan: ")
	fmt.Scan(&ipAddress)

	var scanRate int
	fmt.Print("Enter scan rate (ports per second): ")
	fmt.Scan(&scanRate)

	fmt.Println("Scanning ports for IP address:", ipAddress)

	for port := 1; port <= 65535; port++ {
		target := ipAddress + ":" + strconv.Itoa(port)
		conn, err := net.DialTimeout("tcp", target, time.Duration(1)*time.Second)
		if err != nil {
			continue
		}
		conn.Close()

		service, _ := net.LookupService("tcp", strconv.Itoa(port))
		fmt.Println("Port", port, "open - Service:", service)
		time.Sleep(time.Duration(1000/scanRate) * time.Millisecond)
	}
}
