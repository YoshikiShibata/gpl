package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func cmdEprt(cmds []string, cc *clientConn) (dataConn net.Conn, err error) {
	param := cmds[1]
	ipInfo := strings.Split(param[1:len(param)-1], param[0:1])
	// protocol type, network address, port
	if len(ipInfo) != 3 {
		log.Printf("incorrect IP info: %v\n", ipInfo)
	}

	var ipadd string
	switch ipInfo[0] {
	case "1": // IPv4
		ipadd = ipInfo[1] + ":" + ipInfo[2]
	case "2": // IPv6
		ipadd = fmt.Sprintf("[%s]:%s", ipInfo[1], ipInfo[2])
	default:
		panic(fmt.Errorf("Unknown Protocol: %s", ipInfo[0]))
	}

	log.Printf("ip add : %s\n", ipadd)
	dataConn, err = net.Dial("tcp", ipadd)
	if err != nil {
		return nil, err
	}
	log.Printf("Data Connect established\n")
	if err = cc.writeResponseCode(statusCommandOk); err != nil {
		dataConn.Close()
		return nil, err
	}

	return dataConn, nil
}
