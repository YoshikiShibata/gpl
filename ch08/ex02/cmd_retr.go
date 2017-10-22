// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func cmdRetr(cmds []string, cc *clientConn, dataConn net.Conn, transferType string) error {
	defer dataConn.Close()

	if err := cc.writeResponseCode(statusTransferStarting); err != nil {
		return err
	}

	f, err := os.Open(cmds[1])
	if err != nil {
		cc.writeResponse(statusActionNotTaken, err.Error())
		return err
	}

	defer f.Close()
	log.Printf("cmdRetr: start transfer")

	switch transferType {
	case typeASCII:
		ascii := asciiText{dataConn, nil}
		io.Copy(&ascii, f)
	case typeIMAGE:
		io.Copy(dataConn, f)
	default:
		return fmt.Errorf("Unknown transfer type : %s", transferType)
	}

	if err := cc.writeResponseCode(statusClosingDataConnection); err != nil {
		return err
	}
	log.Printf("cmdRetr: completed")
	return nil
}
