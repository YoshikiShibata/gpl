// Copyright © 2016 Yoshiki Shibata

// Package bzip provides a writer that uses bzip2 compression (bzip.org).
package bzip

import (
	"io"
	"os/exec"
	"sync"
)

// + Exericse 13.4
type writer struct {
	sync.Mutex
	cmd *exec.Cmd
	w   io.WriteCloser
	wg  sync.WaitGroup
}

// NewWriter returns a writer for bzip2-compressed streams.
func NewWriter(out io.Writer) (io.WriteCloser, error) {
	var w writer
	w.cmd = exec.Command("/usr/bin/bzip2")
	stdout, err := w.cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stdin, err := w.cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	w.w = stdin

	if err := w.cmd.Start(); err != nil {
		return nil, err
	}
	w.wg.Add(1)
	go func() {
		io.Copy(out, stdout)
		w.wg.Done()
	}()

	return &w, nil
}

func (w *writer) Write(data []byte) (int, error) {
	w.Lock()
	defer w.Unlock()

	var total int // uncompressed bytes written

	for len(data) > 0 {
		n, err := w.w.Write(data)
		if err != nil {
			return total + n, err
		}
		total += n
		data = data[total:]
	}
	return total, nil
}

// Close flushes the compressed data and closes the stream.
// It does not close the underlying io.Writer.
func (w *writer) Close() error {
	w.Lock()
	defer w.Unlock()

	w.w.Close()
	w.wg.Wait()
	if err := w.cmd.Wait(); err != nil {
		return err
	}
	return nil
}

//- Exercise 13.4
