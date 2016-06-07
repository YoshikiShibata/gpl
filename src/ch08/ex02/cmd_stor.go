// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func cmdStor(cmds []string, cc *clientConn, dataConn net.Conn, transferType string) error {
	defer dataConn.Close()

	if err := cc.writeResponseCode(statusTransferStarting); err != nil {
		return err
	}

	f, err := os.Create(cmds[1])
	if err != nil {
		cc.writeResponse(statusActionNotTaken, err.Error())
		return err
	}

	defer f.Close()

	log.Printf("cmdStor: start transfer")

	switch transferType {
	case type_ASCII:
		ascii := asciiText{nil, dataConn}
		io.Copy(f, &ascii)
	case type_IMAGE:
		io.Copy(f, dataConn)
	default:
		return fmt.Errorf("Unknown transfer type : %s", transferType)
	}

	if err := cc.writeResponseCode(statusClosingDataConnection); err != nil {
		return err
	}
	log.Printf("cmdStor: completed")
	return nil
}
