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
	Reader io.ReadSeeker
	Header *Header
	Ifd    []*IFD
	Cfg    []image.Config

	rs *seekioReader
}

func OpenReader(r io.Reader) (p *Reader, err error) {
	rs := openSeekioReader(r, -1)
	defer func() {
		if err != nil && rs != nil {
			rs.Close()
		}
	}()

	p = &Reader{}
	if p.Header, err = ReadHeader(rs); err != nil {
		return
	}
	if !p.Header.Valid() {
		err = fmt.Errorf("tiff: OpenReader, invalid header: %v", p.Header)
		return
	}

	var ifd *IFD
	for offset := p.Header.Offset; offset != 0; offset = ifd.Offset {
		ifd, err = ReadIFD(rs, p.Header, offset)
		if err != nil {
			return nil, err
		}
		cfg, err := ifd.ImageConfig()
		if err != nil {
			return nil, err
		}

		p.Ifd = append(p.Ifd, ifd)
		p.Cfg = append(p.Cfg, cfg)
	}

	p.Reader = rs
	p.rs = rs
	return
}

func (p *Reader) ImageNum() int {
	return len(p.Ifd)
}

func (p *Reader) ImageConfig(idx int) image.Config {
	return p.Cfg[idx]
}

func (p *Reader) DecodeImage(idx int) (m image.Image, err error) {
	imgRect := image.Rect(0, 0, p.Cfg[idx].Width, p.Cfg[idx].Height)
	imageType := p.Ifd[idx].ImageType()

	switch imageType {
	case ImageType_Bilevel, ImageType_BilevelInvert:
		m = image.NewGray(imgRect)
	case ImageType_Gray, ImageType_GrayInvert:
		if p.Ifd[idx].Depth() == 16 {
			m = image.NewGray16(imgRect)
		} else {
			m = image.NewGray(imgRect)
		}
	case ImageType_Paletted:
		m = image.NewPaletted(imgRect, p.Ifd[idx].ColorMap())
	case ImageType_NRGBA:
		if p.Ifd[idx].Depth() == 16 {
			m = image.NewNRGBA64(imgRect)
		} else {
			m = image.NewNRGBA(imgRect)
		}
	case ImageType_RGB, ImageType_RGBA:
		if p.Ifd[idx].Depth() == 16 {
			m = image.NewRGBA64(imgRect)
		} else {
			m = image.NewRGBA(imgRect)
		}
	}
	if m == nil {
		err = fmt.Errorf("tiff: Decode, unknown format")
		return
	}

	blocksAcross := p.Ifd[idx].BlocksAcross()
	blocksDown := p.Ifd[idx].BlocksDown()

	for i := 0; i < blocksAcross; i++ {
		for j := 0; j < blocksDown; j++ {
			if err = p.Ifd[idx].DecodeBlock(p.rs, i, j, m); err != nil {
				return
			}
		}
	}
	return
}

func (p *Reader) Close() (err error) {
	if p != nil {
		if p.rs != nil {
			err = p.rs.Close()
		}
		*p = Reader{}
	}
	return
}
