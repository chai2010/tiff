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
	Ifd    [][]*IFD

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
	for offset := p.Header.FirstIFD; offset != 0; offset = ifd.NextIFD {
		var ifdList []*IFD

		if ifd, err = ReadIFD(rs, p.Header, offset); err != nil {
			return
		}
		ifdList = append(ifdList, ifd)

		subIfdOffsets, _ := ifd.TagGetter().GetSubIFD()
		for _, subOffset := range subIfdOffsets {
			ifd, _ = ReadIFD(rs, p.Header, subOffset)
			ifdList = append(ifdList, ifd)
		}
		p.Ifd = append(p.Ifd, ifdList)
	}

	p.Reader = rs
	p.rs = rs
	return
}

func (p *Reader) ImageNum() int {
	return len(p.Ifd)
}
func (p *Reader) SubImageNum(i int) int {
	return len(p.Ifd[i])
}

func (p *Reader) ImageConfig(i, j int) (image.Config, error) {
	return p.Ifd[i][j].ImageConfig()
}

func (p *Reader) DecodeImage(i, j int) (m image.Image, err error) {
	cfg, err := p.ImageConfig(i, j)
	if err != nil {
		return
	}
	imgRect := image.Rect(0, 0, cfg.Width, cfg.Height)
	imageType := p.Ifd[i][j].ImageType()

	switch imageType {
	case ImageType_Bilevel, ImageType_BilevelInvert:
		m = image.NewGray(imgRect)
	case ImageType_Gray, ImageType_GrayInvert:
		if p.Ifd[i][j].Depth() == 16 {
			m = image.NewGray16(imgRect)
		} else {
			m = image.NewGray(imgRect)
		}
	case ImageType_Paletted:
		m = image.NewPaletted(imgRect, p.Ifd[i][j].ColorMap())
	case ImageType_NRGBA:
		if p.Ifd[i][j].Depth() == 16 {
			m = image.NewNRGBA64(imgRect)
		} else {
			m = image.NewNRGBA(imgRect)
		}
	case ImageType_RGB, ImageType_RGBA:
		if p.Ifd[i][j].Depth() == 16 {
			m = image.NewRGBA64(imgRect)
		} else {
			m = image.NewRGBA(imgRect)
		}
	}
	if m == nil {
		err = fmt.Errorf("tiff: Decode, unknown format")
		return
	}

	blocksAcross := p.Ifd[i][j].BlocksAcross()
	blocksDown := p.Ifd[i][j].BlocksDown()

	for col := 0; col < blocksAcross; col++ {
		for row := 0; row < blocksDown; row++ {
			if err = p.Ifd[i][j].DecodeBlock(p.rs, col, row, m); err != nil {
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
