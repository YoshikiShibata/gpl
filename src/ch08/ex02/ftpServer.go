// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/textproto"
	"os"
	"os/exec"
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
	statusPathCreated           = 257

	statusUserOK                   = 331
	statusCommandNotImplemented502 = 502
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

	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("%v", err)
		pwd = "/"
	}

	var dataConn net.Conn

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
		case "USER":
			err = cc.writeResponseCode(statusUserOK)
			if err != nil {
				log.Printf("%v", err)
			}
		case "PASS":
			err = cc.writeResponse(statusLoggedIn, welcomeMessage)
			if err != nil {
				log.Printf("%v", err)
			}
		case "SYST":
			err = cc.writeResponse(statusName, "UNIX")
			if err != nil {
				log.Printf("%v", err)
			}
		case "PWD":
			log.Printf("pwd = %s", pwd)
			err = cc.writeResponse(statusPathCreated,
				fmt.Sprintf(`"%s" is the current directory`, pwd))
			if err != nil {
				log.Printf("%v", err)
			}
		case "PORT":
			if dataConn, err = cmdPort(cmds, cc); err != nil {
				log.Printf("%v", err)
			}

		case "EPRT":
			if dataConn, err = cmdEprt(cmds, cc); err != nil {
				log.Printf("%v", err)
			}

		case "QUIT":
			err = cc.writeResponse(statusLoggedOut, "bye")
			if err != nil {
				log.Printf("%v", err)
			}

		case "LIST":
			if err := cc.writeResponseCode(statusTransferStarting); err != nil {
				log.Printf("%v", err)
			}

			if len(cmds) == 1 {
				execls(nil, dataConn)
			} else {
				execls(cmds[1:], dataConn)
			}

			if err := cc.writeResponseCode(statusClosingDataConnection); err != nil {
				log.Printf("%v", err)
			}
			dataConn.Close()
			dataConn = nil

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

	ascii := asciiText{conn}
	go io.Copy(&ascii, stdout)
	go io.Copy(&ascii, stderr)

	if err := cmd.Start(); err != nil {
		conn.Write([]byte(err.Error()))
		return
	}

	cmd.Wait()
	log.Printf("execls done\n")
}
