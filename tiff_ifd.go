// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
)

type IFD struct {
	Header   *Header
	EntryMap map[TagType]*IFDEntry
	Offset   int64 // next IFD
}

func NewIFD(hdr *Header, width, height, depth, channels int, kind reflect.Kind) (ifd *IFD) {
	return
}

func ReadIFD(r io.Reader, h *Header, offset int64) (ifd *IFD, err error) {
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

	if h.TiffType == TiffType_ClassicTIFF {
		ifd, err = readIFD(rs, h, offset)
		return
	} else {
		ifd, err = readIFD8(rs, h, offset)
		return
	}
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
		EntryMap: make(map[TagType]*IFDEntry),
	}

	// read IFDEntry
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
	p.Offset = int64(nextIfdOffset)

	// read IFDEntry Data
	for _, entry := range p.EntryMap {
		if entry.Data, err = readIFDEntryData(r, entry); err != nil {
			return
		}
	}
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
		EntryMap: make(map[TagType]*IFDEntry),
	}

	// read IFDEntry
	var entryNum uint64
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
	p.Offset = int64(nextIfdOffset)

	// read IFDEntry8 Data
	for _, entry := range p.EntryMap {
		if entry.Data, err = readIFDEntry8Data(r, entry); err != nil {
			return
		}
	}

	return
}

func readIFDEntry(r io.ReadSeeker, h *Header) (entry *IFDEntry, err error) {
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
		Count:    int(elemCount),
		Offset:   int64(elemOffset),
	}
	return
}

func readIFDEntry8(r io.ReadSeeker, h *Header) (entry *IFDEntry, err error) {
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
		Count:    int(elemCount),
		Offset:   int64(elemOffset),
	}
	return
}

func readIFDEntryData(r io.ReadSeeker, entry *IFDEntry) (data []byte, err error) {
	valSize := entry.DataType.ByteSize() * entry.Count
	if valSize <= 4 {
		var buf bytes.Buffer
		var offset = uint32(entry.Offset)
		binary.Write(&buf, entry.Header.ByteOrder, offset)
		data = buf.Bytes()
		return
	}

	if entry.Offset < int64(entry.Header.HeadSize()) {
		err = fmt.Errorf("tiff: readIFDEntryData, bad offset %v", entry.Offset)
		return
	}
	if _, err = r.Seek(entry.Offset, 0); err != nil {
		return
	}
	data = make([]byte, valSize)
	if _, err = r.Read(data); err != nil {
		return
	}
	return
}

func readIFDEntry8Data(r io.ReadSeeker, entry *IFDEntry) (data []byte, err error) {
	valSize := entry.DataType.ByteSize() * entry.Count
	if valSize <= 8 {
		var buf bytes.Buffer
		var offset = uint64(entry.Offset)
		binary.Write(&buf, entry.Header.ByteOrder, offset)
		data = buf.Bytes()
		return
	}

	if entry.Offset < int64(entry.Header.HeadSize()) {
		err = fmt.Errorf("tiff: readIFDEntryData, bad offset %v", entry.Offset)
		return
	}
	if _, err = r.Seek(entry.Offset, 0); err != nil {
		return
	}
	data = make([]byte, valSize)
	if _, err = r.Read(data); err != nil {
		return
	}
	return
}
