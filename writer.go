// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"image"
	"io"
)

type _Writer struct {
	Writer io.WriteSeeker
	Header *Header
	Ifd    []*IFD
	Cfg    []image.Config
	Opt    []*Options

	ws *seekioWriter
}

func _OpenWriter(w io.Writer, cfg []image.Config, opt []*Options) (p *_Writer, err error) {
	return
}

func (p *_Writer) ImageNum() int {
	return len(p.Ifd)
}

func (p *_Writer) ImageConfig(idx int) image.Config {
	return p.Cfg[idx]
}

func (p *_Writer) EncodeImage(idx int, m image.Image) (err error) {
	return
}

func (p *_Writer) Close() (err error) {
	if p != nil && p.ws != nil {
		err = p.ws.Close()
	}
	*p = _Writer{}
	return
}
