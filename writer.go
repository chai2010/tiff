// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"image"
	"io"
)

type Writer struct {
	Writer io.WriteSeeker
	Header *Header
	Ifd    []*IFD
	Cfg    []image.Config
	Opt    []*Options

	ws *seekioWriter
}

func OpenWriter(w io.Writer, cfg []image.Config, opt []*Options) (p *Writer, err error) {
	return
}

func (p *Writer) ImageNum() int {
	return len(p.Ifd)
}

func (p *Writer) ImageConfig(idx int) image.Config {
	return p.Cfg[idx]
}

func (p *Writer) EncodeImage(idx int, m image.Image) (err error) {
	return
}

func (p *Writer) Close() (err error) {
	if p != nil && p.ws != nil {
		err = p.ws.Close()
	}
	*p = Writer{}
	return
}
