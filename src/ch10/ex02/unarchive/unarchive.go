// Copyright © 2016 Yoshiki Shibata. All rights reserved.

// Package unarchive supports unarchiving a zip file or a tar file.

package unarchive

import (
	"io"
	"os"
)

// File represent a file or directory. Only a file can be opened.
type File interface {
	Name() string // the name of the file
	FileInfo() os.FileInfo
	Open() (io.Reader, error)
	Close()
}

// Reader returns a next File. If there is no next File, os.EOF will be
// returned as an error
type Reader interface {
	Next() (*File, error)
}

// OpenReader read the specified file and returns a Reader to access the
// contents of the archive file.
func OpenReader(name string) (Reader, error) {
	panic("Not Implemented Yet")
}

// RegisterFormat registers an archive format for use by Decode.
// Name is the name of the format like "zip" or "tar".
// Magic is the magic prefix that identifies the archive format.
func RegisterFormat(name, magic string, decode func(io.Reader) (Reader, error)) {
	panic("Not Implemented Yet")
}
