// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"fmt"
	"net"
	"os"
)

func cmdNlst(cmds []string, cc *clientConn, dataConn net.Conn, cwd string) error {
	defer dataConn.Close()
	if err := cc.writeResponseCode(statusTransferStarting); err != nil {
		return err
	}

	var directory string

	if len(cmds) == 1 {
		directory = cwd
	} else {
		directory = cwd + string([]rune{os.PathSeparator}) + cmds[1]
	}

	if err := listFiles(directory, dataConn); err != nil {
		err2 := cc.writeResponseCode(statusRequestedFileActionNotTaken)
		if err2 != nil {
			return err2
		}
		return err
	}

	if err := cc.writeResponseCode(statusClosingDataConnection); err != nil {
		return err
	}
	return nil
}

func listFiles(directory string, dataConn net.Conn) error {
	fileInfo, err := os.Stat(directory)
	if err != nil {
		return err
	}

	if !fileInfo.IsDir() {
		return fmt.Errorf("%s not a directory", directory)
	}

	f, err := os.Open(directory)
	if err != nil {
		return err
	}
	defer f.Close()

	names, err := f.Readdirnames(0)
	if err != nil {
		return err
	}

	for _, name := range names {
		dataConn.Write([]byte(name))
		dataConn.Write([]byte("\n"))
	}
	return nil
}
