package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/textproto"
	"os"
	"strconv"
	"strings"
)

const (
	StatusName        = 215
	StatusReady       = 220
	StatusLoggedIn    = 230
	StatusPathCreated = 257

	StatusUserOK = 331
)

func main() {
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
	err := cc.writeResponseCode(StatusReady)
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

	for {
		var line string
		if line, err = cc.readLine(); err != nil {
			log.Printf("%v", err)
			break
		}
		fmt.Printf("%s\n", line)
		cmds := strings.Split(line, " ")
		switch cmds[0] {
		case "USER":
			err = cc.writeResponseCode(StatusUserOK)
			if err != nil {
				log.Printf("%v", err)
			}
		case "PASS":
			err = cc.writeResponseCode(StatusLoggedIn)
			if err != nil {
				log.Printf("%v", err)
			}
		case "SYST":
			err = cc.writeResponse(StatusName, "UNIX")
			if err != nil {
				log.Printf("%v", err)
			}
		case "PWD":
			err = cc.writeResponse(StatusPathCreated, pwd)
			if err != nil {
				log.Printf("%v", err)
			}
		case "PORT":
			ipInfo := strings.Split(cmds[1], ",")
			if len(ipInfo) != 6 {
				log.Printf("Unexpected ipInfo: %v", ipInfo)
				continue
			}
			ph, err := strconv.Atoi(ipInfo[4])
			if err != nil {
				log.Printf("%v", err)
				continue
			}
			pl, err := strconv.Atoi(ipInfo[5])
			if err != nil {
				log.Printf("%v", err)
				continue
			}

			port := ph*256 + pl
			ipadd := fmt.Sprintf("%s.%s.%s.%s:%d",
				ipInfo[0], ipInfo[1], ipInfo[2], ipInfo[3], port)
			fmt.Printf("ip add [%s]\n", ipadd)

		default:
			fmt.Printf("%v: Not Implemented Yet (%s)\n", cmds, line)
		}
	}
}
