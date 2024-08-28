package main

import (
	"fmt"
	"net"
	"time"
)

func ScanPort(hostname string, port int) {
	address := fmt.Sprintf("%s:%d", hostname, port)
	_, err := net.DialTimeout("tcp", address, time.Second*5)
	if err != nil {
		fmt.Printf("Port %d  \tis closed\n", port)
		return
	}
	fmt.Printf("Port %d is open\n", port)
}

func main() {
	hostname := "http://scanme.nmap.org"
	portsToScan := [6]int{9090, 5432, 80, 443, 22, 8080}
	for _, port := range portsToScan {
		ScanPort(hostname, port)
	}
}
