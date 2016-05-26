// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"log"
	"net"
)

func establishDataConnection(address string, cc *clientConn) (net.Conn, error) {
	dataConn, err := net.Dial("tcp", address)
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
