// Copyright © 2016 Yoshiki Shibata. All rights reserved.

// Package unarchive supports unarchiving a zip file or a tar file.
package unarchive

import (
	"bufio"
	"errors"
	"io"
	"os"
)

// ErrFormat indicates that decoding encountered an unknown format
var ErrFormat = errors.New("unarchive: unknown format")

// File represent a file or directory. Only a file can be opened.
type File interface {
	Name() string // the name of the file
	FileInfo() os.FileInfo
	Open() (io.ReadCloser, error)
}

// Reader returns a next File. If there is no next File, os.EOF will be
// returned as an error
type Reader interface {
	Next() (File, error)
}

// A format haolds an archive format's name, magice header and how to unarchive it.
type format struct {
	name, magic      string
	peekSize, offset int
	decode           func(string) (Reader, error)
}

// Formats is the list of registered formats
var formats []format

// A reader is an io.Reader that can also peek ahead
type reader interface {
	io.Reader
	Peek(int) ([]byte, error)
}

// as Reader converts an io.Reader to a reader
func asReader(r io.Reader) reader {
	if rr, ok := r.(reader); ok {
		return rr
	}
	return bufio.NewReader(r)
}

// match reports whether magic matches b.
func match(magic string, b []byte) bool {
	if len(magic) != len(b) {
		return false
	}
	for i, c := range b {
		if magic[i] != c {
			return false
		}
	}
	return true
}

// sniff determines the format of r's data.
func sniff(r reader) format {
	for _, f := range formats {
		b, err := r.Peek(f.peekSize)
		if err == nil && match(f.magic, b[f.offset:]) {
			return f
		}
	}
	return format{}
}

// OpenReader read the specified file and returns a Reader to access the
// contents of the archive file.
func OpenReader(name string) (Reader, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := asReader(file)
	f := sniff(r)
	if f.decode == nil {
		return nil, ErrFormat
	}

	return f.decode(name)
}

// RegisterFormat registers an archive format for use by Decode.
// Name is the name of the format like "zip" or "tar".
// Magic is the magic prefix that identifies the archive format.
// PeekSize is the number of bytes to be peeked.
// offset is the offset from where Magic is checked
func RegisterFormat(name, magic string, peekSize, offset int, decode func(string) (Reader, error)) {
	formats = append(formats, format{name, magic, peekSize, offset, decode})
}
