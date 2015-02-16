// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package seekio provides Seeker for Readers and Writers.

package tiff

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
)

var (
	_ io.ReadSeeker  = (*seekioReader)(nil)
	_ io.ReaderAt    = (*seekioReader)(nil)
	_ io.ReadCloser  = (*seekioReader)(nil)
	_ io.WriteSeeker = (*seekioWriter)(nil)
	_ io.WriterAt    = (*seekioWriter)(nil)
	_ io.WriteCloser = (*seekioWriter)(nil)
)

type seekioReader struct {
	r   io.Reader
	rs  io.ReadSeeker
	fp  *os.File
	buf []byte
	off int
	err error
}

func openSeekioReader(r io.Reader, maxBufferSize int) *seekioReader {
	if rs, ok := r.(io.ReadSeeker); ok {
		return &seekioReader{rs: rs}
	}
	data, err := ioutil.ReadAll(r)
	return &seekioReader{r: r, buf: data, err: err}
}

func (p *seekioReader) Read(data []byte) (n int, err error) {
	if p.err != nil {
		err = p.err
		return
	}
	if p.rs != nil {
		return p.rs.Read(data)
	}
	if p.off >= len(p.buf) { // Note len(nil)==0
		return 0, io.EOF
	}
	n = copy(data, p.buf[p.off:])
	p.off += n
	return
}

func (p *seekioReader) ReadAt(data []byte, off int64) (n int, err error) {
	if p.err != nil {
		err = p.err
		return
	}
	old, err := p.Seek(0, 1)
	if err != nil {
		return
	}
	newoff, err := p.Seek(off, 0)
	if err != nil {
		return
	}
	if newoff != off {
		err = errors.New("tiff: seekioReader.ReadAt, seek to new offset failed!")
		return
	}
	n, err = p.Read(data)
	if err != nil {
		return
	}
	if off, err = p.Seek(old, 0); off != old {
		err = errors.New("tiff: seekioReader.ReadAt, seek to old offset failed!")
		return
	}
	return
}

func (p *seekioReader) Seek(offset int64, whence int) (ret int64, err error) {
	if p.err != nil {
		err = p.err
		return
	}
	if p.rs != nil {
		return p.rs.Seek(offset, whence)
	}
	switch whence {
	case 0:
		ret = 0
	case 1:
		ret = int64(p.off)
	case 2:
		ret = int64(len(p.buf))
	default:
		return int64(p.off), io.EOF
	}
	ret += offset
	if ret >= int64(len(p.buf)) {
		err = io.EOF
		return
	}
	p.off = int(ret)
	return
}

func (p *seekioReader) Close() error {
	return p.err
}

type seekioWriter struct {
	w   io.Writer
	ws  io.WriteSeeker
	fp  *os.File
	buf []byte
	off int
	err error
}

func openSeekioWriter(w io.Writer, maxBufferSize int) (*seekioWriter, error) {
	if ws, ok := w.(io.WriteSeeker); ok {
		return &seekioWriter{ws: ws}, nil
	}
	return &seekioWriter{w: w}, nil
}

func (p *seekioWriter) Write(data []byte) (n int, err error) {
	if p.err != nil {
		err = p.err
		return
	}
	if p.ws != nil {
		n, err = p.ws.Write(data)
		return
	}
	p.grow(p.off + len(data))
	n = copy(p.buf[p.off:], data)
	p.off += n
	return
}

func (p *seekioWriter) WriteAt(data []byte, off int64) (n int, err error) {
	if p.err != nil {
		err = p.err
		return
	}
	old, err := p.Seek(0, 1)
	if err != nil {
		return
	}
	newoff, err := p.Seek(off, 0)
	if err != nil {
		return
	}
	if newoff != off {
		err = errors.New("tiff: seekioWriter.WriteAt, seek to new offset failed!")
		return
	}
	n, err = p.Write(data)
	if err != nil {
		return
	}
	if off, err = p.Seek(old, 0); off != old {
		err = errors.New("tiff: seekioWriter.WriteAt, seek to old offset failed!")
		return
	}
	return
}

func (p *seekioWriter) grow(n int) {
	if n > cap(p.buf) {
		buf := make([]byte, n, 2*cap(p.buf)+n)
		copy(buf, p.buf[0:p.off])
		p.buf = buf
	}
}

func (p *seekioWriter) Seek(offset int64, whence int) (ret int64, err error) {
	if p.err != nil {
		err = p.err
		return
	}
	if p.ws != nil {
		return p.ws.Seek(offset, whence)
	}
	switch whence {
	case 0:
		ret = 0
	case 1:
		ret = int64(p.off)
	case 2:
		ret = int64(len(p.buf))
	default:
		return int64(p.off), io.EOF
	}
	ret += offset
	if ret < 0 || int64(ret) != ret {
		return int64(p.off), io.EOF
	}
	p.off = int(ret)
	p.grow(p.off)
	return
}

func (p *seekioWriter) Close() error {
	if p.ws != nil {
		return nil
	}
	if _, err := p.w.Write(p.buf[:p.off]); err != nil {
		return err
	}
	*p = seekioWriter{}
	return nil
}
