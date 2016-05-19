// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func cmdPort(cmds []string, cc *clientConn) (dataConn net.Conn, err error) {

	ipInfo := strings.Split(cmds[1], ",")
	if len(ipInfo) != 6 {
		return nil, fmt.Errorf("Unexpected ipInfo: %v", ipInfo)
	}

	ph, err := strconv.Atoi(ipInfo[4])
	if err != nil {
		return nil, err
	}

	pl, err := strconv.Atoi(ipInfo[5])
	if err != nil {
		return nil, err
	}

	port := ph*256 + pl
	ipadd := fmt.Sprintf("%s.%s.%s.%s:%d",
		ipInfo[0], ipInfo[1], ipInfo[2], ipInfo[3], port)
	log.Printf("ip add [%s]\n", ipadd)
	dataConn, err = net.Dial("tcp", ipadd)
	if err != nil {
		return nil, err
	}
	log.Printf("Data Connect established\n")
	if err = cc.writeResponseCode(StatusCommandOk); err != nil {
		dataConn.Close()
		return nil, err
	}

	return dataConn, nil
}
