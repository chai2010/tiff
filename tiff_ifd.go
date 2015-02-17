// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type IFD struct {
	Header   *Header
	Offset   int64
	EntryMap map[TagType]*IFDEntry
	Next     int64
}

func ReadIFD(r io.Reader, h *Header, offset int64) (ifds []*IFD, err error) {
	var rs io.ReadSeeker
	if rs, _ = r.(io.ReadSeeker); rs == nil {
		seekioReader := openSeekioReader(r, 0)
		defer seekioReader.Close()
		rs = seekioReader
	}

	if !h.Valid() {
		err = fmt.Errorf("tiff: ReadIFD, invalid header: %v", h)
		return
	}
	if h.Version == ClassicTIFF {
		for offset != 0 {
			var ifd *IFD
			if ifd, err = readIFD(rs, h, offset); err != nil {
				return
			}
			ifds = append(ifds, ifd)
			offset = ifd.Next
		}
	} else {
		for offset != 0 {
			var ifd *IFD
			if ifd, err = readIFD8(rs, h, offset); err != nil {
				return
			}
			ifds = append(ifds, ifd)
			offset = ifd.Next
		}
	}
	return
}

func readIFD(r io.ReadSeeker, h *Header, offset int64) (p *IFD, err error) {
	if offset == 0 {
		return
	}
	if _, err = r.Seek(offset, 0); err != nil {
		return
	}

	p = &IFD{
		Header:   h,
		Offset:   offset,
		EntryMap: make(map[TagType]*IFDEntry),
	}

	var entryNum uint16
	if err = binary.Read(r, h.ByteOrder, &entryNum); err != nil {
		return
	}

	for i := 0; i < int(entryNum); i++ {
		var entry *IFDEntry
		if entry, err = readIFDEntry(r, h); err != nil {
			return
		}
		p.EntryMap[entry.Tag] = entry
	}

	var nextIfdOffset uint32
	if err = binary.Read(r, h.ByteOrder, &nextIfdOffset); err != nil {
		return
	}
	p.Next = int64(nextIfdOffset)
	return
}

func readIFD8(r io.ReadSeeker, h *Header, offset int64) (p *IFD, err error) {
	if offset == 0 {
		return
	}
	if _, err = r.Seek(offset, 0); err != nil {
		return
	}

	p = &IFD{
		Header:   h,
		Offset:   offset,
		EntryMap: make(map[TagType]*IFDEntry),
	}

	var entryNum uint32
	if err = binary.Read(r, h.ByteOrder, &entryNum); err != nil {
		return
	}

	for i := 0; i < int(entryNum); i++ {
		var entry *IFDEntry
		if entry, err = readIFDEntry8(r, h); err != nil {
			return
		}
		p.EntryMap[entry.Tag] = entry
	}

	var nextIfdOffset uint64
	if err = binary.Read(r, h.ByteOrder, &nextIfdOffset); err != nil {
		return
	}
	p.Next = int64(nextIfdOffset)
	return
}

func readIFDEntry(r io.Reader, h *Header) (entry *IFDEntry, err error) {
	var entryTag TagType
	if err = binary.Read(r, h.ByteOrder, &entryTag); err != nil {
		return
	}

	var entryDataType DataType
	if err = binary.Read(r, h.ByteOrder, &entryDataType); err != nil {
		return
	}

	var elemCount uint32
	if err = binary.Read(r, h.ByteOrder, &elemCount); err != nil {
		return
	}

	var elemOffset uint32
	if err = binary.Read(r, h.ByteOrder, &elemOffset); err != nil {
		return
	}
	entry = &IFDEntry{
		Header:   h,
		Tag:      entryTag,
		DataType: entryDataType,
		Count:    uint64(elemCount),
		Offset:   uint64(elemOffset),
	}
	err = entry.Read(r)
	return
}

func readIFDEntry8(r io.Reader, h *Header) (entry *IFDEntry, err error) {
	var entryTag TagType
	if err = binary.Read(r, h.ByteOrder, &entryTag); err != nil {
		return
	}

	var entryDataType DataType
	if err = binary.Read(r, h.ByteOrder, &entryDataType); err != nil {
		return
	}

	var elemCount uint64
	if err = binary.Read(r, h.ByteOrder, &elemCount); err != nil {
		return
	}

	var elemOffset uint64
	if err = binary.Read(r, h.ByteOrder, &elemOffset); err != nil {
		return
	}
	entry = &IFDEntry{
		Header:   h,
		Tag:      entryTag,
		DataType: entryDataType,
		Count:    elemCount,
		Offset:   elemOffset,
	}
	err = entry.Read(r)
	return
}

func (p *IFD) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "tiff.IFD {\n")
	fmt.Fprintf(&buf, "  Offset: 0x%08x,\n", p.Offset)
	for _, v := range p.EntryMap {
		fmt.Fprintf(&buf, "  %v,\n", v)
	}
	fmt.Fprintf(&buf, "  Next: 0x%08x,\n", p.Next)
	fmt.Fprintf(&buf, "}\n")
	return buf.String()
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
		return 2 + len(p.EntryMap)*12 + 4
	} else {
		return 8 + len(p.EntryMap)*20 + 8
	}
}

func (p *IFD) ImageSize() (width, height int) {
	return
}

func (p *IFD) BitsPerSample() int {
	return 0
}

func (p *IFD) Compression() CompressType {
	return CompressType_Nil
}

func (p *IFD) CellSize() (width, height float64) {
	return
}

func (p *IFD) BlockSize() (width, height int) {
	return
}

func (p *IFD) BlockOffsets() []int64 {
	return nil
}

func (p *IFD) DocumentName() string {
	return ""
}

func (p *IFD) ImageDescription() string {
	return ""
}

func (p *IFD) Make() string {
	return ""
}

func (p *IFD) Model() string {
	return ""
}

func (p *IFD) Copyright() string {
	return ""
}
