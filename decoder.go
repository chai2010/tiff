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

	return p.ImageConfig(0, 0)
}

func DecodeConfigAll(r io.Reader) (cfg [][]image.Config, errors [][]error, err error) {
	var p *Reader
	if p, err = OpenReader(r); err != nil {
		return
	}
	defer p.Close()

	cfg = make([][]image.Config, len(p.Ifd))
	errors = make([][]error, len(p.Ifd))

Loop:
	for i := 0; i < len(p.Ifd); i++ {
		errors[i] = make([]error, len(p.Ifd[i]))
		for j := 0; j < len(p.Ifd[i]); j++ {
			if cfg[i][j], errors[i][j] = p.Ifd[i][j].ImageConfig(); errors[i][j] != nil {
				break Loop
			}
		}
	}
	for i := 0; i < len(errors); i++ {
		for j := 0; j < len(errors[i]); j++ {
			if errors[i][j] != nil {
				err = errors[i][j]
				return
			}
		}
	}

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

	m, err = p.DecodeImage(0, 0)
	return
}

func DecodeAll(r io.Reader) (m [][]image.Image, errors [][]error, err error) {
	var p *Reader
	if p, err = OpenReader(r); err != nil {
		return
	}
	defer p.Close()

	m = make([][]image.Image, p.ImageNum())
	errors = make([][]error, p.ImageNum())

Loop:
	for i := 0; i < p.ImageNum(); i++ {
		m[i] = make([]image.Image, p.SubImageNum(i))
		errors[i] = make([]error, len(p.Ifd[i]))
		for j := 0; j < p.SubImageNum(i); j++ {
			if m[i][j], errors[i][j] = p.DecodeImage(i, j); errors[i][j] != nil {
				break Loop
			}
		}
	}
	for i := 0; i < len(errors); i++ {
		for j := 0; j < len(errors[i]); j++ {
			if errors[i][j] != nil {
				err = errors[i][j]
				return
			}
		}
	}
	return
}

func init() {
	image.RegisterFormat("tiff", ClassicTiffLittleEnding, Decode, DecodeConfig)
	image.RegisterFormat("tiff", ClassicTiffBigEnding, Decode, DecodeConfig)
	image.RegisterFormat("tiff", BigTiffLittleEnding, Decode, DecodeConfig)
	image.RegisterFormat("tiff", BigTiffBigEnding, Decode, DecodeConfig)
}
