package main

import "net"

func cmdEprt(cmds []string, cc *clientConn) (net.Conn, error) {
	address, err := parseEPRTAddress(cmds[1])
	if err != nil {
		return nil, err
	}
	return establishDataConnection(address, cc)
}
