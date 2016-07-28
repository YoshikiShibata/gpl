// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package tar

import (
	"archive/tar"
	"ch10/ex02/unarchive"
	"io"
	"os"
)

const magicNumber = "ustar"
const offset = 257

func init() {
	unarchive.RegisterFormat("tar", magicNumber, offset+len(magicNumber), offset, decode)
}

type tarReader struct {
	file *os.File
	r    *tar.Reader
}

type tarFile struct {
	r *tar.Reader
	h *tar.Header
}

func (r *tarReader) Next() (unarchive.File, error) {
	h, err := r.r.Next()
	if err != nil {
		if err == io.EOF {
			r.file.Close()
			return nil, io.EOF
		}
		return nil, err
	}
	return &tarFile{r.r, h}, nil
}

func (f *tarFile) Name() string {
	return f.h.Name
}

func (f *tarFile) FileInfo() os.FileInfo {
	return f.h.FileInfo()
}

type readCloser struct {
	reader *tar.Reader
}

func (rc *readCloser) Close() error {
	return nil
}

func (rc *readCloser) Read(b []byte) (int, error) {
	return rc.reader.Read(b)
}

func (f *tarFile) Open() (io.ReadCloser, error) {
	return &readCloser{f.r}, nil
}

func decode(name string) (unarchive.Reader, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	r := tar.NewReader(f)
	return &tarReader{f, r}, nil
}
