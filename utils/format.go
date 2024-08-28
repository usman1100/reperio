package utils

import (
	"strconv"
	"strings"
)

func ParsePortsFromArgs(ports string) []int {
	// comma separated list of ports
	commaSeparatedPorts := strings.Split(ports, ",")
	numericPortList := []int{}

	for _, port := range commaSeparatedPorts {
		numericPort, err := strconv.Atoi(port)
		if err != nil {
			continue
		}

		numericPortList = append(numericPortList, numericPort)
	}

	return numericPortList
}
