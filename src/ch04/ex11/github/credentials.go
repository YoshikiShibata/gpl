// Copyright © 2016 Yoshiki Shibata

package github

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

type Credentials struct {
	username string
	password string
}

func (c *Credentials) Query() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Username for 'https://github.com': ")
	username, _ := reader.ReadString('\n')
	c.username = strings.TrimSpace(username)

	fmt.Print(fmt.Sprintf("Password for 'https://%s@github.com': ", c.username))
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println("\nPassword Error: %v\n", err)
	}
	password := string(bytePassword)
	c.password = strings.TrimSpace(password)
}
