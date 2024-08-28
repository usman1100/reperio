package main

import (
	"fmt"
	"net"
	"time"

	"github.com/fatih/color"
)

func ScanPort(hostname string, port int) {
	address := fmt.Sprintf("%s:%d", hostname, port)
	conn, err := net.DialTimeout("tcp", address, time.Second*5)
	if err != nil {
		fmt.Printf("Port %d \tis ", port)
		color.Red("closed\n")
		return
	}
	defer conn.Close()

	fmt.Printf("Port %d \tis ", port)
	color.Green("open\n")
}

func main() {
	hostname := "scanme.nmap.org"
	portsToScan := [6]int{9090, 5432, 80, 443, 22, 8080}

	color.Yellow("Scanning ports on %s\n", hostname)

	for _, port := range portsToScan {
		ScanPort(hostname, port)
	}
}
