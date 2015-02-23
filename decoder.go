// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"image"
	"io"
)

// DecodeConfig returns the color model and dimensions of a TIFF image without
// decoding the entire image.
func DecodeConfig(r io.Reader) (cfg image.Config, err error) {
	var p *Reader
	if p, err = OpenReader(r); err != nil {
		return
	}
	defer p.Close()

	cfg = p.ImageConfig(0)
	return
}

func DecodeConfigAll(r io.Reader) (cfg []image.Config, err error) {
	var p *Reader
	if p, err = OpenReader(r); err != nil {
		return
	}
	defer p.Close()

	cfg = append(cfg, p.Cfg...)
	return
}

// Decode reads a TIFF image from r and returns it as an image.Image.
// The type of Image returned depends on the contents of the TIFF.
func Decode(r io.Reader) (m image.Image, err error) {
	var p *Reader
	if p, err = OpenReader(r); err != nil {
		return
	}
	defer p.Close()

	m, err = p.DecodeImage(0)
	return
}

func DecodeAll(r io.Reader) (m []image.Image, err error) {
	var p *Reader
	if p, err = OpenReader(r); err != nil {
		return
	}
	defer p.Close()

	m = make([]image.Image, p.ImageNum())
	for i := 0; i < p.ImageNum(); i++ {
		if m[i], err = p.DecodeImage(i); err != nil {
			return
		}
	}
	return
}
