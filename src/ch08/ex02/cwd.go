// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"os"
	"sync"
)

var cwdLock sync.Mutex

// cwd is used to keep track of the current directory per a login user
// While executing some command such as List, Get, Put, Pwd, the current
// directory for the user must be restored.
type cwd struct {
	current string
}

// newCwd creates an instance whose current directory is the HOME directory
func newCwd() *cwd {
	return &cwd{os.Getenv("HOME")}
}

// changeCWD changes the current directory to the specified directory
func (c *cwd) changeCWD(dir string) error {
	cwdLock.Lock()
	defer cwdLock.Unlock()

	if err := os.Chdir(c.current); err != nil {
		return err
	}

	if err := os.Chdir(dir); err != nil {
		return err
	}

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	c.current = pwd
	return nil
}

// pwd returns the current directory.
func (c *cwd) pwd() string {
	return c.current
}

// execute run the passed fuction after restoring the current directory
func (c *cwd) execute(f func() error) error {
	cwdLock.Lock()
	defer cwdLock.Unlock()

	if err := os.Chdir(c.current); err != nil {
		return err
	}

	return f()
}
