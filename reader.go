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

	for offset := p.Header.FirstIFD; offset != 0; {
		var ifd *IFD
		var ifdList []*IFD

		if ifd, err = ReadIFD(rs, p.Header, offset); err != nil {
			return
		}
		ifdList = append(ifdList, ifd)
		offset = ifd.NextIFD

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

func (p *Reader) ImageBlocksAcross(i, j int) int {
	return p.Ifd[i][j].BlocksAcross()
}

func (p *Reader) ImageBlocksDown(i, j int) int {
	return p.Ifd[i][j].BlocksDown()
}

func (p *Reader) ImageBlockBounds(i, j, col, row int) image.Rectangle {
	return p.Ifd[i][j].BlockBounds(col, row)
}

func (p *Reader) DecodeImage(i, j int) (m image.Image, err error) {
	cfg, err := p.ImageConfig(i, j)
	if err != nil {
		return
	}
	imgRect := image.Rect(0, 0, cfg.Width, cfg.Height)
	if m, err = NewImageWithIFD(imgRect, p.Ifd[i][j]); err != nil {
		return
	}

	blocksAcross := p.ImageBlocksAcross(i, j)
	blocksDown := p.ImageBlocksDown(i, j)

	for col := 0; col < blocksAcross; col++ {
		for row := 0; row < blocksDown; row++ {
			if err = p.Ifd[i][j].DecodeBlock(p.rs, col, row, m); err != nil {
				return
			}
		}
	}
	return
}

func (p *Reader) DecodeImageBlock(i, j, col, row int) (m image.Image, err error) {
	r := p.ImageBlockBounds(i, j, col, row)
	if m, err = NewImageWithIFD(r, p.Ifd[i][j]); err != nil {
		return
	}
	if err = p.Ifd[i][j].DecodeBlock(p.rs, col, row, m); err != nil {
		return
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
