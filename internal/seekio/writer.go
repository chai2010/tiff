// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package seekio

import (
	"io"
	"io/ioutil"
	"os"
)

type Writer struct {
	w0 io.Writer
	ws io.WriteSeeker
	fp *os.File
}

func NewWriter(w0 io.Writer, tempDir string) (w *Writer, err error) {
	if ws, ok := w0.(io.WriteSeeker); ok {
		w = &Writer{
			w0: w0,
			ws: ws,
		}
		return
	}
	fp, err := ioutil.TempFile(tempDir, "seekio.Writer.")
	if err != nil {
		return nil, err
	}
	w = &Writer{
		w0: w0,
		ws: fp,
		fp: fp,
	}
	return
}

func (w *Writer) Seek(offset int64, whence int) (newoff int64, err error) {
	return w.ws.Seek(offset, whence)
}

func (w *Writer) Write(p []byte) (n int, err error) {
	return w.ws.Write(p)
}

func (w *Writer) Release() (err error) {
	if w.fp != nil {
		defer os.Remove(w.fp.Name())
		defer w.fp.Close()

		if _, err = w.fp.Seek(0, 0); err != nil {
			return
		}
		if _, err = io.Copy(w.w0, w.fp); err != nil {
			if err != io.EOF {
				return
			}
			err = nil
		}
	}
	*w = Writer{}
	return
}
