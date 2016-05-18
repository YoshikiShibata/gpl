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
	"strconv"
	"strings"
)

const (
	StatusTransferStarting      = 125
	StatusCommandOk             = 200
	StatusCommandNotImplemented = 202
	StatusName                  = 215
	StatusReady                 = 220
	StatusLoggedOut             = 221
	StatusClosingDataConnection = 226
	StatusLoggedIn              = 230
	StatusPathCreated           = 257

	StatusUserOK = 331
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
			err = cc.writeResponseCode(StatusUserOK)
			if err != nil {
				log.Printf("%v", err)
			}
		case "PASS":
			err = cc.writeResponse(StatusLoggedIn, welcomeMessage)
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
			dataConn, err = net.Dial("tcp", ipadd)
			if err != nil {
				log.Printf("%v", err)
				continue
			}
			log.Printf("Data Connect established\n")
			err = cc.writeResponseCode(StatusCommandOk)
			if err != nil {
				log.Printf("%v", err)
			}

		case "QUIT":
			err = cc.writeResponse(StatusLoggedOut, "bye")
			if err != nil {
				log.Printf("%v", err)
			}

		case "LIST":
			if err := cc.writeResponseCode(StatusTransferStarting); err != nil {
				log.Printf("%v", err)
			}

			if len(cmds) == 1 {
				execls(nil, dataConn)
			} else {
				execls(cmds[1:], dataConn)
			}

			if err := cc.writeResponseCode(StatusClosingDataConnection); err != nil {
				log.Printf("%v", err)
			}
			dataConn.Close()
			dataConn = nil

		default:
			fmt.Printf("%v: Not Implemented Yet (%s)\n", cmds, line)
			err = cc.writeResponseCode(StatusCommandNotImplemented)
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

	go io.Copy(conn, stdout)
	go io.Copy(conn, stderr)

	if err := cmd.Start(); err != nil {
		conn.Write([]byte(err.Error()))
		return
	}

	cmd.Wait()
	log.Printf("execls done\n")
}
