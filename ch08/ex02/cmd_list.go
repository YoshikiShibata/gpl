// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func cmdList(cmds []string, cc *clientConn, dataConn net.Conn) error {
	if err := cc.writeResponseCode(statusTransferStarting); err != nil {
		return err
	}

	if len(cmds) == 1 {
		execls(nil, dataConn)
	} else {
		execls(cmds[1:], dataConn)
	}

	defer dataConn.Close()

	if err := cc.writeResponseCode(statusClosingDataConnection); err != nil {
		return err
	}
	return nil
}

func execls(params []string, conn net.Conn) {
	cmd := exec.Command("/bin/ls", params...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		conn.Write([]byte(err.Error()))
		return
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		conn.Write([]byte(err.Error()))
		return
	}

	if conn == nil {
		panic("Data connection has not been established")
	}

	ascii := asciiText{conn, nil}
	go io.Copy(&ascii, stdout)
	go io.Copy(&ascii, stderr)

	if err := cmd.Start(); err != nil {
		conn.Write([]byte(err.Error()))
		return
	}

	cmd.Wait()
	log.Printf("execls done\n")
}
