package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/usman1100/reperio/utils"
)

func ScanPort(wg *sync.WaitGroup, hostname string, port int) {
	defer wg.Done()
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
	host := flag.String("host", "localhost", "hostname to scan")
	ports := flag.String("ports", "80,443,22,8080", "comma separated list of ports to scan")
	flag.Usage = func() {
		fmt.Println("Usage: reperio -host <hostname> -ports <ports>")
		flag.PrintDefaults()
	}
	flag.Parse()

	var wg sync.WaitGroup

	portsToScan := utils.ParsePortsFromArgs(*ports)

	color.Yellow("Scanning ports on %s\n", *host)

	for _, port := range portsToScan {
		wg.Add(1)
		go ScanPort(&wg, *host, port)
	}

	wg.Wait()
}
