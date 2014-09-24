// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"fmt"
	"io"
)

type IFD struct {
	Header *Header
	Entry  []IFDEntry
	Next   uint64
}

func ReadIFD(r io.ReaderAt, h *Header) (ifd *IFD, err error) {
	if !h.Valid() {
		err = fmt.Errorf("tiff.go: ReadIFD, invalid header: %v", h)
		return
	}
	ifd = &IFD{Header: h}
	if h.Version == ClassicTIFF {
		if err = ifd.readIFD(r); err != nil {
			return
		}
	} else {
		if err = ifd.readIFD8(r); err != nil {
			return
		}
	}
	return
}

func (p *IFD) readIFD(r io.ReaderAt) (err error) {
	return
}

func (p *IFD) readIFD8(r io.ReaderAt) (err error) {
	return
}

func (p *IFD) readIFDEntry(r io.ReaderAt, offset uint64) (entry IFDEntry, err error) {
	return
}

func (p *IFD) Valid() bool {
	if p.Header == nil || !p.Header.Valid() {
		return false
	}
	return true
}

func (p *IFD) IfdSize() int {
	if !p.Valid() {
		return 0
	}
	if p.Header.Version == ClassicTIFF {
		return 2 + len(p.Entry)*12 + 4
	} else {
		return 8 + len(p.Entry)*20 + 8
	}
}
