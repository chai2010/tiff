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
	ifd    []*IFD
	cfg    []image.Config
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
	if !hdr.Valid() {
		err = fmt.Errorf("tiff: OpenReader, invalid header: %v", hdr)
		return
	}

	var ifdList []*IFD
	var cfgList []image.Config
	for offset := hdr.Offset; offset != 0; {
		ifd, err := ReadIFD(rs, hdr, offset)
		if err != nil {
			return nil, err
		}
		cfg, err := ifd.ImageConfig()
		if err != nil {
			return nil, err
		}

		ifdList = append(ifdList, ifd)
		cfgList = append(cfgList, cfg)
		offset = ifd.Offset
	}

	p = &Reader{
		rs:     rs,
		header: hdr,
		ifd:    ifdList,
		cfg:    cfgList,
	}
	return
}

func (p *Reader) ImageNum() int {
	return len(p.ifd)
}

func (p *Reader) ImageConfig(idx int) image.Config {
	return p.cfg[idx]
}

func (p *Reader) DecodeImage(idx int) (m image.Image, err error) {
	imgRect := image.Rect(0, 0, p.cfg[idx].Width, p.cfg[idx].Height)
	imageType := p.ifd[idx].ImageType()

	switch imageType {
	case ImageType_Bilevel, ImageType_BilevelInvert:
		m = image.NewGray(imgRect)
	case ImageType_Gray, ImageType_GrayInvert:
		if p.ifd[idx].Depth() == 16 {
			m = image.NewGray16(imgRect)
		} else {
			m = image.NewGray(imgRect)
		}
	case ImageType_Paletted:
		m = image.NewPaletted(imgRect, p.ifd[idx].ColorMap())
	case ImageType_NRGBA:
		if p.ifd[idx].Depth() == 16 {
			m = image.NewNRGBA64(imgRect)
		} else {
			m = image.NewNRGBA(imgRect)
		}
	case ImageType_RGB, ImageType_RGBA:
		if p.ifd[idx].Depth() == 16 {
			m = image.NewRGBA64(imgRect)
		} else {
			m = image.NewRGBA(imgRect)
		}
	}
	if m == nil {
		err = fmt.Errorf("tiff: Decode, unknown format")
		return
	}

	blocksAcross := p.ifd[idx].BlocksAcross()
	blocksDown := p.ifd[idx].BlocksDown()

	for i := 0; i < blocksAcross; i++ {
		for j := 0; j < blocksDown; j++ {
			if err = p.ifd[idx].DecodeBlock(p.rs, i, j, m); err != nil {
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

	cfg = append(cfg, p.cfg...)
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
		if m[i], err = p.DecodeImage(0); err != nil {
			return
		}
	}
	return
}
