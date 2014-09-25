// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package seekio

import (
	"io"
	"io/ioutil"
	"os"
)

type Reader struct {
	r0 io.Reader
	rs io.ReadSeeker
	fp *os.File
}

func NewReader(r0 io.Reader, tempDir string) (r *Reader, err error) {
	if rs, ok := r0.(io.ReadSeeker); ok {
		r = &Reader{
			r0: r0,
			rs: rs,
		}
		return
	}

	// create temp file
	fp, err := ioutil.TempFile(tempDir, "seekio.Reader.")
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			defer os.Remove(fp.Name())
			fp.Close()
			r = nil
		}
	}()

	// copy r0 to temp file
	if _, err = io.Copy(fp, r0); err != nil {
		if err != io.EOF {
			return nil, err
		}
		err = nil
	}
	r = &Reader{
		r0: r0,
		rs: fp,
		fp: fp,
	}
	return
}

func (r *Reader) Seek(offset int64, whence int) (newoff int64, err error) {
	return r.rs.Seek(offset, whence)
}

func (r *Reader) Read(p []byte) (n int, err error) {
	return r.rs.Read(p)
}

func (r *Reader) Release() (err error) {
	if r.fp != nil {
		defer os.Remove(r.fp.Name())
		r.fp.Close()
	}
	*r = Reader{}
	return
}
