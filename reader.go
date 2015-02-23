// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"fmt"
	"image"
	"io"
)

type Reader struct {
	rs     *seekioReader
	header *Header
	ifd    *IFD
	cfg    image.Config
}

func OpenReader(r io.Reader) (p *Reader, err error) {
	rs := openSeekioReader(r, -1)
	defer func() {
		if err != nil && rs != nil {
			rs.Close()
		}
	}()

	hdr, err := ReadHeader(rs)
	if err != nil {
		return
	}
	ifd, err := ReadIFD(rs, hdr, hdr.Offset)
	if err != nil {
		return
	}
	cfg, err := ifd.ImageConfig()
	if err != nil {
		return
	}

	p = &Reader{
		rs:     rs,
		header: hdr,
		ifd:    ifd,
		cfg:    cfg,
	}
	return
}

func (p *Reader) ImageConfig() image.Config {
	return p.cfg
}

func (p *Reader) Decode() (m image.Image, err error) {
	imgRect := image.Rect(0, 0, p.cfg.Width, p.cfg.Height)
	imageType := p.ifd.ImageType()

	switch imageType {
	case ImageType_Bilevel, ImageType_BilevelInvert:
		m = image.NewGray(imgRect)
	case ImageType_Gray, ImageType_GrayInvert:
		if p.ifd.Depth() == 16 {
			m = image.NewGray16(imgRect)
		} else {
			m = image.NewGray(imgRect)
		}
	case ImageType_Paletted:
		m = image.NewPaletted(imgRect, p.ifd.ColorMap())
	case ImageType_NRGBA:
		if p.ifd.Depth() == 16 {
			m = image.NewNRGBA64(imgRect)
		} else {
			m = image.NewNRGBA(imgRect)
		}
	case ImageType_RGB, ImageType_RGBA:
		if p.ifd.Depth() == 16 {
			m = image.NewRGBA64(imgRect)
		} else {
			m = image.NewRGBA(imgRect)
		}
	}
	if m == nil {
		err = fmt.Errorf("tiff: Decode, unknown format")
		return
	}

	blocksAcross := p.ifd.BlocksAcross()
	blocksDown := p.ifd.BlocksDown()

	for i := 0; i < blocksAcross; i++ {
		for j := 0; j < blocksDown; j++ {
			if err = p.ifd.DecodeBlock(p.rs, i, j, m); err != nil {
				return
			}
		}
	}
	return
}

func (p *Reader) Close() (err error) {
	if p != nil && p.rs != nil {
		err = p.rs.Close()
	}
	*p = Reader{}
	return
}

// DecodeConfig returns the color model and dimensions of a TIFF image without
// decoding the entire image.
func DecodeConfig(r io.Reader) (cfg image.Config, err error) {
	var reader *Reader
	if reader, err = OpenReader(r); err != nil {
		return
	}
	defer reader.Close()

	cfg = reader.ImageConfig()
	return
}

// Decode reads a TIFF image from r and returns it as an image.Image.
// The type of Image returned depends on the contents of the TIFF.
func Decode(r io.Reader) (m image.Image, err error) {
	var reader *Reader
	if reader, err = OpenReader(r); err != nil {
		return
	}
	defer reader.Close()

	m, err = reader.Decode()
	return
}

func DecodeAll(r io.Reader) (m []image.Image, err error) {
	return
}
