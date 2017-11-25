// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package zip

import (
	"archive/zip"
	"io"
	"os"

	"github.com/YoshikiShibata/gpl/ch10/ex02/unarchive"
)

const magicNumber = "PK\003\004"

func init() {
	unarchive.RegisterFormat("zip", magicNumber, len(magicNumber), 0, decode)
}

type zipReader struct {
	readCloser *zip.ReadCloser
	current    int
}

type zipFile struct {
	file *zip.File
}

func (r *zipReader) Next() (unarchive.File, error) {
	r.current++
	if len(r.readCloser.File) <= r.current {
		r.readCloser.Close()
		return nil, io.EOF
	}
	return &zipFile{r.readCloser.File[r.current]}, nil
}

func (f *zipFile) Name() string {
	return f.file.Name
}

func (f *zipFile) FileInfo() os.FileInfo {
	return f.file.FileInfo()
}

func (f *zipFile) Open() (io.ReadCloser, error) {
	return f.file.Open()
}

func decode(name string) (unarchive.Reader, error) {
	r, err := zip.OpenReader(name)
	if err != nil {
		return nil, err
	}

	return &zipReader{r, -1}, nil
}
