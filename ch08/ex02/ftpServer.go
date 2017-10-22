// Copyright © 2016 Yoshiki Shibata. All rights reserved.

// A Simple FTP server. This server supports only the following FTP client commands:
// ls, pwd, cd, put, get, bin, ascii
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/textproto"
	"os"
	"strings"
)

const (
	statusTransferStarting      = 125
	statusCommandOk             = 200
	statusCommandNotImplemented = 202
	statusName                  = 215
	statusReady                 = 220
	statusLoggedOut             = 221
	statusClosingDataConnection = 226
	statusLoggedIn              = 230
	statusFileActionCompleted   = 250
	statusPathCreated           = 257

	statusUserOK                      = 331
	statusRequestedFileActionNotTaken = 450
	statusCommandNotImplemented502    = 502
	statusActionNotTaken              = 550
)

const (
	typeASCII = "A"
	typeIMAGE = "I"
)

const (
	welcomeMessage = "Welcome to FTP server written in Go (v0.0)"
)

func main() {
	fmt.Printf("Home = %s\n", os.Getenv("HOME"))
	if err := os.Chdir(os.Getenv("HOME")); err != nil {
		fmt.Printf("%v\n", err)
	}

	in, err := net.Listen("tcp", ":21")
	if err != nil {
		fmt.Printf("Listen: %v\n", err)
		return
	}
	for {
		conn, err := in.Accept()
		if err != nil {
			fmt.Printf("Accept: %v\n", err)
		}
		go handleConnection(conn)
	}
}

type clientConn struct {
	conn net.Conn
	r    *textproto.Reader
}

func newClientConn(conn net.Conn) *clientConn {
	var cc clientConn
	cc.conn = conn
	cc.r = textproto.NewReader(bufio.NewReader(conn))
	return &cc
}

func (cc *clientConn) writeResponse(code int, message string) error {
	var res string

	if message == "" {
		res = fmt.Sprintf("%d\n", code)
	} else {
		res = fmt.Sprintf("%d %s\n", code, message)
	}
	_, err := io.WriteString(cc.conn, res)
	return err
}

func (cc *clientConn) writeResponseCode(code int) error {
	return cc.writeResponse(code, "")
}

func (cc *clientConn) readLine() (string, error) {
	return cc.r.ReadLine()
}

func handleConnection(conn net.Conn) {
	fmt.Printf("Connected\n")
	cc := newClientConn(conn)
	err := cc.writeResponseCode(statusReady)
	if err != nil {
		log.Printf("%v", err)
		conn.Close()
		return
	}

	var dataConn net.Conn
	var transferType string
	cwd := newCwd()

	for {
		var line string
		if line, err = cc.readLine(); err != nil {
			if err == io.EOF {
				log.Printf("Disconnected\n")
				return
			}
			log.Printf("%v", err)
			return
		}
		fmt.Printf("%s\n", line)
		cmds := strings.Split(line, " ")

		switch cmds[0] {
		case "RETR":
			err = cwd.execute(func() error {
				return cmdRetr(cmds, cc, dataConn, transferType)
			})

			if err != nil {
				log.Printf("%v", err)
			}

		case "STOR":
			err = cwd.execute(func() error {
				return cmdStor(cmds, cc, dataConn, transferType)
			})

			if err != nil {
				log.Printf("%v", err)
			}

		case "TYPE":
			switch cmds[1] {
			case "I":
				transferType = typeIMAGE
				err = cc.writeResponseCode(statusCommandOk)
			case "A":
				transferType = typeASCII
				err = cc.writeResponseCode(statusCommandOk)
			default:
				fmt.Printf("Unspported Type(%s)\n", cmds[1])
				err = cc.writeResponseCode(statusCommandNotImplemented)
			}
			fmt.Printf("Current Type = %s\n", transferType)

		case "CWD":
			if err = cwd.changeCWD(cmds[1]); err != nil {
				err = cc.writeResponse(statusActionNotTaken, err.Error())
			} else {
				err = cc.writeResponseCode(statusFileActionCompleted)
			}
			if err != nil {
				log.Printf("%v", err)
			}

		case "USER": // ignore user
			err = cc.writeResponseCode(statusUserOK)
			if err != nil {
				log.Printf("%v", err)
			}

		case "PASS": // ignore password
			err = cc.writeResponse(statusLoggedIn, welcomeMessage)
			if err != nil {
				log.Printf("%v", err)
			}

		case "SYST":
			err = cc.writeResponse(statusName, "UNIX")
			if err != nil {
				log.Printf("%v", err)
			}

		case "PWD", "XPWD":
			pwd := cwd.pwd()
			log.Printf("pwd = %s", pwd)
			err = cc.writeResponse(statusPathCreated,
				fmt.Sprintf(`"%s" is the current directory`, pwd))
			if err != nil {
				log.Printf("%v", err)
			}

		case "PORT": // For ubuntu
			if dataConn, err = cmdPort(cmds, cc); err != nil {
				log.Printf("%v", err)
			}

		case "EPRT": // Foc Mac OS X
			if dataConn, err = cmdEprt(cmds, cc); err != nil {
				log.Printf("%v", err)
			}

		case "QUIT":
			err = cc.writeResponse(statusLoggedOut, "bye")
			if err != nil {
				log.Printf("%v", err)
			}

		case "LIST":
			err = cwd.execute(func() error {
				return cmdList(cmds, cc, dataConn)
			})

			if err != nil {
				log.Printf("%v", err)
			}

		case "NLST":
			err = cwd.execute(func() error {
				return cmdNlst(cmds, cc, dataConn, cwd.pwd())
			})

			if err != nil {
				log.Printf("%v", err)
			}

		case "FEAT", "EPSV", "LPSV", "LPRT":
			if err = cc.writeResponseCode(statusCommandNotImplemented502); err != nil {
				log.Printf("%v", err)
			}

		default:
			fmt.Printf("%v: Not Implemented Yet (%s)\n", cmds, line)
			err = cc.writeResponseCode(statusCommandNotImplemented)
			if err != nil {
				log.Printf("%v", err)
			}
		}
	}
}
