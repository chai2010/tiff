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
	Offset   uint64
	EntryMap map[TagType]*IFDEntry
	Next     uint64
}

func ReadIFD(r io.ReadSeeker, h *Header, offset uint64) (ifd *IFD, err error) {
	if !h.Valid() {
		err = fmt.Errorf("tiff.go: ReadIFD, invalid header: %v", h)
		return
	}
	ifd = &IFD{
		Header:   h,
		Offset:   offset,
		EntryMap: make(map[TagType]*IFDEntry),
	}
	if h.Version == ClassicTIFF {
		if err = ifd.readIFD(r, offset); err != nil {
			return
		}
	} else {
		if err = ifd.readIFD8(r, offset); err != nil {
			return
		}
	}
	return
}

func (p *IFD) readIFD(r io.ReadSeeker, offset uint64) (err error) {
	if _, err = r.Seek(int64(offset), 0); err != nil {
		return
	}

	var entryNum uint16
	if err = binary.Read(r, p.Header.ByteOrder, &entryNum); err != nil {
		return
	}

	for i := 0; i < int(entryNum); i++ {
		var entry *IFDEntry
		if entry, err = p.readIFDEntry(r); err != nil {
			return
		}
		p.EntryMap[entry.Tag] = entry
	}

	var nextIfdOffset uint32
	if err = binary.Read(r, p.Header.ByteOrder, &nextIfdOffset); err != nil {
		return
	}
	p.Next = uint64(nextIfdOffset)

	return
}

func (p *IFD) readIFD8(r io.ReadSeeker, offset uint64) (err error) {
	if _, err = r.Seek(int64(offset), 0); err != nil {
		return
	}

	var entryNum uint32
	if err = binary.Read(r, p.Header.ByteOrder, &entryNum); err != nil {
		return
	}

	for i := 0; i < int(entryNum); i++ {
		var entry *IFDEntry
		if entry, err = p.readIFDEntry8(r); err != nil {
			return
		}
		p.EntryMap[entry.Tag] = entry
	}

	var nextIfdOffset uint64
	if err = binary.Read(r, p.Header.ByteOrder, &nextIfdOffset); err != nil {
		return
	}
	p.Next = nextIfdOffset

	return
}

func (p *IFD) readIFDEntry(r io.ReadSeeker) (entry *IFDEntry, err error) {
	var entryTag TagType
	if err = binary.Read(r, p.Header.ByteOrder, &entryTag); err != nil {
		return
	}

	var entryDataType DataType
	if err = binary.Read(r, p.Header.ByteOrder, &entryDataType); err != nil {
		return
	}

	var elemCount uint32
	if err = binary.Read(r, p.Header.ByteOrder, &elemCount); err != nil {
		return
	}

	var elemOffset uint32
	if err = binary.Read(r, p.Header.ByteOrder, &elemOffset); err != nil {
		return
	}
	entry = &IFDEntry{
		IFD:      p,
		Tag:      entryTag,
		DataType: entryDataType,
		Count:    uint64(elemCount),
		Offset:   uint64(elemOffset),
	}
	return
}

func (p *IFD) readIFDEntry8(r io.ReadSeeker) (entry *IFDEntry, err error) {
	var entryTag TagType
	if err = binary.Read(r, p.Header.ByteOrder, &entryTag); err != nil {
		return
	}

	var entryDataType DataType
	if err = binary.Read(r, p.Header.ByteOrder, &entryDataType); err != nil {
		return
	}

	var elemCount uint64
	if err = binary.Read(r, p.Header.ByteOrder, &elemCount); err != nil {
		return
	}

	var elemOffset uint64
	if err = binary.Read(r, p.Header.ByteOrder, &elemOffset); err != nil {
		return
	}
	entry = &IFDEntry{
		IFD:      p,
		Tag:      entryTag,
		DataType: entryDataType,
		Count:    elemCount,
		Offset:   elemOffset,
	}
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
