// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// parsePORTAddress parses the address parameter of PORT command and
// returns an address which can be passed to net.Dial.
//
// A port command would be:
//
//	PORT h1,h2,h3,h4,p1,p2
//
// where h1 is the high order 8 bits of the internet host address.
// See RFC959
func parsePORTAddress(address string) (string, error) {
	ipInfo := strings.Split(address, ",")
	if len(ipInfo) != 6 {
		return "", fmt.Errorf("Unexpected address format: %v", ipInfo)
	}

	p1, err := strconv.Atoi(ipInfo[4])
	if err != nil {
		return "", err
	}

	p2, err := strconv.Atoi(ipInfo[5])
	if err != nil {
		return "", err
	}

	port := p1*256 + p2
	ipadd := fmt.Sprintf("%s.%s.%s.%s:%d",
		ipInfo[0], ipInfo[1], ipInfo[2], ipInfo[3], port)
	log.Printf("ip add [%s]\n", ipadd)

	return ipadd, nil
}

// parseEPRTAddress parses the address parameter of EPRT command and
// returns an address which can be passed to net.Dial
//
// EPRT<space><d><net-prt><d><net-addr><d><tcp-port><d>
// See RFC2428
func parseEPRTAddress(address string) (string, error) {
	ipInfo := strings.Split(address[1:len(address)-1], address[0:1])
	// protocol type, network address, port
	if len(ipInfo) != 3 {
		return "", fmt.Errorf("incorrect IP info: %v", ipInfo)
	}

	switch ipInfo[0] {
	case "1": // IPv4
		return ipInfo[1] + ":" + ipInfo[2], nil
	case "2": // IPv6
		return fmt.Sprintf("[%s]:%s", ipInfo[1], ipInfo[2]), nil
	default:
		return "", fmt.Errorf("Unknown Protocol: %s", ipInfo[0])
	}
}
