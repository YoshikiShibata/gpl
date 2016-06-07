// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"os"
	"sync"
)

var cwdLock sync.Mutex

type cwd struct {
	current string
}

func newCwd() *cwd {
	return &cwd{os.Getenv("HOME")}
}

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

func (c *cwd) pwd() string {
	return c.current
}

func (c *cwd) execute(f func() error) error {
	cwdLock.Lock()
	defer cwdLock.Unlock()

	if err := os.Chdir(c.current); err != nil {
		return err
	}

	return f()
}
