// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

const (
	ClassicTIFF = 42 // uint16, Header.Version
	BigTIFF     = 43
)

type Header struct {
	ByteOrder binary.ByteOrder
	Version   uint16
	Offset    int64
}

func NewHeader(isBigTiff bool, offset int64) *Header {
	if isBigTiff {
		return &Header{
			ByteOrder: binary.LittleEndian,
			Version:   BigTIFF,
			Offset:    offset,
		}
	} else {
		return &Header{
			ByteOrder: binary.LittleEndian,
			Version:   ClassicTIFF,
			Offset:    offset,
		}
	}
}

func ReadHeader(r io.Reader) (header *Header, err error) {
	var ra io.ReaderAt
	if ra, _ = r.(io.ReaderAt); ra == nil {
		rs := openSeekioReader(r, 0)
		defer rs.Close()
		ra = rs
	}

	// read classic TIFF header
	var data [8]byte
	if _, err = ra.ReadAt(data[:8], 0); err != nil {
		return
	}
	header = new(Header)

	// byte order
	switch {
	case data[0] == 'I' && data[1] == 'I':
		header.ByteOrder = binary.LittleEndian
	case data[0] == 'M' && data[1] == 'M':
		header.ByteOrder = binary.BigEndian
	default:
		err = fmt.Errorf("tiff: ReadHeader, bad order: %v", data[:2])
		return
	}

	// version: ClassicTIFF or BigTIFF
	header.Version = header.ByteOrder.Uint16(data[2:4])
	if v := header.Version; v != ClassicTIFF && v != BigTIFF {
		err = fmt.Errorf("tiff: ReadHeader, bad version: %v", data[2:4])
		return
	}

	// offset
	switch header.Version {
	case ClassicTIFF:
		header.Offset = int64(header.ByteOrder.Uint32(data[4:8]))
	case BigTIFF:
		byte46 := header.ByteOrder.Uint16(data[4:6])
		byte68 := header.ByteOrder.Uint16(data[6:8])
		if byte46 != 8 || byte68 != 0 {
			err = fmt.Errorf("tiff: ReadHeader, bad offset: %v", data[4:8])
			return
		}
		if _, err = ra.ReadAt(data[:8], 8); err != nil {
			return
		}
		header.Offset = int64(header.ByteOrder.Uint64(data[0:8]))
	}
	if header.Offset < int64(header.HeadSize()) {
		err = fmt.Errorf("tiff: ReadHeader, bad offset: %v", data[4:8])
		return
	}

	return
}

func (p *Header) Bytes() []byte {
	if !p.Valid() {
		return nil
	}

	var d [16]byte
	switch p.ByteOrder {
	case binary.LittleEndian:
		d[0], d[1] = 'I', 'I'
	case binary.BigEndian:
		d[0], d[1] = 'M', 'M'
	}

	if p.Version == ClassicTIFF {
		p.ByteOrder.PutUint16(d[2:4], p.Version)
		p.ByteOrder.PutUint32(d[4:8], uint32(p.Offset))
		return d[:8]
	} else {
		p.ByteOrder.PutUint16(d[2:4], p.Version)
		p.ByteOrder.PutUint16(d[4:6], 8)
		p.ByteOrder.PutUint16(d[6:8], 0)
		p.ByteOrder.PutUint64(d[8:], uint64(p.Offset))
		return d[:16]
	}
}

func (p *Header) Valid() bool {
	if x := p.ByteOrder; x != binary.LittleEndian && x != binary.BigEndian {
		return false
	}
	if x := p.Version; x != ClassicTIFF && x != BigTIFF {
		return false
	}
	if p.Offset < int64(p.HeadSize()) {
		return false
	}
	if p.Version == ClassicTIFF {
		if p.Offset > math.MaxUint32 {
			return false
		}
	}
	return true
}

func (p *Header) HeadSize() int {
	if p.Version == ClassicTIFF {
		return 8
	}
	if p.Version == BigTIFF {
		return 16
	}
	return 0
}

func (p *Header) IsBigTiff() bool {
	return p.Version == BigTIFF
}

func (p *Header) String() string {
	orderName := "Unknown"
	switch p.ByteOrder {
	case binary.LittleEndian:
		orderName = "LittleEndian"
	case binary.BigEndian:
		orderName = "BigEndian"
	}
	return fmt.Sprintf(
		`tiff.Header{ ByteOrder:%s; Version:%d; Offset:0x%08x }`,
		orderName, p.Version, p.Offset,
	)
}
