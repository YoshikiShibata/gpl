// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import "net"

func cmdPort(cmds []string, cc *clientConn) (net.Conn, error) {
	address, err := parsePORTAddress(cmds[1])
	if err != nil {
		return nil, err
	}
	return establishDataConnection(address, cc)
}
